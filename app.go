package main

/*
	Importing dependencies Tview for working with tui and tcell for coloring
*/

import (
	"example.com/m/pages"
	"github.com/rivo/tview"
)

var (
	app *tview.Application
	page *tview.Pages
)

func main() {
	/*
		Bootstraping new application
		Our application is using flex box for design system
		Have two box called Login and Servers for fetching servers
		Then we have some properties
		Our app will fetch servers from an internal server(Iran)
	*/
	// ===================================================
	// Initiating new tview app
	// Going into first page
	// ===================================================
	app = tview.NewApplication()
	page = tview.NewPages()
	pages, pagesErr := pages.Pages(app, page)
	if pagesErr != nil {
		panic(pagesErr)
	}
	err := app.SetRoot(pages, true).SetFocus(pages).EnableMouse(true).Run()
	if err != nil {
		panic(err)
	}
}