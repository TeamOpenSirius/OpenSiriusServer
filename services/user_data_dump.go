package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/vmihailenco/msgpack/v5"
)

type UserDataDumpSummary struct {
	TotalSize               int                 `json:"total_size"`
	ObjectCount             int                 `json:"object_count"`
	UnionCounts             map[string]int      `json:"union_counts"`
	Objects                 []UserObjectSummary `json:"objects"`
	Warnings                []string            `json:"warnings"`
	PayloadCompressed       bool                `json:"payload_compressed"`
	PayloadFirstExtType     int                 `json:"payload_first_ext_type"`
	PayloadBlockCount       int                 `json:"payload_block_count"`
	PayloadDecompressedSize int                 `json:"payload_decompressed_size,omitempty"`
}

type UserObjectSummary struct {
	Index           int            `json:"index"`
	UnionID         int64          `json:"union_id"`
	PrimaryID       any            `json:"primary_id,omitempty"`
	DataLen         int            `json:"data_len"`
	ObjectKey       string         `json:"object_key"`
	FieldTypes      []string       `json:"field_types,omitempty"`
	ImportantFields map[string]any `json:"important_fields,omitempty"`
}

func DumpUserDataResponse(dir string, body []byte, payload []any) (string, string, error) {
	if dir == "" {
		dir = "data/dumps"
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", "", err
	}

	stamp := time.Now().Format("20060102_150405")
	binPath := filepath.Join(dir, "api_data_user_seed_"+stamp+".bin")
	summaryPath := filepath.Join(dir, "api_data_user_seed_"+stamp+".summary.json")

	if err := os.WriteFile(binPath, body, 0644); err != nil {
		return "", "", err
	}

	summary := BuildUserDataSummary(len(body), body, payload)
	rawSummary, err := json.MarshalIndent(summary, "", "  ")
	if err != nil {
		return binPath, "", err
	}
	if err := os.WriteFile(summaryPath, rawSummary, 0644); err != nil {
		return binPath, "", err
	}

	return binPath, summaryPath, nil
}

