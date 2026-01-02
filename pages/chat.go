package pages

import (
	"github.com/rivo/tview"
)

func ChatPage(app *tview.Application, pagePtr *tview.Pages) (*tview.Flex, error) {
	loginForm := tview.NewForm().
		AddInputField(NICKNAME, "", 30, charValidator, nil)
	flex := tview.NewFlex().
		AddItem(loginForm, 0, 1, false)

	return flex, nil
}