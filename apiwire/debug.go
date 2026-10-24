package apiwire

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/vmihailenco/msgpack/v5"
)

type MsgpackDebugSummary struct {
	OriginalSize     int
	DecodedSize      int
	WasLZ4BlockArray bool
	TopLevelKind     string
	ArrayLen         int
	FieldTypes       []string
}

func SummarizeRequestBody(body []byte) MsgpackDebugSummary {
	s := MsgpackDebugSummary{OriginalSize: len(body), DecodedSize: len(body)}
	decoded, err := DecodeMaybeLZ4BlockArray(body)
	if err != nil {
		return s
	}
	s.DecodedSize = len(decoded)
	s.WasLZ4BlockArray = !bytes.Equal(decoded, body)
	if len(decoded) == 0 {
		return s
	}
	var v any
	if err := msgpack.Unmarshal(decoded, &v); err != nil {
		return s
	}
	v = NormalizeMsgpackValue(v)
	s.TopLevelKind = kindName(v)
	if arr, ok := v.([]any); ok {
		s.ArrayLen = len(arr)
		s.FieldTypes = make([]string, len(arr))
		for i, item := range arr {
			s.FieldTypes[i] = kindName(item)
		}
	}
	return s
}

func kindName(v any) string {
	if v == nil {
		return "nil"
	}
	if _, ok := v.([]any); ok {
		return "array"
	}
	return fmt.Sprintf("%s", reflect.TypeOf(v))
}