func BuildUserDataSummary(totalSize int, body []byte, payload []any) UserDataDumpSummary {
	s := UserDataDumpSummary{
		TotalSize:   totalSize,
		ObjectCount: len(payload),
		UnionCounts: make(map[string]int),
		Objects:     make([]UserObjectSummary, 0, len(payload)),
		Warnings:    []string{},
	}

	seenKeys := map[string]int{}
	unionSeen := map[int64]int{}
	userID := int64(0)
	userTutorial := int64(-1)
	profileID, preferenceID, currencyID, restrictionID, userBonusID := int64(0), int64(0), int64(0), int64(0), int64(0)
	parties := map[int64]bool{}
	characters := map[int64]bool{}
	partySlotRefs := []struct {
		slotID      any
		partyID     int64
		characterID int64
	}{}
	characterMasterIDs := map[int64]bool{}
	characterBaseMasterIDs := map[int64]bool{}
	gameHints := map[int64]bool{}
	homeDisplayRefs := map[string]int64{}

	for i, raw := range payload {
		unionID, data, ok := splitDataObject(raw)
		if !ok {
			s.Warnings = append(s.Warnings, fmt.Sprintf("object %d is not [unionId,dataArray]", i))
			continue
		}

		unionSeen[unionID]++
		s.UnionCounts[strconv.FormatInt(unionID, 10)]++

		key := objectKey(unionID, data)
		if prev, exists := seenKeys[key]; exists {
			s.Warnings = append(s.Warnings, fmt.Sprintf("duplicate object key %s at index %d and %d", key, prev, i))
		}
		seenKeys[key] = i

		obj := UserObjectSummary{
			Index:           i,
			UnionID:         unionID,
			PrimaryID:       valueAt(data, 0),
			DataLen:         len(data),
			ObjectKey:       key,
			FieldTypes:      fieldTypes(data),
			ImportantFields: importantFields(unionID, data),
		}
		s.Objects = append(s.Objects, obj)

		switch unionID {
		case UnionUser:
			userID = int64At(data, 0)
			userTutorial = int64At(data, 15)
			if userID == 0 {
				s.Warnings = append(s.Warnings, "User.Id is zero")
			}
			if userTutorial != 99 {
				s.Warnings = append(s.Warnings, fmt.Sprintf("User.TutorialStatus is %d, expected 99", userTutorial))
			}
			if stringAt(data, 13) == "" {
				s.Warnings = append(s.Warnings, "User.HashUserId is empty")
			}
			for _, idx := range []int{4, 12, 17} {
				if timeAt(data, idx).IsZero() {
					s.Warnings = append(s.Warnings, fmt.Sprintf("User DateTime key %d is zero", idx))
				}
			}
		case UnionUserProfile:
			profileID = int64At(data, 0)
		case UnionUserPreference:
			preferenceID = int64At(data, 0)
		case UnionCurrency:
			currencyID = int64At(data, 0)
		case UnionRestriction:
			restrictionID = int64At(data, 0)
		case UnionUserBonus:
			userBonusID = int64At(data, 0)
		case UnionGameHint:
			gameHints[int64At(data, 1)] = boolAt(data, 2)
		case UnionCharacter:
			characters[int64At(data, 0)] = true
			characterMasterIDs[int64At(data, 1)] = true
		case UnionCharacterBase:
			characterBaseMasterIDs[int64At(data, 1)] = true
		case UnionParty:
			parties[int64At(data, 0)] = true
		case UnionPartySlot:
			partySlotRefs = append(partySlotRefs, struct {
				slotID      any
				partyID     int64
				characterID int64
			}{slotID: valueAt(data, 0), partyID: int64At(data, 1), characterID: int64At(data, 3)})
		case UnionHomeDisplayPreference:
			homeDisplayRefs["HomeCharacterBaseMasterId"] = ptrInt64Value(valueAt(data, 1))
			homeDisplayRefs["MemberCharacterBaseMasterId"] = ptrInt64Value(valueAt(data, 3))
			homeDisplayRefs["StoryCharacterBaseMasterId"] = ptrInt64Value(valueAt(data, 4))
			homeDisplayRefs["ShopCharacterBaseMasterId"] = ptrInt64Value(valueAt(data, 5))
			homeDisplayRefs["IllustCharacterMasterId"] = int64At(data, 11)
			homeDisplayRefs["LoginBonusCharacterBaseMasterId"] = ptrInt64Value(valueAt(data, 14))
		}
	}

	for _, required := range []int64{0, 1, 2, 3, 4, 5, 6, 7, 66, 67, 128, 148} {
		if unionSeen[required] == 0 {
			s.Warnings = append(s.Warnings, fmt.Sprintf("missing required union id %d", required))
		}
	}
	if unionSeen[UnionCharacter] < 1 {
		s.Warnings = append(s.Warnings, "Character count is zero")
	}
	if unionSeen[UnionCharacterBase] < 1 {
		s.Warnings = append(s.Warnings, "CharacterBase count is zero")
	}
	if unionSeen[UnionParty] < 1 {
		s.Warnings = append(s.Warnings, "Party count is zero")
	}
	if unionSeen[UnionPartySlot] < 1 {
		s.Warnings = append(s.Warnings, "PartySlot count is zero")
	}
	if unionSeen[UnionGameHint] < 7 {
		s.Warnings = append(s.Warnings, fmt.Sprintf("GameHint count is %d, expected at least 7", unionSeen[UnionGameHint]))
	}
	for _, category := range []int64{19, 21, 29, 32, 33, 34, 43} {
		if !gameHints[category] {
			s.Warnings = append(s.Warnings, fmt.Sprintf("GameHint pageCategory %d missing or not read", category))
		}
	}
	accountID := firstNonZeroInt64(profileID, preferenceID, currencyID, restrictionID, userBonusID)
	for name, id := range map[string]int64{
		"UserProfile.Id":    profileID,
		"UserPreference.Id": preferenceID,
		"Currency.Id":       currencyID,
		"Restriction.Id":    restrictionID,
		"UserBonus.Id":      userBonusID,
	} {
		if accountID != 0 && id != 0 && id != accountID {
			s.Warnings = append(s.Warnings, fmt.Sprintf("%s=%d does not match account object id=%d", name, id, accountID))
		}
	}
	for name, id := range homeDisplayRefs {
		if id == 0 {
			continue
		}
		if name == "IllustCharacterMasterId" {
			if !characterMasterIDs[id] {
				s.Warnings = append(s.Warnings, fmt.Sprintf("HomeDisplayPreference.%s references missing CharacterMasterId %d", name, id))
			}
			continue
		}
		if !characterBaseMasterIDs[id] {
			s.Warnings = append(s.Warnings, fmt.Sprintf("HomeDisplayPreference.%s references missing CharacterBaseMasterId %d", name, id))
		}
	}
	for _, ref := range partySlotRefs {
		if !parties[ref.partyID] {
			s.Warnings = append(s.Warnings, fmt.Sprintf("PartySlot %v references missing PartyId %d", ref.slotID, ref.partyID))
		}
		if !characters[ref.characterID] {
			s.Warnings = append(s.Warnings, fmt.Sprintf("PartySlot %v references missing CharacterId %d", ref.slotID, ref.characterID))
		}
	}
	analyzeEnvelopeCompression(body, &s)
	return s
}

