package apiwire

import (
	"io"
	"math"
	"time"

	"github.com/vmihailenco/msgpack/v5"
)

func ReadRequestBody(r io.Reader) ([]byte, error) { return io.ReadAll(r) }

func DecodeRequestArray(body []byte) ([]any, error) {
	raw, err := DecodeMaybeLZ4BlockArray(body)
	if err != nil {
		return nil, err
	}
	if len(raw) == 0 {
		return []any{}, nil
	}
	var arr []any
	if err := msgpack.Unmarshal(raw, &arr); err != nil {
		return nil, err
	}
	return NormalizeMsgpackValue(arr).([]any), nil
}

func NormalizeMsgpackValue(value any) any {
	switch v := value.(type) {
	case []any:
		out := make([]any, len(v))
		for i, item := range v {
			out[i] = NormalizeMsgpackValue(item)
		}
		return out
	case map[string]any:
		out := make(map[string]any, len(v))
		for key, item := range v {
			out[key] = NormalizeMsgpackValue(item)
		}
		return out
	case map[any]any:
		out := make(map[any]any, len(v))
		for key, item := range v {
			out[NormalizeMsgpackValue(key)] = NormalizeMsgpackValue(item)
		}
		return out
	default:
		return v
	}
}

func ValueAt(arr []any, idx int) any {
	if idx < 0 || idx >= len(arr) {
		return nil
	}
	return arr[idx]
}

func Int64At(arr []any, idx int) int64     { return Int64Value(ValueAt(arr, idx)) }
func IntAt(arr []any, idx int) int         { return int(Int64At(arr, idx)) }
func StringAt(arr []any, idx int) string   { v, _ := ValueAt(arr, idx).(string); return v }
func BoolAt(arr []any, idx int) bool       { v, _ := ValueAt(arr, idx).(bool); return v }
func Float64At(arr []any, idx int) float64 { return Float64Value(ValueAt(arr, idx)) }

func TimeAt(arr []any, idx int) time.Time {
	v := ValueAt(arr, idx)
	switch t := v.(type) {
	case time.Time:
		return t
	case string:
		if parsed, ok := ParseTimeString(t); ok {
			return parsed
		}
	}
	return time.Time{}
}

func NullableInt64At(arr []any, idx int) *int64 {
	v := ValueAt(arr, idx)
	if v == nil {
		return nil
	}
	n := Int64Value(v)
	return &n
}

func Int64Value(v any) int64 {
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
		if n > uint64(math.MaxInt64) {
			return 0
		}
		return int64(n)
	case float32:
		return int64(n)
	case float64:
		return int64(n)
	default:
		return 0
	}
}

func Float64Value(v any) float64 {
	switch n := v.(type) {
	case float64:
		return n
	case float32:
		return float64(n)
	case int:
		return float64(n)
	case int8:
		return float64(n)
	case int16:
		return float64(n)
	case int32:
		return float64(n)
	case int64:
		return float64(n)
	case uint:
		return float64(n)
	case uint8:
		return float64(n)
	case uint16:
		return float64(n)
	case uint32:
		return float64(n)
	case uint64:
		return float64(n)
	default:
		return 0
	}
}

func ParseTimeString(value string) (time.Time, bool) {
	layouts := []string{"2006-01-02 15:04:05.999999999", "2006-01-02 15:04:05", time.RFC3339Nano, time.RFC3339}
	for _, layout := range layouts {
		if t, err := time.ParseInLocation(layout, value, time.Local); err == nil {
			return t, true
		}
	}
	return time.Time{}, false
}
