package main

import (
	_ "embed"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// intro is a component that displays a simple "intro World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type plo2 struct {
	app.Compo
}

// The Render method is where the component appearance is defined. Here, a
// markdown file is displayed as content.
//
//go:embed documents/plo2.md
var entry4Content string

func (h *plo2) Render() app.UI {
	return newPage().
		Title("Understand, engage, and serve users and their communities").
		Icon(schoolSVG).
		Index(
			newIndexLink().Title("Understand, engage, and serve users and their communities").Href("/plo2"),
			app.Div().Class("separator"),
		).
		Content(
			newMarkdownDoc().MD(entry4Content),
		)
}
