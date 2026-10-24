package apiwire

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vmihailenco/msgpack/v5"
)

const MsgpackContentType = "application/vnd.msgpack"

type ResponseOptions struct {
	CompressPayload   bool
	CompressThreshold int
	LZ4BlockSize      int
}

func DefaultResponseOptions() ResponseOptions {
	return ResponseOptions{CompressPayload: false, CompressThreshold: 64 * 1024, LZ4BlockSize: 64 * 1024}
}

func EncodeResult(payload any, updates any, deletes any, effects any, opt ResponseOptions) ([]byte, error) {
	if updates == nil {
		updates = []any{}
	}
	if deletes == nil {
		deletes = []any{}
	}
	if effects == nil {
		effects = []any{}
	}

	var packedPayload []byte
	encodedPayload := payload
	if opt.CompressPayload {
		rawPayload, err := msgpack.Marshal(payload)
		if err != nil {
			return nil, err
		}
		if opt.CompressThreshold <= 0 || len(rawPayload) >= opt.CompressThreshold {
			packed, err := EncodeLZ4BlockArray(rawPayload, opt.LZ4BlockSize)
			if err != nil {
				return nil, err
			}
			packedPayload = packed
		}
	}

	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)
	if err := enc.Encode([]any{}); err != nil {
		return nil, err
	}
	if packedPayload != nil {
		if _, err := buf.Write(packedPayload); err != nil {
			return nil, err
		}
	} else if err := enc.Encode(encodedPayload); err != nil {
		return nil, err
	}
	for _, obj := range []any{updates, deletes, effects} {
		if err := enc.Encode(obj); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func WriteResult(c *gin.Context, payload any, updates any, deletes any, effects any, opt ResponseOptions) error {
	body, err := EncodeResult(payload, updates, deletes, effects, opt)
	if err != nil {
		return err
	}
	c.Header("Content-Type", MsgpackContentType)
	c.Header("X-FM", "0")
	c.Data(http.StatusOK, MsgpackContentType, body)
	return nil
}

func WritePayload(c *gin.Context, payload any) error {
	return WriteResult(c, payload, nil, nil, nil, DefaultResponseOptions())
}

func WritePayloadWithUpdates(c *gin.Context, payload any, updates any) error {
	return WriteResult(c, payload, updates, nil, nil, DefaultResponseOptions())
}

func WriteBoolean(c *gin.Context, success bool, updates any) error {
	return WritePayloadWithUpdates(c, BooleanResult(success), updates)
}