func analyzeEnvelopeCompression(body []byte, s *UserDataDumpSummary) {
	if len(body) == 0 {
		return
	}
	payloadOffset, err := skipMsgpackObjects(body, 1)
	if err != nil || payloadOffset >= len(body) {
		return
	}
	payloadBytes := body[payloadOffset:]

	compressed, extType, blockCount, decompSize := detectLZ4BlockArray(payloadBytes)
	s.PayloadCompressed = compressed
	s.PayloadFirstExtType = extType
	s.PayloadBlockCount = blockCount
	s.PayloadDecompressedSize = decompSize
}

func skipMsgpackObjects(body []byte, count int) (int, error) {
	r := newMsgpackReader(body)
	for i := 0; i < count; i++ {
		if err := r.skipValue(); err != nil {
			return 0, err
		}
	}
	return r.pos, nil
}

func detectLZ4BlockArray(data []byte) (compressed bool, extType int, blockCount int, decompSize int) {
	r := newMsgpackReader(data)
	arrayLen, ok := r.readArrayLen()
	if !ok || arrayLen < 2 {
		return false, 0, 0, 0
	}
	ext, ok := r.readExt()
	if !ok || ext.Type != 98 {
		return false, 0, 0, 0
	}
	extType = int(ext.Type)
	sizes, err := decodeBlockSizes(ext.Data)
	if err != nil || len(sizes) == 0 || arrayLen != len(sizes)+1 {
		return true, extType, 0, 0
	}
	blockCount = len(sizes)
	for _, sz := range sizes {
		decompSize += sz
	}
	return true, extType, blockCount, decompSize
}

type msgpackReader struct {
	data []byte
	pos  int
}

func newMsgpackReader(data []byte) *msgpackReader {
	return &msgpackReader{data: data}
}

func (r *msgpackReader) readByte() (byte, bool) {
	if r.pos >= len(r.data) {
		return 0, false
	}
	b := r.data[r.pos]
	r.pos++
	return b, true
}

func (r *msgpackReader) readN(n int) ([]byte, bool) {
	if r.pos+n > len(r.data) {
		return nil, false
	}
	d := r.data[r.pos : r.pos+n]
	r.pos += n
	return d, true
}

func (r *msgpackReader) readArrayLen() (int, bool) {
	b, ok := r.readByte()
	if !ok {
		return 0, false
	}
	if b >= 0x90 && b <= 0x9f {
		return int(b & 0x0f), true
	}
	switch b {
	case 0xdc:
		d, ok := r.readN(2)
		if !ok {
			return 0, false
		}
		return int(uint16(d[0])<<8 | uint16(d[1])), true
	case 0xdd:
		d, ok := r.readN(4)
		if !ok {
			return 0, false
		}
		return int(uint32(d[0])<<24 | uint32(d[1])<<16 | uint32(d[2])<<8 | uint32(d[3])), true
	default:
		r.pos--
		return 0, false
	}
}

