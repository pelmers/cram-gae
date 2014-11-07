package cram

import (
	"encoding/json"
	"fmt"
	"github.com/pelmers/cram/tokenize/js"
	"log"
	"math"
	"net/http"
)

// The body of the POST request the client generates.
type PostForm struct {
	// the code to be crammed
	Contents string
	// the shape to transform the code into (e.g. "rectangle")
	Shape string
	// the height:width ratio to use
	Ratio float64
}

func handlejs(w http.ResponseWriter, r *http.Request) {
	// only do POST requests
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	form := PostForm{}
	json.NewDecoder(r.Body).Decode(&form)
	reshapedCode := reshapeJS(form.Contents, form.Shape, form.Ratio)
	fmt.Fprint(w, reshapedCode)
}

func reshapeJS(code, shape string, ratio float64) string {
	tokenizer := js.NewJSTokenizer()
	reshaper := pickReshaper(shape)
	tokens := tokenizer.Tokenize(code, []string{})
	log.Print(ratio)
	return reshaper(tokens, math.Abs(ratio))
}
