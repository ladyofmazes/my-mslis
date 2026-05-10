package main

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"github.com/maxence-charriere/go-app/v10/pkg/ui"
)

type menu struct {
	app.Compo

	Iclass string

	appInstallable bool
}

func newMenu() *menu {
	return &menu{}
}

func (m *menu) Class(v string) *menu {
	m.Iclass = app.AppendClass(m.Iclass, v)
	return m
}

func (m *menu) OnNav(ctx app.Context) {
	m.appInstallable = ctx.IsAppInstallable()
}

func (m *menu) OnAppInstallChange(ctx app.Context) {
	m.appInstallable = ctx.IsAppInstallable()
}

func (m *menu) Render() app.UI {
	linkClass := "link heading fit unselectable"

	isFocus := func(path string) string {
		if app.Window().URL().Path == path {
			return "focus"
		}
		return ""
	}

	return ui.Scroll().
		Class("menu").
		Class(m.Iclass).
		HeaderHeight(headerHeight).
		Header(
			ui.Stack().
				Class("fill").
				Middle().
				Content(
					app.Header().Body(
						app.A().
							Class("heading").
							Class("app-title").
							Href("https://go-app.dev/").
							Text("Go-App"),
					),
				),
		).
		Content(
			app.Nav().Body(
				app.Div().Class("separator"),

				ui.Link().
					Class(linkClass).
					Icon(homeSVG).
					Label("Home").
					Href("/").
					Class(isFocus("/")),

				app.Div().Class("separator"),

				ui.Link().
					Class(linkClass).
					Icon(imgFolderSVG).
					Label("Philosophy").
					Href("/philosophy").
					Class(isFocus("/philosophy")),

				ui.Link().
					Class(linkClass).
					Icon(imgFolderSVG).
					Label("PLO1").
					Href("/plo1").
					Class(isFocus("/plo1")),

				ui.Link().
					Class(linkClass).
					Icon(imgFolderSVG).
					Label("PLO2").
					Href("/plo2").
					Class(isFocus("/plo2")),
				app.Div().Class("separator"),
				ui.Link().
					Class(linkClass).
					Icon(userLockSVG).
					Label("Resume").
					Href("https://d3aobyces4igtg.cloudfront.net/"),

				app.Div().Class("separator"),
			),
		)
}

func (m *menu) installApp(ctx app.Context, e app.Event) {
	ctx.NewAction(installApp)
}
