package masterdata

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

const decodedImportSchemaVersion = 3

type DecodedJSONFile struct {
	ID         uint      `gorm:"primaryKey"`
	MasterName string    `gorm:"column:table_name;size:128;not null;uniqueIndex"`
	SQLTable  string    `gorm:"size:128;not null;index"`
	SchemaVersion int      `gorm:"not null;default:0;index"`
	FileName  string    `gorm:"size:255;not null"`
	Source    string    `gorm:"size:1024;not null"`
	SHA256    string    `gorm:"size:64;not null;index"`
	Size      int64     `gorm:"not null"`
	RowCount  int       `gorm:"not null"`
	ImportedAt time.Time `gorm:"not null;index"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (DecodedJSONFile) TableName() string { return "master_json_files" }

type ImportDecodedResult struct {
	Root          string
	Files         int
	ImportedFiles int
	SkippedFiles  int
	Rows          int
}

func ImportDecodedJSON(ctx context.Context, db *gorm.DB, root string, log *zap.Logger) (*ImportDecodedResult, error) {
	if db == nil {
		return nil, fmt.Errorf("database is nil")
	}
	root = strings.TrimSpace(root)
	if root == "" {
		return &ImportDecodedResult{}, nil
	}
	info, err := os.Stat(root)
	if err != nil {
		if os.IsNotExist(err) {
			if log != nil {
				log.Warn("master decoded json root not found", zap.String("root", root))
			}
			return &ImportDecodedResult{Root: root}, nil
		}
		return nil, fmt.Errorf("stat master decoded json root: %w", err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("master decoded json root is not a directory: %s", root)
	}
	if err := db.WithContext(ctx).AutoMigrate(&DecodedJSONFile{}); err != nil {
		return nil, fmt.Errorf("migrate decoded master json tables: %w", err)
	}

	res := &ImportDecodedResult{Root: root}
	err = filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.EqualFold(filepath.Ext(path), ".json") {
			return nil
		}
		res.Files++
		imported, rows, err := importDecodedJSONFile(ctx, db, root, path)
		if err != nil {
			return err
		}
		if imported {
			res.ImportedFiles++
			res.Rows += rows
		} else {
			res.SkippedFiles++
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("import decoded master json: %w", err)
	}
	if log != nil {
		log.Info("master decoded json import complete",
			zap.String("root", res.Root),
			zap.Int("files", res.Files),
			zap.Int("imported_files", res.ImportedFiles),
			zap.Int("skipped_files", res.SkippedFiles),
			zap.Int("rows", res.Rows),
		)
	}
	return res, nil
}

func importDecodedJSONFile(ctx context.Context, db *gorm.DB, root string, path string) (bool, int, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return false, 0, fmt.Errorf("read %s: %w", path, err)
	}
	sum := sha256.Sum256(body)
	hash := hex.EncodeToString(sum[:])
	tableName := decodedTableName(root, path)

	var existing DecodedJSONFile
	sqlTable := decodedSQLTableName(tableName)
	err = db.WithContext(ctx).Where("table_name = ?", tableName).First(&existing).Error
	if err == nil && existing.SHA256 == hash && existing.SchemaVersion == decodedImportSchemaVersion && db.Migrator().HasTable(sqlTable) {
		return false, existing.RowCount, nil
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, 0, err
	}

	rows, err := decodeJSONRows(body)
	if err != nil {
		return false, 0, fmt.Errorf("decode %s: %w", path, err)
	}
	info, err := os.Stat(path)
	if err != nil {
		return false, 0, err
	}

	now := time.Now()
	tableData := buildDecodedTableData(rows)
	assignChildTableNames(sqlTable, &tableData)
	err = db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Migrator().DropTable(sqlTable); err != nil {
			return err
		}
		if err := dropChildTables(tx, sqlTable); err != nil {
			return err
		}
		if err := tx.Where("table_name = ?", tableName).Delete(&DecodedJSONFile{}).Error; err != nil {
			return err
		}
		if len(tableData.ParentRows) > 0 {
			if err := createDecodedJSONTable(tx, sqlTable, tableData.ParentColumns, false); err != nil {
				return err
			}
			for i, row := range tableData.ParentRows {
				if err := insertDecodedJSONRow(tx, sqlTable, tableData.ParentColumns, i, row, "_row_index", "_master_id"); err != nil {
					return err
				}
			}
		}
		for _, child := range tableData.ChildTables {
			if len(child.Rows) == 0 {
				continue
			}
			if err := createDecodedJSONTable(tx, child.SQLTable, child.Columns, true); err != nil {
				return err
			}
			for i, row := range child.Rows {
				if err := insertDecodedJSONRow(tx, child.SQLTable, child.Columns, i, row, "_row_index", "_parent_master_id"); err != nil {
					return err
				}
			}
		}
		file := DecodedJSONFile{
			MasterName: tableName,
			SQLTable:  sqlTable,
			SchemaVersion: decodedImportSchemaVersion,
			FileName:  filepath.Base(path),
			Source:    filepath.ToSlash(path),
			SHA256:    hash,
			Size:      info.Size(),
			RowCount:  len(rows),
			ImportedAt: now,
		}
		if err := tx.Create(&file).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, 0, err
	}
	return true, len(rows), nil
}

func decodedTableName(root string, path string) string {
	rel, err := filepath.Rel(root, path)
	if err != nil {
		rel = filepath.Base(path)
	}
	rel = strings.TrimSuffix(filepath.ToSlash(rel), ".json")
	rel = strings.Trim(rel, "/")
	if rel == "" {
		return strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	}
	return rel
}

func decodedSQLTableName(tableName string) string {
	return limitedIdentifier("master_"+safeIdentifier(tableName), 64)
}

func dropChildTables(db *gorm.DB, parentTable string) error {
	tables, err := db.Migrator().GetTables()
	if err != nil {
		return err
	}
	prefix := parentTable + "_"
	for _, table := range tables {
		if strings.HasPrefix(table, prefix) {
			if err := db.Migrator().DropTable(table); err != nil {
				return err
			}
		}
	}
	return nil
}

func decodeJSONRows(body []byte) ([]map[string]any, error) {
	dec := json.NewDecoder(bytes.NewReader(body))
	dec.UseNumber()
	var v any
	if err := dec.Decode(&v); err != nil {
		return nil, err
	}
	switch t := v.(type) {
	case []any:
		rows := make([]map[string]any, 0, len(t))
		for i, item := range t {
			rows = append(rows, normalizeDecodedRow(fmt.Sprintf("%d", i), item))
		}
		return rows, nil
	case map[string]any:
		keys := make([]string, 0, len(t))
		for key := range t {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		rows := make([]map[string]any, 0, len(keys))
		for _, key := range keys {
			rows = append(rows, normalizeDecodedRow(key, t[key]))
		}
		return rows, nil
	default:
		return []map[string]any{normalizeDecodedRow("0", v)}, nil
	}
}

func normalizeDecodedRow(key string, row any) map[string]any {
	if m, ok := row.(map[string]any); ok {
		if _, exists := m["__key"]; !exists && key != "" {
			m["__key"] = key
		}
		return m
	}
	return map[string]any{"__key": key, "value": row}
}

func extractMasterID(m map[string]any) string {
	for _, key := range []string{"Id", "ID", "id"} {
		if v, ok := m[key]; ok {
			return jsonValueString(v)
		}
	}
	for key, v := range m {
		lower := strings.ToLower(key)
		if strings.HasSuffix(lower, "masterid") || strings.HasSuffix(lower, "master_id") {
			if s := jsonValueString(v); s != "" {
				return s
			}
		}
	}
	if v, ok := m["__key"]; ok {
		return jsonValueString(v)
	}
	return ""
}

type decodedColumn struct {
	JSONKey string
	SQLName string
	Kind    string
}

type decodedTableData struct {
	ParentRows    []map[string]any
	ParentColumns []decodedColumn
	ChildTables   []decodedChildTable
}

type decodedChildTable struct {
	Name     string
	SQLTable string
	Rows     []map[string]any
	Columns  []decodedColumn
}

func buildDecodedTableData(rows []map[string]any) decodedTableData {
	parentRows := make([]map[string]any, 0, len(rows))
	childRows := map[string][]map[string]any{}
	for i, row := range rows {
		parent := map[string]any{}
		parentID := extractMasterID(row)
		flattenDecodedObject("", row, parent, func(field string, arr []any) {
			childName := safeIdentifier(field)
			if childName == "" {
				childName = "items"
			}
			for itemIndex, item := range arr {
				childGlobalIndex := len(childRows[childName])
				child := normalizeDecodedRow(fmt.Sprintf("%d", itemIndex), item)
				flatChild := map[string]any{
					"_parent_row_index": i,
					"_parent_master_id":  parentID,
					"_item_index":        itemIndex,
				}
				flattenDecodedObject("", child, flatChild, func(nestedField string, nestedArr []any) {
					grandChildName := safeIdentifier(childName + "_" + safeIdentifier(nestedField))
					for nestedIndex, nestedItem := range nestedArr {
						grandChild := normalizeDecodedRow(fmt.Sprintf("%d", nestedIndex), nestedItem)
						flatGrandChild := map[string]any{
							"_parent_row_index": childGlobalIndex,
							"_parent_master_id":  parentID,
							"_item_index":        nestedIndex,
						}
						flattenDecodedObject("", grandChild, flatGrandChild, func(deeperField string, deeperArr []any) {
							flatGrandChild[safeIdentifier(deeperField)] = deeperArr
						})
						childRows[grandChildName] = append(childRows[grandChildName], flatGrandChild)
					}
				})
				childRows[childName] = append(childRows[childName], flatChild)
			}
		})
		parentRows = append(parentRows, parent)
	}

	out := decodedTableData{
		ParentRows:    parentRows,
		ParentColumns: inferColumns(parentRows),
	}
	childNames := make([]string, 0, len(childRows))
	for name := range childRows {
		childNames = append(childNames, name)
	}
	sort.Strings(childNames)
	for _, name := range childNames {
		childTable := decodedChildTable{
			Name:     name,
			SQLTable: "",
			Rows:     childRows[name],
			Columns:  inferColumns(childRows[name]),
		}
		out.ChildTables = append(out.ChildTables, childTable)
	}
	return out
}

func assignChildTableNames(parentSQLTable string, data *decodedTableData) {
	for i := range data.ChildTables {
		data.ChildTables[i].SQLTable = limitedIdentifier(parentSQLTable+"_"+safeIdentifier(data.ChildTables[i].Name), 64)
	}
}

func flattenDecodedObject(prefix string, row map[string]any, out map[string]any, array func(field string, arr []any)) {
	keys := make([]string, 0, len(row))
	for key := range row {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		value := row[key]
		field := safeIdentifier(key)
		if prefix != "" {
			field = prefix + "_" + field
		}
		switch t := value.(type) {
		case []any:
			array(field, t)
		case map[string]any:
			flattenDecodedObject(field, t, out, array)
		default:
			out[field] = value
		}
	}
}

func inferColumns(rows []map[string]any) []decodedColumn {
	kinds := map[string]string{}
	for _, row := range rows {
		for key, value := range row {
			kind := inferJSONKind(value)
			if current := kinds[key]; current == "" {
				kinds[key] = kind
			} else {
				kinds[key] = mergeJSONKind(current, kind)
			}
		}
	}
	keys := make([]string, 0, len(kinds))
	for key := range kinds {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	used := map[string]int{"_row_index": 1, "_master_id": 1}
	out := make([]decodedColumn, 0, len(keys))
	for _, key := range keys {
		name := safeIdentifier(key)
		if name == "" {
			name = "value"
		}
		if n := used[name]; n > 0 {
			used[name] = n + 1
			name = fmt.Sprintf("%s_%d", name, n+1)
			used[name] = 1
		} else {
			used[name] = 1
		}
		out = append(out, decodedColumn{JSONKey: key, SQLName: name, Kind: kinds[key]})
	}
	return out
}

func inferJSONKind(v any) string {
	switch t := v.(type) {
	case nil:
		return "null"
	case bool:
		return "bool"
	case json.Number:
		if strings.ContainsAny(t.String(), ".eE") {
			return "real"
		}
		return "int"
	case float64:
		if t == float64(int64(t)) {
			return "int"
		}
		return "real"
	case string:
		return "text"
	default:
		return "json"
	}
}

func mergeJSONKind(a, b string) string {
	if a == b || b == "null" {
		return a
	}
	if a == "null" {
		return b
	}
	if (a == "int" && b == "real") || (a == "real" && b == "int") {
		return "real"
	}
	return "json"
}

func createDecodedJSONTable(db *gorm.DB, table string, columns []decodedColumn, child bool) error {
	defs := []string{}
	primaryColumn := "_row_index"
	indexColumn := "_master_id"
	if child {
		primaryColumn = "_row_index"
		indexColumn = "_parent_master_id"
		defs = append(defs,
			quoteIdent("_parent_row_index")+" INTEGER NOT NULL",
			quoteIdent("_item_index")+" INTEGER NOT NULL",
		)
	}
	defs = append(defs,
		quoteIdent(primaryColumn)+" INTEGER NOT NULL",
		quoteIdent(indexColumn)+" TEXT",
	)
	for _, col := range columns {
		if col.SQLName == primaryColumn || col.SQLName == indexColumn || col.SQLName == "_parent_row_index" || col.SQLName == "_item_index" {
			continue
		}
		defs = append(defs, quoteIdent(col.SQLName)+" "+sqlTypeForKind(db.Dialector.Name(), col.Kind))
	}
	sql := fmt.Sprintf("CREATE TABLE %s (%s, PRIMARY KEY (%s))",
		quoteIdent(table),
		strings.Join(defs, ", "),
		quoteIdent(primaryColumn),
	)
	if err := db.Exec(sql).Error; err != nil {
		return err
	}
	if err := db.Exec(fmt.Sprintf("CREATE INDEX %s ON %s (%s)",
		quoteIdent(limitedIdentifier("idx_"+table+"_"+indexColumn, 64)),
		quoteIdent(table),
		quoteIdent(indexColumn),
	)).Error; err != nil {
		return err
	}
	if child {
		if err := db.Exec(fmt.Sprintf("CREATE INDEX %s ON %s (%s)",
			quoteIdent(limitedIdentifier("idx_"+table+"_parent_row", 64)),
			quoteIdent(table),
			quoteIdent("_parent_row_index"),
		)).Error; err != nil {
			return err
		}
	}
	for _, col := range columns {
		if !isIndexableMasterColumn(col) {
			continue
		}
		if err := db.Exec(fmt.Sprintf("CREATE INDEX %s ON %s (%s)",
			quoteIdent(limitedIdentifier("idx_"+table+"_"+col.SQLName, 64)),
			quoteIdent(table),
			quoteIdent(col.SQLName),
		)).Error; err != nil {
			return err
		}
	}
	return nil
}

func insertDecodedJSONRow(db *gorm.DB, table string, columns []decodedColumn, index int, row map[string]any, indexColumn string, masterIDColumn string) error {
	names := []string{indexColumn, masterIDColumn}
	masterID := extractMasterID(row)
	if masterIDColumn == "_parent_master_id" {
		masterID = jsonValueString(row["_parent_master_id"])
	}
	values := []any{index, masterID}
	placeholders := []string{"?", "?"}
	if v, ok := row["_parent_row_index"]; ok && indexColumn != "_parent_row_index" {
		names = append(names, "_parent_row_index")
		values = append(values, v)
		placeholders = append(placeholders, "?")
	}
	if v, ok := row["_item_index"]; ok && indexColumn != "_item_index" {
		names = append(names, "_item_index")
		values = append(values, v)
		placeholders = append(placeholders, "?")
	}
	for _, col := range columns {
		if col.SQLName == indexColumn || col.SQLName == masterIDColumn || col.SQLName == "_parent_row_index" || col.SQLName == "_item_index" {
			continue
		}
		names = append(names, col.SQLName)
		values = append(values, sqlValue(row[col.JSONKey], col.Kind))
		placeholders = append(placeholders, "?")
	}
	quoted := make([]string, 0, len(names))
	for _, name := range names {
		quoted = append(quoted, quoteIdent(name))
	}
	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		quoteIdent(table),
		strings.Join(quoted, ", "),
		strings.Join(placeholders, ", "),
	)
	return db.Exec(sql, values...).Error
}

func isIndexableMasterColumn(col decodedColumn) bool {
	name := strings.ToLower(col.SQLName)
	if name == "id" || strings.HasSuffix(name, "_id") || strings.HasSuffix(name, "master_id") {
		return col.Kind == "int" || col.Kind == "real" || col.Kind == "bool" || col.Kind == "text"
	}
	return false
}

func sqlTypeForKind(dialect string, kind string) string {
	switch kind {
	case "int":
		return "BIGINT"
	case "real":
		return "DOUBLE"
	case "bool":
		if dialect == "mysql" {
			return "TINYINT(1)"
		}
		return "BOOLEAN"
	default:
		if dialect == "mysql" {
			return "LONGTEXT"
		}
		return "TEXT"
	}
}

func sqlValue(v any, kind string) any {
	if v == nil {
		return nil
	}
	switch kind {
	case "int", "real", "bool", "text":
		switch t := v.(type) {
		case json.Number:
			return t.String()
		case string, bool, float64:
			return t
		default:
			body, _ := json.Marshal(t)
			return string(body)
		}
	default:
		body, _ := json.Marshal(v)
		return string(body)
	}
}

func safeIdentifier(s string) string {
	s = strings.TrimSuffix(strings.TrimSpace(s), ".json")
	s = strings.ReplaceAll(s, "\\", "/")
	var b strings.Builder
	lastUnderscore := false
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			if b.Len() > 0 && !lastUnderscore {
				b.WriteByte('_')
			}
			b.WriteRune(r + ('a' - 'A'))
			lastUnderscore = false
			continue
		}
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
			b.WriteRune(r)
			lastUnderscore = false
			continue
		}
		if !lastUnderscore && b.Len() > 0 {
			b.WriteByte('_')
			lastUnderscore = true
		}
	}
	out := strings.Trim(b.String(), "_")
	if out == "" {
		return "value"
	}
	if out[0] >= '0' && out[0] <= '9' {
		out = "m_" + out
	}
	return out
}

func limitedIdentifier(s string, max int) string {
	if max <= 0 || len(s) <= max {
		return s
	}
	sum := sha256.Sum256([]byte(s))
	suffix := "_" + hex.EncodeToString(sum[:])[:12]
	head := max - len(suffix)
	if head <= 0 {
		return hex.EncodeToString(sum[:])[:max]
	}
	return strings.TrimRight(s[:head], "_") + suffix
}

func quoteIdent(s string) string {
	return "`" + strings.ReplaceAll(s, "`", "``") + "`"
}

func jsonValueString(v any) string {
	switch t := v.(type) {
	case nil:
		return ""
	case string:
		return strings.TrimSpace(t)
	case json.Number:
		return t.String()
	case float64:
		return fmt.Sprintf("%.0f", t)
	default:
		return fmt.Sprintf("%v", t)
	}
}
