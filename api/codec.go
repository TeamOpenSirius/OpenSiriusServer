package api

import "ymst_server/apiwire"

func DecodeMaybeLz4BlockArray(body []byte) ([]byte, error) {
	return apiwire.DecodeMaybeLZ4BlockArray(body)
}
