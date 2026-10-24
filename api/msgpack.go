package api

import (
	"github.com/gin-gonic/gin"

	"ymst_server/apiwire"
)

const MsgpackContentType = apiwire.MsgpackContentType

func WriteApiMsgpackResult(c *gin.Context, payload any) error {
	return apiwire.WritePayload(c, payload)
}

func WriteApiMsgpackResultWithUpdates(c *gin.Context, payload any, updates any) error {
	return apiwire.WritePayloadWithUpdates(c, payload, updates)
}

func WriteBooleanResult(c *gin.Context, success bool, updates any) error {
	return apiwire.WriteBoolean(c, success, updates)
}

func EncodeApiMsgpackResult(payload any) ([]byte, error) {
	return EncodeApiMsgpackResultFull(payload, nil, nil, nil)
}

func EncodeApiMsgpackResultFull(payload any, updates any, deletes any, effects any) ([]byte, error) {
	return apiwire.EncodeResult(payload, updates, deletes, effects, apiwire.DefaultResponseOptions())
}

func WriteApiMsgpackResultFull(c *gin.Context, payload any, updates any, deletes any, effects any) error {
	return apiwire.WriteResult(c, payload, updates, deletes, effects, apiwire.DefaultResponseOptions())
}

func WriteMsgpackResult(c *gin.Context, payload any) error { return WriteApiMsgpackResult(c, payload) }
