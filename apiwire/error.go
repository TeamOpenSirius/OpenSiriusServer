package apiwire

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WriteErrorStatus(c *gin.Context, status int) {
	c.Header("Content-Type", MsgpackContentType)
	c.Header("X-FM", "0")
	c.Status(status)
	if status == http.StatusNoContent {
		return
	}
}
