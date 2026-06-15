package main

import (
	_ "embed"
	"log"
	"net/http"
	"strings"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type intro struct {
	app.Compo
}

// The Render method is where the component appearance is defined. Here, a
// markdown file is displayed as content.
//
//go:embed documents/intro.md
var entry1Content string

func (h *intro) Render() app.UI {
	return newPage().
		Title("Introduction").
		Icon(homeSVG).
		Index(
			newIndexLink().Title("My MSLIS").Href("/"),
			app.Div().Class("separator"),
		).
		Content(
			newMarkdownDoc().MD(entry1Content),
		)
}

type root struct {
	app.Compo
	current string
}

func (r *root) OnMount(ctx app.Context) {
	r.syncHash()

	// React to fragment changes: in-app clicks, back/forward, deep links.
	app.Window().Call("addEventListener", "hashchange",
		app.FuncOf(func(this app.Value, args []app.Value) any {
			ctx.Dispatch(func(ctx app.Context) {
				r.syncHash()
				app.Log("nav ->", r.current) // diagnostic; delete later
			})
			return nil
		}),
	)

	// Stop go-app from grabbing our "#" links for path routing.
	// Capture phase + stopPropagation runs before go-app's own click
	// handler; we just set the hash, which fires the listener above.
	app.Window().Get("document").Call("addEventListener", "click",
		app.FuncOf(func(this app.Value, args []app.Value) any {
			e := args[0]
			a := e.Get("target").Call("closest", "a")
			if a.IsNull() {
				return nil
			}
			href := a.Call("getAttribute", "href").String()
			if strings.HasPrefix(href, "#") {
				e.Call("preventDefault")
				e.Call("stopPropagation")
				app.Window().Get("location").Set("hash", strings.TrimPrefix(href, "#"))
			}
			return nil
		}),
		true, // capture
	)
}

func (r *root) syncHash() {
	r.current = strings.TrimPrefix(
		app.Window().Get("location").Get("hash").String(), "#")
}

func (r *root) Render() app.UI {
	switch r.current {
	case "philosophy":
		return &philosophy{}
	case "plo1":
		return &plo1{}
	case "plo2":
		return &plo2{}
	case "plo3":
		return &plo3{}
	case "plo4":
		return &plo4{}
	case "plo5":
		return &plo5{}
	case "conclusion":
		return &conclusion{}
	default:
		return &intro{}
	}
}

// The main function is the entry point where the app is configured and started.
// It is executed in 2 different environments: A client (the web browser) and a
// server.
func main() {
	// The first thing to do is to associate the components with a path.
	//
	// This is done by calling the Route() function,  which tells go-app what
	// component to display for a given path, on both client and server-side.
	app.Route("/", func() app.Composer { return &root{} })
	// Once the routes set up, the next thing to do is to either launch the app
	// or the server that serves the app.
	//
	// When executed on the client-side, the RunWhenOnBrowser() function
	// launches the app,  starting a loop that listens for app events and
	// executes client instructions. Since it is a blocking call, the code below
	// it will never be executed.
	//
	// When executed on the server-side, RunWhenOnBrowser() does nothing, which
	// lets room for server implementation without the need for precompiling
	// instructions.
	app.RunWhenOnBrowser()

	// Add this check - if we're in browser, block forever
	if app.IsClient {
		select {} // Block forever - prevent Go runtime from exiting
	}

	// Finally, launching the server that serves the app is done by using the Go
	// standard HTTP package.
	//
	// The Handler is an HTTP handler that serves the client and all its
	// required resources to make it work into a web browser. Here it is
	// configured to handle requests with a path that starts with "/".
	http.Handle("/", &app.Handler{
		Name:        "Home",
		Description: "Home Page",
		Resources:   app.LocalDir("."),
		Styles: []string{
			// "https://fonts.googleapis.com/css2?family=Montserrat:wght@400;500&display=swap",
			"/app.css",
			"/web/css/prism.css",
			"/web/css/docs.css",
		},
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
