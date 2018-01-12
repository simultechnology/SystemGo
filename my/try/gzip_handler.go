package main

import (
	"github.com/NYTimes/gziphandler"
	"io"
	"net/http"
)

func main() {
	withoutGz := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		io.WriteString(w, "Hello, World")
	})

	withGz := gziphandler.GzipHandler(withoutGz)

	http.Handle("/", withGz)
	http.ListenAndServe("0.0.0.0:48000", nil)
}
