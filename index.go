package cram

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	render(w, "index.html")
}

func init() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/post/js", handlejs)
}
