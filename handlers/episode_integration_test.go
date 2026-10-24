package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestParseEpisodeMasterIDUsesPathParam(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = httptest.NewRequest(http.MethodGet, "/api/Episodes/140271/GetDetails", nil)
	c.Params = gin.Params{{Key: "episodeMasterId", Value: "140271"}}

	id, ok := parseEpisodeMasterID(c)
	if !ok || id != 140271 {
		t.Fatalf("episode id = %d, %v", id, ok)
	}
}

func TestParseEpisodeMasterIDQueryOverridesPath(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = httptest.NewRequest(http.MethodGet, "/api/Episodes/1/GetDetails?episodeMasterId=140271", nil)
	c.Params = gin.Params{{Key: "episodeMasterId", Value: "1"}}

	id, ok := parseEpisodeMasterID(c)
	if !ok || id != 140271 {
		t.Fatalf("episode id = %d, %v", id, ok)
	}
}