type simpleExt struct {
	Type int8
	Data []byte
}

func (r *msgpackReader) readExt() (simpleExt, bool) {
	b, ok := r.readByte()
	if !ok {
		return simpleExt{}, false
	}
	var sz uint32
	switch b {
	case 0xd4:
		sz = 1
	case 0xd5:
		sz = 2
	case 0xd6:
		sz = 4
	case 0xd7:
		sz = 8
	case 0xd8:
		sz = 16
	case 0xc7:
		n, ok := r.readByte()
		if !ok {
			return simpleExt{}, false
		}
		sz = uint32(n)
	case 0xc8:
		d, ok := r.readN(2)
		if !ok {
			return simpleExt{}, false
		}
		sz = uint32(d[0])<<8 | uint32(d[1])
	case 0xc9:
		d, ok := r.readN(4)
		if !ok {
			return simpleExt{}, false
		}
		sz = uint32(d[0])<<24 | uint32(d[1])<<16 | uint32(d[2])<<8 | uint32(d[3])
	default:
		r.pos--
		return simpleExt{}, false
	}
	typeByte, ok := r.readByte()
	if !ok {
		return simpleExt{}, false
	}
	data, ok := r.readN(int(sz))
	if !ok {
		return simpleExt{}, false
	}
	return simpleExt{Type: int8(typeByte), Data: data}, true
}

