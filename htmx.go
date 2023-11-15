// package htmxgo provides utilities for serving htmx library.
package htmxgo

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	_ "embed"

	"github.com/andybalholm/brotli"
)

// on each application start we force clients to reload the library
var lastModified = time.Now().UTC()

// ServeHTTP serves the htmx library. Compression is supported.
// Sets Content-Type to application/javascript.
// Tries to use the best compression available by Accept-Encoding header.
// Concurrent safe.
func ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	encoding := req.Header.Get("Accept-Encoding")
	header := rw.Header()
	data := lib

	// priority is important! br > gzip > deflate
	switch {
	case strings.Contains(encoding, "br"):
		header.Set("Content-Encoding", "br")
		data = libBrotli()
	case strings.Contains(encoding, "gzip"):
		header.Set("Content-Encoding", "gzip")
		data = libGzip()
	case strings.Contains(encoding, "deflate"):
		header.Set("Content-Encoding", "deflate")
		data = libDeflate()
	}

	header.Set("Content-Type", "application/javascript")
	header.Set("Content-Length", strconv.Itoa(len(data)))

	http.ServeContent(rw, req, "", lastModified, bytes.NewReader(data))
}

//go:embed htmx.min.js
var lib []byte

var libDeflate = compressOnceMust(flate.NewWriter(nil, flate.BestCompression))

var libGzip = compressOnceMust(gzip.NewWriterLevel(nil, gzip.BestCompression))

var libBrotli = compressOnceMust(brotli.NewWriterLevel(nil, brotli.BestCompression), nil)

type compressor interface {
	io.WriteCloser
	Reset(io.Writer)
}

func compressOnceMust(c compressor, _ error) func() []byte {
	return sync.OnceValue(func() []byte {
		buf := &bytes.Buffer{}
		c.Reset(buf)

		if _, err := c.Write(lib); err != nil {
			panic("compressing: " + err.Error())
		}

		if err := c.Close(); err != nil {
			panic("compressing: " + err.Error())
		}

		return buf.Bytes()
	})
}
