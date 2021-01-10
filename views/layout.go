// Package views has functions to register HTTP handlers that create views in the form of HTML pages.
package views

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

// PageProps are properties for every Page.
type PageProps struct {
	Title string
	Body  g.Node
}

// Page returns a Node that renders an HTML document with the given props.
func Page(props PageProps) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:    props.Title,
		Language: "en",
		Head: []g.Node{
			Link(Rel("stylesheet"), Href("/static/styles/app.css"), Type("text/css")),
		},
		Body: []g.Node{
			Container(
				props.Body,
			),
		},
	})
}

// Container restricts the width of the children, centers, and adds some padding.
func Container(children ...g.Node) g.Node {
	return Div(Class("max-w-7xl mx-auto px-4 sm:px-6 lg:px-8"), g.Group(children))
}
