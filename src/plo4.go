package main

import (
	_ "embed"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// intro is a component that displays a simple "intro World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type plo4 struct {
	app.Compo
}

// The Render method is where the component appearance is defined. Here, a
// markdown file is displayed as content.
//
//go:embed documents/plo4.md
var entry6Content string

func (h *plo4) Render() app.UI {
	return newPage().
		Title("Lead and manage people and projects in an equitable, just, and culturally competent manner").
		Icon(schoolSVG).
		Index(
			newIndexLink().Title("Lead and manage people and projects in an equitable, just, and culturally competent manner").Href("#plo4"),
			app.Div().Class("separator"),
		).
		Content(
			newMarkdownDoc().MD(entry6Content),
		)
}
