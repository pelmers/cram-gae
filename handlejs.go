package cram

import (
	"encoding/json"
	"fmt"
	"github.com/pelmers/cram/tokenize/js"
	"math"
	"net/http"
)

// Endpoint for post/js
func handlejs(w http.ResponseWriter, r *http.Request) {
	// only do POST requests
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	form := PostForm{}
	json.NewDecoder(r.Body).Decode(&form)
	reshapedCode := reshapejs(form.Contents, form.Shape, form.Ratio)
	fmt.Fprint(w, reshapedCode)
}

func reshapejs(code, shape string, ratio float64) string {
	tokenizer := js.NewJSTokenizer()
	reshaper := pickReshaper(shape)
	tokens := tokenizer.Tokenize(code, []string{})
	return reshaper(tokens, math.Abs(ratio))
}
