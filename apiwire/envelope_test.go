package apiwire

import (
	"bytes"
	"testing"

	"github.com/vmihailenco/msgpack/v5"
)

func TestEncodeResultWritesFiveObjects(t *testing.T) {
	body, err := EncodeResult([]any{"payload"}, nil, nil, nil, DefaultResponseOptions())
	if err != nil {
		t.Fatal(err)
	}
	dec := msgpack.NewDecoder(bytes.NewReader(body))
	var common []any
	var payload []any
	var updates []any
	var deletes []any
	var effects []any
	if err := dec.Decode(&common); err != nil {
		t.Fatal(err)
	}
	if err := dec.Decode(&payload); err != nil {
		t.Fatal(err)
	}
	if err := dec.Decode(&updates); err != nil {
		t.Fatal(err)
	}
	if err := dec.Decode(&deletes); err != nil {
		t.Fatal(err)
	}
	if err := dec.Decode(&effects); err != nil {
		t.Fatal(err)
	}
	if len(common) != 0 || len(payload) != 1 || len(updates) != 0 || len(deletes) != 0 || len(effects) != 0 {
		t.Fatalf("unexpected envelope")
	}
}

func TestEncodeResultCompressedPayload(t *testing.T) {
	payload := []any{bytes.Repeat([]byte("a"), 4096)}
	body, err := EncodeResult(payload, nil, nil, nil, ResponseOptions{CompressPayload: true, CompressThreshold: 1, LZ4BlockSize: 1024})
	if err != nil {
		t.Fatal(err)
	}
	if len(body) == 0 || body[0] != 0x90 {
		t.Fatalf("first envelope object is not empty common array")
	}
	decodedPayloadRaw, err := DecodeMaybeLZ4BlockArray(body[1:])
	if err != nil {
		t.Fatal(err)
	}
	var decoded []any
	if err := msgpack.Unmarshal(decodedPayloadRaw, &decoded); err != nil {
		t.Fatal(err)
	}
	if len(decoded) != 1 {
		t.Fatalf("payload len = %d", len(decoded))
	}
}
