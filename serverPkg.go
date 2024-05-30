package server

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	root = flag.String("root", "web/", "Root of the server.")
	port = flag.Int("port", 3380, "Port to serve app over.")
)

type wasmHandler struct {
	fs http.Handler
}

func (h *wasmHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.Path, ".wasm") {
		w.Header().Set("content-type", "application/wasm")
	}
	h.fs.ServeHTTP(w, r)
}

func startHttpServer() {
	// Recommend to run with goroutine so it doesn't block the execution of other JavaScript code; go startHttpServer()
	// flag.Parse()
	if *root == "" || !filepath.IsAbs(*root) {
		log.Fatalln("Root must be an absolute path with non-empty value")
	}
	if *port < 0 || *port > 65535 {
		log.Fatalf("Port must be in the range 0 - 65535, got: %d", *port)
	}
	fs := http.FileServer(http.Dir(*root))
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), &wasmHandler{fs})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Serving files from "+*root, "on http://localhost:"+strconv.Itoa(*port))
}
