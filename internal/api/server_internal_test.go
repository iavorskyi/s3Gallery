package api

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/iavorskyi/s3gallery/Services/auth"
	"github.com/iavorskyi/s3gallery/internal/model"
	"github.com/iavorskyi/s3gallery/internal/store/testS3store"
	"github.com/iavorskyi/s3gallery/internal/store/testStore"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_signUp(t *testing.T) {
	srv := newServer(testStore.New(), testS3store.New())

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "success",
			payload: map[string]string{
				"email":    "user@example.com",
				"password": "123456",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "success",
			payload: map[string]string{
				"email":    "user@example.com",
				"password": "1",
			},
			expectedCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			rec := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPost, "/sign-up", nil)
			if err != nil {
				t.Fatal(err)
			}

			jsonbytes, err := json.Marshal(tc.payload)
			if err != nil {
				t.Fatal(err)
			}
			ctx, _ := gin.CreateTestContext(rec)
			ctx.Request = req
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

			//srv.ServeHTTP(rec, req)

			srv.signUp(ctx)
			assert.Equal(t, tc.expectedCode, rec.Code)

		})
	}
}

func TestServer_listItems(t *testing.T) {
	srv := newServer(testStore.New(), testS3store.New())

	// add valid user
	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/sign-up", nil)
	if err != nil {
		t.Fatal(err)
	}

	jsonbytes, err := json.Marshal(model.User{Email: "test@gmail.com", Password: "123456"})
	if err != nil {
		t.Fatal(err)
	}
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
	srv.ServeHTTP(rec, req)

	// success test listItems
	rec2 := httptest.NewRecorder()
	req2, err := http.NewRequest(http.MethodGet, "/api/albums/test/items/", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctx2, _ := gin.CreateTestContext(rec2)
	token, err := auth.GenerateToken(model.User{Email: "test@gmail.com", Password: "123456"}, srv.store)
	if err != nil {
		t.Fatal(err)
	}
	req2.Header.Set("Authorization", "Bearer "+token)
	ctx2.Request = req2
	srv.ServeHTTP(rec2, req2)
	assert.Equal(t, http.StatusOK, rec2.Code)
}
