package main

import (
	_ "embed"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// intro is a component that displays a simple "intro World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type plo1 struct {
	app.Compo
}

// The Render method is where the component appearance is defined. Here, a
// markdown file is displayed as content.
//
//go:embed documents/plo1.md
var entry3Content string

func (h *plo1) Render() app.UI {
	return newPage().
		Title("Advance Information Equity and Justice").
		Icon(schoolSVG).
		Index(
			newIndexLink().Title("Advance Information Equity and Justice").Href("/plo1"),
			app.Div().Class("separator"),
		).
		Content(
			newMarkdownDoc().MD(entry3Content),
		)
}
