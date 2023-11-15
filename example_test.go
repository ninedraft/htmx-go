package htmxgo_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	htmxgo "github.com/ninedraft/htmx-go"
)

func ExampleServeHTTP() {
	req := httptest.NewRequest("GET", "/htmx.min.js", nil)
	req.Header.Set("Accept-Encoding", "br")

	mux := http.NewServeMux()
	mux.HandleFunc("/htmx.min.js", htmxgo.ServeHTTP)

	rw := httptest.NewRecorder()

	mux.ServeHTTP(rw, req)

	fmt.Println(rw.Header().Get("Content-Encoding"))
	fmt.Println(rw.Header().Get("Content-Type"))

	// Output:
	// br
	// application/javascript
}
