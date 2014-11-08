package cram

import (
	"github.com/pelmers/cram/shapes"
	"strings"
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
	// default choice is to just concatenate everything on one line
	return func(tok []string, _ float64) string {
		return strings.Join(tok, "")
	}
}
