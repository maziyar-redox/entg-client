package pages

import (
	"github.com/rivo/tview"
)

func Pages(app *tview.Application, pagePtr *tview.Pages) (*tview.Pages, error) {
	loginPage, loginErr := LoginPage(app, pagePtr)
	if loginErr != nil {
		panic(loginErr)
	}
	chatPage, chatErr := ChatPage(app, pagePtr)
	if chatErr != nil {
		panic(chatErr)
	}
	pagePtr.
		AddPage("login_page", loginPage, true, true).
		AddPage("chat_page", chatPage, true, false)
	return pagePtr, nil
}