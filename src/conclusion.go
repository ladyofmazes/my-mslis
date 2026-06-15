package main

import (
	_ "embed"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// intro is a component that displays a simple "intro World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type conclusion struct {
	app.Compo
}

// The Render method is where the component appearance is defined. Here, a
// markdown file is displayed as content.
//
//go:embed documents/conclusion.md
var entry8Content string

func (h *conclusion) Render() app.UI {
	return newPage().
		Title("Conclusion").
		Icon(schoolSVG).
		Index(
			newIndexLink().Title("Conclusion").Href("#conclusion"),
			app.Div().Class("separator"),
		).
		Content(
			newMarkdownDoc().MD(entry8Content),
		)
}
