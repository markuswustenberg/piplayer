package views

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"

	"piplayer/model"
)

func HomeBody(albums []model.Album) g.Node {
	return Div(Class("py-4"),
		A(Href("/pause"), g.Text("Pause"), Class("mb-16 inline-flex items-center px-2.5 py-1.5 border border-transparent text-xs font-medium rounded shadow-sm text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500")),

		Ul(Class("space-y-8"),
			g.Group(g.Map(len(albums), func(i int) g.Node {
				return Li(A(Href("/play/"+albums[i].ID), g.Textf("%v - %v", albums[i].Artist, albums[i].Name)))
			})),
		),
	)
}
