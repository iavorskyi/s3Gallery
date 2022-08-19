package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http/httptest"
	"testing"
)

func Test_PingHandler(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(rec)

	s.apiPing(ctx)
	assert.Equal(t, rec.Body.String(), "PONG")
}
