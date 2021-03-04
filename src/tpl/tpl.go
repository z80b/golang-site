package tpl

import("github.com/flosch/pongo2")

func mainPage() {
	return pongo2.Must(pongo2.FromFile("templates/main.html"))
}
