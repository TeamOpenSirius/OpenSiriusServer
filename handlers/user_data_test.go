package handlers

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vmihailenco/msgpack/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"ymst_server/apiwire"
	"ymst_server/core/config"
	"ymst_server/core/database"
	"ymst_server/middleware"
	"ymst_server/repositories"
)

func TestUserDataHandlerReturnsFiveObjectEnvelopeFromGeneratedRawState(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, err := gorm.Open(sqlite.Open("file:user_data_handler_"+time.Now().Format("150405.000000000")+"?mode=memory&cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := database.AutoMigrate(db); err != nil {
		t.Fatal(err)
	}
	users := repositories.NewUserRepository(db, "").WithUserKey("usr_handler")
	h := NewUserDataHandler(nil, config.GameConfig{}, config.DebugConfig{}, users)
	r := gin.New()
	r.Use(func(c *gin.Context) {
		c.Set(middleware.ContextUserKey, "usr_handler")
		c.Next()
	})
	h.RegisterRoutes(r)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/data/user", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	body := w.Body.Bytes()

	payloadEnd, payloadRaw, err := decodeEnvelopeWithCompressedPayload(body)
	if err != nil {
		t.Fatal(err)
	}

	dec := msgpack.NewDecoder(bytes.NewReader(body))
	var common []any
	if err := dec.Decode(&common); err != nil {
		t.Fatal(err)
	}
	_ = common

	dec = msgpack.NewDecoder(bytes.NewReader(body[payloadEnd:]))
	var updates []any
	if err := dec.Decode(&updates); err != nil {
		t.Fatal(err)
	}
	var deletes []any
	if err := dec.Decode(&deletes); err != nil {
		t.Fatal(err)
	}
	var effects []any
	if err := dec.Decode(&effects); err != nil {
		t.Fatal(err)
	}

	if len(common) != 0 || len(updates) != 0 || len(deletes) != 0 || len(effects) != 0 {
		t.Fatalf("unexpected envelope sizes: common=%d updates=%d deletes=%d effects=%d", len(common), len(updates), len(deletes), len(effects))
	}
	if len(payloadRaw) == 0 {
		t.Fatal("empty payload")
	}
	counts := map[int64]int{}
	for _, raw := range payloadRaw {
		obj, ok := raw.([]any)
		if !ok || len(obj) < 2 {
			continue
		}
		counts[apiwire.Int64Value(obj[0])]++
	}
	if counts[repositories.UnionUser] != 1 || counts[repositories.UnionCharacter] < 5 || counts[repositories.UnionCurrency] != 1 {
		t.Fatalf("missing required unions: %#v", counts)
	}
}

func decodeEnvelopeWithCompressedPayload(body []byte) (payloadEndOffset int, payload []any, err error) {
	object1Start := msgpackEncodedLen(body, 0, 1)
	if object1Start < 0 {
		return 0, nil, fmt.Errorf("could not locate object1")
	}
	object1End, ok := msgpackSkipValue(body, object1Start)
	if !ok {
		return 0, nil, fmt.Errorf("could not parse object1 end")
	}
	compressedBytes := body[object1Start:object1End]
	decompressed, err := apiwire.DecodeMaybeLZ4BlockArray(compressedBytes)
	if err != nil {
		return 0, nil, fmt.Errorf("lz4 decode: %w", err)
	}
	var decoded []any
	if err := msgpack.NewDecoder(bytes.NewReader(decompressed)).Decode(&decoded); err != nil {
		return 0, nil, fmt.Errorf("decode decompressed payload: %w", err)
	}
	return object1End, decoded, nil
}

func TestUserDataPayloadAlwaysCompressed(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, err := gorm.Open(sqlite.Open("file:user_data_cmp_"+time.Now().Format("150405.000000000")+"?mode=memory&cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := database.AutoMigrate(db); err != nil {
		t.Fatal(err)
	}
	users := repositories.NewUserRepository(db, "").WithUserKey("usr_cmp")
	h := NewUserDataHandler(nil, config.GameConfig{}, config.DebugConfig{}, users)
	r := gin.New()
	r.Use(func(c *gin.Context) {
		c.Set(middleware.ContextUserKey, "usr_cmp")
		c.Next()
	})
	h.RegisterRoutes(r)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/data/user", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	body := w.Body.Bytes()

	object1Start := msgpackEncodedLen(body, 0, 1)
	if object1Start < 0 || object1Start >= len(body) {
		t.Fatal("could not locate payload in envelope")
	}
	object1End, ok := msgpackSkipValue(body, object1Start)
	if !ok {
		t.Fatal("could not parse object1")
	}

	dec := msgpack.NewDecoder(bytes.NewReader(body))
	var common []any
	if err := dec.Decode(&common); err != nil {
		t.Fatal("decode object0:", err)
	}
	if len(common) != 0 {
		t.Fatalf("object0 should be [], got len=%d", len(common))
	}

	dec = msgpack.NewDecoder(bytes.NewReader(body[object1End:]))
	var updates []any
	if err := dec.Decode(&updates); err != nil {
		t.Fatal("decode object2:", err)
	}
	if len(updates) != 0 {
		t.Fatalf("object2 should be [], got len=%d", len(updates))
	}

	var deletes []any
	if err := dec.Decode(&deletes); err != nil {
		t.Fatal("decode object3:", err)
	}
	if len(deletes) != 0 {
		t.Fatalf("object3 should be [], got len=%d", len(deletes))
	}

	var effects []any
	if err := dec.Decode(&effects); err != nil {
		t.Fatal("decode object4:", err)
	}
	if len(effects) != 0 {
		t.Fatalf("object4 should be [], got len=%d", len(effects))
	}

	compressedBytes := body[object1Start:object1End]
	extType := detectFirstExtType(compressedBytes)
	if extType != 98 {
		t.Fatalf("lz4 block array first ext type want 98, got %d", extType)
	}
	blockCount := detectBlockCount(compressedBytes)

	_, err = apiwire.DecodeMaybeLZ4BlockArray(compressedBytes)
	if err != nil {
		t.Fatalf("decodeMaybeLZ4BlockArray failed: %v", err)
	}

	t.Logf("payload compressed: first ext type = %d, block count = %d", extType, blockCount)
}

func msgpackEncodedLen(body []byte, startOffset int, targetIndex int) int {
	off := startOffset
	for i := 0; i <= targetIndex; i++ {
		if off >= len(body) {
			return -1
		}
		advance, ok := msgpackSkipValue(body, off)
		if !ok {
			return -1
		}
		if i == targetIndex {
			return off
		}
		off = advance
	}
	return off
}

func msgpackSkipValue(body []byte, pos int) (int, bool) {
	if pos >= len(body) {
		return 0, false
	}
	b := body[pos]
	pos++
	switch {
	case b <= 0x7f || (b >= 0xe0):
		return pos, true
	case b == 0xc0, b == 0xc2, b == 0xc3:
		return pos, true
	case b >= 0xa0 && b <= 0xbf:
		return pos + int(b&0x1f), true
	case b >= 0x90 && b <= 0x9f:
		return msgpackSkipNElements(body, pos, int(b&0x0f))
	case b == 0xdc:
		if pos+2 > len(body) {
			return 0, false
		}
		n := int(binary.BigEndian.Uint16(body[pos:]))
		return msgpackSkipNElements(body, pos+2, n)
	case b == 0xdd:
		if pos+4 > len(body) {
			return 0, false
		}
		n := int(binary.BigEndian.Uint32(body[pos:]))
		return msgpackSkipNElements(body, pos+4, n)
	case b == 0x80:
		return msgpackSkipNKeyValues(body, pos, 1)
	case b >= 0x81 && b <= 0x8f:
		return msgpackSkipNKeyValues(body, pos, int(b&0x0f))
	case b == 0xde:
		if pos+2 > len(body) {
			return 0, false
		}
		n := int(binary.BigEndian.Uint16(body[pos:]))
		return msgpackSkipNKeyValues(body, pos+2, n)
	case b == 0xdf:
		if pos+4 > len(body) {
			return 0, false
		}
		n := int(binary.BigEndian.Uint32(body[pos:]))
		return msgpackSkipNKeyValues(body, pos+4, n)
	case b == 0xd9:
		if pos >= len(body) {
			return 0, false
		}
		return pos + 1 + int(body[pos]), true
	case b == 0xda:
		if pos+2 > len(body) {
			return 0, false
		}
		return pos + 2 + int(binary.BigEndian.Uint16(body[pos:])), true
	case b == 0xdb:
		if pos+4 > len(body) {
			return 0, false
		}
		return pos + 4 + int(binary.BigEndian.Uint32(body[pos:])), true
	case b == 0xc4:
		if pos >= len(body) {
			return 0, false
		}
		return pos + 1 + int(body[pos]), true
	case b == 0xc5:
		if pos+2 > len(body) {
			return 0, false
		}
		return pos + 2 + int(binary.BigEndian.Uint16(body[pos:])), true
	case b == 0xc6:
		if pos+4 > len(body) {
			return 0, false
		}
		return pos + 4 + int(binary.BigEndian.Uint32(body[pos:])), true
	case b >= 0xc7 && b <= 0xc9:
		return msgpackSkipExt(body, pos, b)
	case b >= 0xd4 && b <= 0xd8:
		sizes := map[byte]int{0xd4: 1, 0xd5: 2, 0xd6: 4, 0xd7: 8, 0xd8: 16}
		return pos + 1 + sizes[b], true
	case b == 0xca:
		return pos + 4, true
	case b == 0xcb:
		return pos + 8, true
	case b == 0xcc:
		return pos + 1, true
	case b == 0xcd:
		return pos + 2, true
	case b == 0xce:
		return pos + 4, true
	case b == 0xcf:
		return pos + 8, true
	case b == 0xd0:
		return pos + 1, true
	case b == 0xd1:
		return pos + 2, true
	case b == 0xd2:
		return pos + 4, true
	case b == 0xd3:
		return pos + 8, true
	default:
		return 0, false
	}
}

func msgpackSkipNElements(body []byte, pos int, n int) (int, bool) {
	for i := 0; i < n; i++ {
		var ok bool
		pos, ok = msgpackSkipValue(body, pos)
		if !ok {
			return 0, false
		}
	}
	return pos, true
}

func msgpackSkipNKeyValues(body []byte, pos int, n int) (int, bool) {
	for i := 0; i < n*2; i++ {
		var ok bool
		pos, ok = msgpackSkipValue(body, pos)
		if !ok {
			return 0, false
		}
	}
	return pos, true
}

func msgpackSkipExt(body []byte, pos int, code byte) (int, bool) {
	var sz uint32
	switch code {
	case 0xc7:
		if pos >= len(body) {
			return 0, false
		}
		sz = uint32(body[pos])
		pos++
	case 0xc8:
		if pos+2 > len(body) {
			return 0, false
		}
		sz = uint32(binary.BigEndian.Uint16(body[pos:]))
		pos += 2
	case 0xc9:
		if pos+4 > len(body) {
			return 0, false
		}
		sz = binary.BigEndian.Uint32(body[pos:])
		pos += 4
	}
	pos++
	return pos + int(sz), true
}

func detectFirstExtType(data []byte) int {
	if len(data) < 3 {
		return -1
	}
	pos := 0
	b := data[pos]
	pos++
	switch {
	case b >= 0x90 && b <= 0x9f:
		pos += int(b&0x0f) * 0 // no advance for skip, just detect
		// actually skip the array elements to get to first element's ext
		// We know it's an array; first element is at pos
	case b == 0xdc:
		if pos+2 > len(data) {
			return -1
		}
		pos += 2
	case b == 0xdd:
		if pos+4 > len(data) {
			return -1
		}
		pos += 4
	default:
		return -1
	}
	if pos >= len(data) {
		return -1
	}
	return readExtType(data, pos)
}

func readExtType(data []byte, pos int) int {
	if pos >= len(data) {
		return -1
	}
	b := data[pos]
	pos++
	switch b {
	case 0xd4:
		if pos >= len(data) {
			return -1
		}
		return int(data[pos])
	case 0xd5:
		if pos+1 >= len(data) {
			return -1
		}
		return int(data[pos+1])
	case 0xd6:
		if pos+3 >= len(data) {
			return -1
		}
		return int(data[pos+3])
	case 0xd7:
		if pos+7 >= len(data) {
			return -1
		}
		return int(data[pos+7])
	case 0xd8:
		if pos+15 >= len(data) {
			return -1
		}
		return int(data[pos+15])
	case 0xc7, 0xc8, 0xc9:
		var sz int
		switch b {
		case 0xc7:
			if pos >= len(data) {
				return -1
			}
			sz = int(data[pos])
			pos++
		case 0xc8:
			if pos+2 > len(data) {
				return -1
			}
			sz = int(binary.BigEndian.Uint16(data[pos:]))
			pos += 2
		case 0xc9:
			if pos+4 > len(data) {
				return -1
			}
			sz = int(binary.BigEndian.Uint32(data[pos:]))
			pos += 4
		}
		if pos >= len(data) {
			return -1
		}
		_ = sz
		return int(data[pos])
	default:
		return -1
	}
}

func detectBlockCount(data []byte) int {
	if len(data) < 2 {
		return 0
	}
	pos := 0
	b := data[pos]
	pos++
	var arrayLen int
	switch {
	case b >= 0x90 && b <= 0x9f:
		arrayLen = int(b & 0x0f)
	case b == 0xdc:
		if pos+2 > len(data) {
			return 0
		}
		arrayLen = int(binary.BigEndian.Uint16(data[pos:]))
		pos += 2
	case b == 0xdd:
		if pos+4 > len(data) {
			return 0
		}
		arrayLen = int(binary.BigEndian.Uint32(data[pos:]))
		pos += 4
	default:
		return 0
	}
	if arrayLen < 1 {
		return 0
	}
	return arrayLen - 1
}
