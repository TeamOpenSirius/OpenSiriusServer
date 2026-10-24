package apiwire

import "testing"

func TestDecodeRequestArrayPlain(t *testing.T) {
	raw := mustMsgpackMarshal(t, []any{"abc", int64(123), true})
	arr, err := DecodeRequestArray(raw)
	if err != nil {
		t.Fatal(err)
	}
	if StringAt(arr, 0) != "abc" || Int64At(arr, 1) != 123 || !BoolAt(arr, 2) {
		t.Fatalf("unexpected decoded values: %#v", arr)
	}
}

func TestDecodeRequestArrayLZ4(t *testing.T) {
	raw := mustMsgpackMarshal(t, []any{"abc", int64(123)})
	packed, err := EncodeLZ4BlockArray(raw, 1024)
	if err != nil {
		t.Fatal(err)
	}
	arr, err := DecodeRequestArray(packed)
	if err != nil {
		t.Fatal(err)
	}
	if StringAt(arr, 0) != "abc" || Int64At(arr, 1) != 123 {
		t.Fatalf("unexpected decoded values: %#v", arr)
	}
}

func TestScalarHelpersOutOfRange(t *testing.T) {
	var arr []any
	if Int64At(arr, 10) != 0 || StringAt(arr, 10) != "" || BoolAt(arr, 10) || Float64At(arr, 10) != 0 || ValueAt(arr, 10) != nil {
		t.Fatalf("out of range helper returned non-zero")
	}
}

func TestNormalizeMsgpackValueKeepsNilSlot(t *testing.T) {
	arr := NormalizeMsgpackValue([]any{int64(1), nil, int64(3)}).([]any)
	if len(arr) != 3 || arr[1] != nil {
		t.Fatalf("nil slot not preserved: %#v", arr)
	}
}
