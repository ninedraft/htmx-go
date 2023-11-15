# htmxgo

pckage htmxgo provides utilities for serving htmx library.

Look https://htmx.org/ for more information about htmx.

It can be imported as a library into golang project and hooked up to any HTTP router or server. It supports compression and caches compressed content in memory, but lazily and doesn't waste memory on unused encodings.

All functions are safe for concurrent use.

Currently supported encodings are: 
- plain text 
- deflate 
- gzip 
- brotli

## License

- Module: [BSD 3-Clause License](LICENSE)
- HTMX lib: [BSD 2-Clause License](LICENSE-htmx)

## Usage

```go
import "github.com/ninedraft/htmx-go"
```

```sh
go get -v github.com/ninedraft/htmx-go@latest
```

## Example

```go
// your router setup

mux.Get("/assets/htmx.min.js", htmxgo.ServeHTTP)
```

## Versioning
Release versions are the same as htmx versions.
API can and will be broken.

## Development

Updating htmx version:

```sh
wget 'https://unpkg.com/htmx.org@1.9.8/dist/htmx.min.js' -O htmx.min.js
```

## But why?
For pet experiments. I don't want to import htmx via unpkg nor do I want to setup tooling for frontend stuff.
Don't use in production, it will backfire, I swear to you.
It's mostly a meme module.
