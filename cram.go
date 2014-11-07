package cram

import (
	"github.com/pelmers/cram/shapes"
	"net/http"
	"strings"
)

func pickReshaper(option string) shapes.Reshaper {
	switch option {
	case "rect":
		return shapes.Square
	case "triangle":
		return shapes.Triangle
	case "trapezoid":
		return shapes.Trapezoid
	case "ellipse":
		return shapes.Ellipse
	case "diamond":
		return shapes.Diamond
	case "pepper":
		return shapes.Pepper
	}
	// default choice is to just concatenate everything
	return func(tok []string, _ float64) string {
		return strings.Join(tok, "  ")
	}
}

func init() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/post", handlejs)
}
