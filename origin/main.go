package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
)

func main() {

	flag.Parse()

	bindAddr := flag.Arg(0)
	if bindAddr == "" {
		bindAddr = ":80"
	}

	log.Printf("listening on %s", bindAddr)
	err := http.ListenAndServe(bindAddr, EchoHandler{})
	log.Fatalf("serve: %v", err)
}

// EchoHandler handles requests by echoing them back in the response body.
type EchoHandler struct{}

// ServeHTTP echos requests back in the response body.
func (EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Log the request.
	log.Printf("Got request: %s %s %s", r.Method, r.URL.Path, r.Proto)

	// Set an example response header.
	w.Header().Set("x-example-header", "Hello World!")

	// Write the first request line back to the response body.
	resp := fmt.Sprintf("%s %s %s\n", r.Method, r.URL.Path, r.Proto)
	_, err := w.Write([]byte(resp))
	if err != nil {
		log.Printf("write first request line: %v", err)
		return
	}

	// Write the request headers back to the response body.
	keys := make([]string, 0, len(r.Header))
	for k := range r.Header {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		resp := fmt.Sprintf("%s: %s\n", key, r.Header[key])
		_, err := w.Write([]byte(resp))
		if err != nil {
			log.Printf("write header \"%s: %s\": %v", key, r.Header[key], err)
			return
		}
	}
	_, err = w.Write([]byte("\n"))
	if err != nil {
		log.Printf("write header terminator: %v", err)
		return
	}

	// Write the request body back to the response, limited to 1 megabyte.
	body := http.MaxBytesReader(w, r.Body, 1<<20)
	var buf [2048]byte
	for {
		_, err = body.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			w.Write([]byte("An error occurred"))
			log.Printf("echo body: %v", err)
			return
		}
	}
}