func (r *msgpackReader) skipValue() error {
	b, ok := r.readByte()
	if !ok {
		return fmt.Errorf("unexpected eof at pos %d", r.pos)
	}
	switch {
	case b <= 0x7f: // positive fixint
		return nil
	case b >= 0x80 && b <= 0x8f: // fixmap
		_ = int(b & 0x0f) // count
		for i := 0; i < int(b&0x0f); i++ {
			if err := r.skipValue(); err != nil {
				return err
			}
			if err := r.skipValue(); err != nil {
				return err
			}
		}
		return nil
	case b >= 0x90 && b <= 0x9f: // fixarray
		cnt := int(b & 0x0f)
		for i := 0; i < cnt; i++ {
			if err := r.skipValue(); err != nil {
				return err
			}
		}
		return nil
	case b >= 0xa0 && b <= 0xbf: // fixstr
		_, ok := r.readN(int(b & 0x1f))
		if !ok {
			return fmt.Errorf("truncated fixstr")
		}
		return nil
	case b >= 0xe0: // negative fixint
		return nil
	case b == 0xc0: // nil
		return nil
	case b == 0xc2: // false
		return nil
	case b == 0xc3: // true
		return nil
	case b == 0xc4: // bin8
		n, ok := r.readByte()
		if !ok {
			return fmt.Errorf("truncated bin8")
		}
		_, ok = r.readN(int(n))
		return nil
	case b == 0xc5: // bin16
		d, _ := r.readN(2)
		n := int(d[0])<<8 | int(d[1])
		_, _ = r.readN(n)
		return nil
	case b == 0xc6: // bin32
		d, _ := r.readN(4)
		n := int(d[0])<<24 | int(d[1])<<16 | int(d[2])<<8 | int(d[3])
		_, _ = r.readN(n)
		return nil
	case b == 0xc7: // ext8
		n, _ := r.readByte()
		r.pos++ // type byte
		_, _ = r.readN(int(n))
		return nil
	case b == 0xc8: // ext16
		d, _ := r.readN(2)
		r.pos++ // type byte
		_, _ = r.readN(int(d[0])<<8 | int(d[1]))
		return nil
	case b == 0xc9: // ext32
		d, _ := r.readN(4)
		r.pos++ // type byte
		_, _ = r.readN(int(d[0])<<24 | int(d[1])<<16 | int(d[2])<<8 | int(d[3]))
		return nil
	case b == 0xca: // float32
		_, _ = r.readN(4)
		return nil
	case b == 0xcb: // float64
		_, _ = r.readN(8)
		return nil
	case b == 0xcc: // uint8
		_, _ = r.readByte()
		return nil
	case b == 0xcd: // uint16
		_, _ = r.readN(2)
		return nil
	case b == 0xce: // uint32
		_, _ = r.readN(4)
		return nil
	case b == 0xcf: // uint64
		_, _ = r.readN(8)
		return nil
	case b == 0xd0: // int8
		_, _ = r.readByte()
		return nil
	case b == 0xd1: // int16
		_, _ = r.readN(2)
		return nil
	case b == 0xd2: // int32
		_, _ = r.readN(4)
		return nil
	case b == 0xd3: // int64
		_, _ = r.readN(8)
		return nil
	case b == 0xd4, b == 0xd5, b == 0xd6, b == 0xd7, b == 0xd8: // fixext
		sizes := map[byte]int{0xd4: 1, 0xd5: 2, 0xd6: 4, 0xd7: 8, 0xd8: 16}
		r.pos++ // type byte
		_, _ = r.readN(sizes[b])
		return nil
	case b == 0xd9: // str8
		n, _ := r.readByte()
		_, _ = r.readN(int(n))
		return nil
	case b == 0xda: // str16
		d, _ := r.readN(2)
		_, _ = r.readN(int(d[0])<<8 | int(d[1]))
		return nil
	case b == 0xdb: // str32
		d, _ := r.readN(4)
		_, _ = r.readN(int(d[0])<<24 | int(d[1])<<16 | int(d[2])<<8 | int(d[3]))
		return nil
	case b == 0xdc: // array16
		cntByte, _ := r.readN(2)
		cnt := int(cntByte[0])<<8 | int(cntByte[1])
		for i := 0; i < cnt; i++ {
			if err := r.skipValue(); err != nil {
				return err
			}
		}
		return nil
	case b == 0xdd: // array32
		cntByte, _ := r.readN(4)
		cnt := int(cntByte[0])<<24 | int(cntByte[1])<<16 | int(cntByte[2])<<8 | int(cntByte[3])
		for i := 0; i < cnt; i++ {
			if err := r.skipValue(); err != nil {
				return err
			}
		}
		return nil
	case b == 0xde: // map16
		cntByte, _ := r.readN(2)
		cnt := int(cntByte[0])<<8 | int(cntByte[1])
		for i := 0; i < cnt; i++ {
			if err := r.skipValue(); err != nil {
				return err
			}
			if err := r.skipValue(); err != nil {
				return err
			}
		}
		return nil
	case b == 0xdf: // map32
		cntByte, _ := r.readN(4)
		cnt := int(cntByte[0])<<24 | int(cntByte[1])<<16 | int(cntByte[2])<<8 | int(cntByte[3])
		for i := 0; i < cnt; i++ {
			if err := r.skipValue(); err != nil {
				return err
			}
			if err := r.skipValue(); err != nil {
				return err
			}
		}
		return nil
	default:
		return fmt.Errorf("unknown msgpack tag 0x%02x at pos %d", b, r.pos-1)
	}
}

func splitDataObject(raw any) (int64, []any, bool) {
	obj, ok := raw.([]any)
	if !ok || len(obj) < 2 {
		return 0, nil, false
	}
	data, ok := obj[1].([]any)
	if !ok {
		return 0, nil, false
	}
	return int64Value(obj[0]), data, true
}

func objectKey(unionID int64, data []any) string {
	switch unionID {
	case UnionUser:
		return "0"
	case UnionGameHint:
		return fmt.Sprintf("66:%d", int64At(data, 1))
	case UnionViewedShop:
		return fmt.Sprintf("65:%d", int64At(data, 3))
	default:
		return fmt.Sprintf("%d:%d", unionID, int64At(data, 0))
	}
}

