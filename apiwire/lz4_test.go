package apiwire

import (
	"bytes"
	"testing"

	"github.com/vmihailenco/msgpack/v5"
)

func mustMsgpackMarshal(t *testing.T, v any) []byte {
	t.Helper()
	raw, err := msgpack.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	return raw
}

func TestEncodeLZ4BlockArrayRoundTripSmall(t *testing.T) {
	raw := mustMsgpackMarshal(t, []any{"abc", int64(123), true})
	packed, err := EncodeLZ4BlockArray(raw, 1024)
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := DecodeMaybeLZ4BlockArray(packed)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(raw, decoded) {
		t.Fatalf("decoded mismatch")
	}
}

func TestEncodeLZ4BlockArrayRoundTripMultiBlock(t *testing.T) {
	raw := bytes.Repeat([]byte("0123456789abcdef"), 20000)
	packed, err := EncodeLZ4BlockArray(raw, 4096)
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := DecodeMaybeLZ4BlockArray(packed)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(raw, decoded) {
		t.Fatalf("decoded mismatch")
	}
}

func TestDecodeMaybeLZ4BlockArrayPlainBody(t *testing.T) {
	raw := mustMsgpackMarshal(t, []any{"plain"})
	decoded, err := DecodeMaybeLZ4BlockArray(raw)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(raw, decoded) {
		t.Fatalf("plain body changed")
	}
}

func TestDecodeMaybeLZ4BlockArrayNon98Ext(t *testing.T) {
	var body bytes.Buffer
	writeMsgpackArrayLen(&body, 2)
	writeMsgpackExt(&body, 99, []byte{1})
	writeMsgpackBin(&body, []byte{2})
	decoded, err := DecodeMaybeLZ4BlockArray(body.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(body.Bytes(), decoded) {
		t.Fatalf("non-98 ext changed")
	}
}
