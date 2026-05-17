package main

import (
	_ "embed"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// intro is a component that displays a simple "intro World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type plo3 struct {
	app.Compo
}

// The Render method is where the component appearance is defined. Here, a
// markdown file is displayed as content.
//
//go:embed documents/plo3.md
var entry5Content string

func (h *plo3) Render() app.UI {
	return newPage().
		Title("Design and innovate to create equitable, just, and engaging information artifacts, including services, systems, spaces, resources, and technologies").
		Icon(schoolSVG).
		Index(
			newIndexLink().Title("Design and innovate to create equitable, just, and engaging information artifacts, including services, systems, spaces, resources, and technologies").Href("/plo2"),
			app.Div().Class("separator"),
		).
		Content(
			newMarkdownDoc().MD(entry5Content),
		)
}
