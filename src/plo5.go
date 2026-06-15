package main

import (
	_ "embed"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// intro is a component that displays a simple "intro World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type plo5 struct {
	app.Compo
}

// The Render method is where the component appearance is defined. Here, a
// markdown file is displayed as content.
//
//go:embed documents/plo5.md
var entry7Content string

func (h *plo5) Render() app.UI {
	return newPage().
		Title("Demonstrate information literacy and technological agility").
		Icon(schoolSVG).
		Index(
			newIndexLink().Title("Demonstrate information literacy and technological agility").Href("#plo5"),
			app.Div().Class("separator"),
		).
		Content(
			newMarkdownDoc().MD(entry7Content),
		)
}
