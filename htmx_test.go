package htmxgo_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	htmxgo "github.com/ninedraft/htmx-go"
	"github.com/stretchr/testify/assert"
)

func TestServeHTTP(t *testing.T) {
	t.Log(
		"ServeHTTP must write htmx library to response",
	)

	encodings := []string{
		"", "deflate", "gzip", "br",
	}

	for _, encoding := range encodings {
		encoding := encoding

		t.Run(encoding, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest("GET", "/", nil)
			req.Header.Set("Accept-Encoding", encoding)

			got := httptest.NewRecorder()
			got.Body = &bytes.Buffer{}

			htmxgo.ServeHTTP(got, req)

			assert.Equal(t, http.StatusOK, got.Code, "code mismatch")
			assert.Equal(t, "application/javascript", got.Header().Get("Content-Type"), "content-type mismatch")
			assert.Contains(t, got.Header().Get("Content-Encoding"), encoding, "content-encoding mismatch")
			assert.NotEmpty(t, got.Body.Bytes(), "body")
		})
	}
}