func importantFields(unionID int64, data []any) map[string]any {
	switch unionID {
	case UnionUser:
		return map[string]any{"id": valueAt(data, 0), "tutorialStatus": valueAt(data, 15), "playerRank": valueAt(data, 1), "stamina": valueAt(data, 3), "coin": valueAt(data, 7), "freeJewel": valueAt(data, 6), "paidJewel": valueAt(data, 5)}
	case UnionUserProfile:
		return map[string]any{"id": valueAt(data, 0), "name": valueAt(data, 1)}
	case UnionGameHint:
		return map[string]any{"id": valueAt(data, 0), "pageCategory": valueAt(data, 1), "hasAlreadyRead": valueAt(data, 2)}
	case UnionCharacter:
		return map[string]any{"id": valueAt(data, 0), "characterMasterId": valueAt(data, 1), "level": valueAt(data, 2), "characterBaseId": valueAt(data, 6)}
	case UnionCharacterBase:
		return map[string]any{"id": valueAt(data, 0), "characterBaseMasterId": valueAt(data, 1)}
	case UnionParty:
		return map[string]any{"id": valueAt(data, 0), "name": valueAt(data, 2)}
	case UnionPartySlot:
		return map[string]any{"id": valueAt(data, 0), "partyId": valueAt(data, 1), "position": valueAt(data, 2), "characterId": valueAt(data, 3)}
	case UnionCurrency:
		return map[string]any{"id": valueAt(data, 0), "coin": valueAt(data, 1), "freeJewel": valueAt(data, 2), "paidJewel": valueAt(data, 3)}
	case UnionViewedShop:
		return map[string]any{"id": valueAt(data, 0), "category": valueAt(data, 3)}
	default:
		return nil
	}
}

func valueAt(data []any, idx int) any {
	if idx < 0 || idx >= len(data) {
		return nil
	}
	return data[idx]
}

func int64At(data []any, idx int) int64 {
	return int64Value(valueAt(data, idx))
}

func int64Value(v any) int64 {
	switch n := v.(type) {
	case int:
		return int64(n)
	case int8:
		return int64(n)
	case int16:
		return int64(n)
	case int32:
		return int64(n)
	case int64:
		return n
	case uint:
		return int64(n)
	case uint8:
		return int64(n)
	case uint16:
		return int64(n)
	case uint32:
		return int64(n)
	case uint64:
		if n > uint64(^uint(0)>>1) {
			return 0
		}
		return int64(n)
	default:
		return 0
	}
}

func ptrInt64Value(v any) int64 {
	switch p := v.(type) {
	case *int64:
		if p == nil {
			return 0
		}
		return *p
	default:
		return int64Value(v)
	}
}

func boolAt(data []any, idx int) bool {
	v, _ := valueAt(data, idx).(bool)
	return v
}

func stringAt(data []any, idx int) string {
	v, _ := valueAt(data, idx).(string)
	return v
}

func timeAt(data []any, idx int) time.Time {
	v, _ := valueAt(data, idx).(time.Time)
	return v
}

func fieldTypes(data []any) []string {
	types := make([]string, len(data))
	for i, v := range data {
		types[i] = goTypeName(v)
	}
	return types
}

func goTypeName(v any) string {
	if v == nil {
		return "nil"
	}
	switch v.(type) {
	case bool:
		return "bool"
	case int64:
		return "int64"
	case int, int8, int16, int32:
		return "int"
	case uint, uint8, uint16, uint32, uint64:
		return "uint"
	case float32:
		return "float32"
	case float64:
		return "float64"
	case string:
		return "string"
	case time.Time:
		return "time.Time"
	case []any:
		return "[]any"
	case map[string]any:
		return "map"
	default:
		return fmt.Sprintf("%T", v)
	}
}

func firstNonZeroInt64(values ...int64) int64 {
	for _, value := range values {
		if value != 0 {
			return value
		}
	}
	return 0
}

func decodeBlockSizes(payload []byte) ([]int, error) {
	r := bytes.NewReader(payload)
	dec := msgpack.NewDecoder(r)
	var sizes []int
	for r.Len() > 0 {
		v, err := dec.DecodeInt64()
		if err != nil {
			return nil, err
		}
		if v <= 0 {
			return nil, fmt.Errorf("invalid lz4 block size: %d", v)
		}
		sizes = append(sizes, int(v))
	}
	return sizes, nil
}
