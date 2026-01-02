package pages

import (
	"bytes"
	"net/http"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	_ int = iota
	IDLE
	LOGGING_IN
	ERROR
)

var (
	NICKNAME string = "Nickname"
	USER string = "Username"
	PASSWORD string = "Password"
	LOGIN string = "Login"
	QUIT string = "Quit"
	SERVER string = "Server"
	SERVERS_LIST []string = []string{}
	DEFAULT_SERVER string = "https://maziyar-isa.ir/"
	IS_SERVERS_LOADED bool = false
	REPO_URL string = "https://raw.githubusercontent.com/maziyar-redox/entg-protocol/refs/heads/main/servers.txt"
)

func charValidator(textToCheck string, lastChar rune) bool {
	if len(textToCheck) > 10 {
		return false
	}
	return true
}

// TODO = Fix hanging while sending request

func LoginPage(app *tview.Application, pagePtr *tview.Pages) (*tview.Flex, error) {
	// ===================================
	// Loading text view
	// ===================================
	textView := tview.NewTextView().
        SetText("Loading Servers...").
        SetTextAlign(tview.AlignLeft).
        SetDynamicColors(true)
    textView.SetBorder(true).
		SetBorderColor(tcell.ColorBlue).
        SetTitle(" Servers List ")
	// ===================================
	// Login form view
	// ===================================
	loginForm := tview.NewForm().
		AddInputField(NICKNAME, "", 30, charValidator, nil).
		AddInputField(USER, "", 30, charValidator, nil).
		AddPasswordField(PASSWORD, "", 30, '*', nil).
		AddTextView(SERVER, "Loading...", 30, 3, true, false).
		AddButton(LOGIN, func() {
			if IS_SERVERS_LOADED == false {
				return
			}
			pagePtr.RemovePage("login_page")
			pagePtr.ShowPage("chat_page")
		}).
		AddButton(QUIT, func() {
			app.Stop()
		})
	loginForm.SetBorder(true).
		SetBorderColor(tcell.ColorYellow).
		SetTitle(" Login Box ")
	// ===================================
	// Servers form view
	// ===================================
	serversList := tview.NewList().ShowSecondaryText(false)
	serversList.SetBorder(true).
		SetBorderColor(tcell.ColorBlue).
		SetTitle(" Servers Box ")
	serversList.SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		go func() {
			app.QueueUpdateDraw(func() {
				loginForm.RemoveFormItem(3)
				loginForm.AddTextView(SERVER, mainText, 30, 3, true, false)
			})
		}()
	})
	// ===================================
	// Login page
	// ===================================
	flex := tview.NewFlex().
		AddItem(loginForm, 0, 1, false).
		AddItem(textView, 50, 1, false)
	// ===================================
	// Fetching Servers
	// ===================================
	go func() {
		jsonBody := []byte("")
		bodyReader := bytes.NewReader(jsonBody)
		req, err := http.NewRequest(http.MethodGet, REPO_URL, bodyReader)
		if err != nil {
			app.QueueUpdateDraw(func ()  {
				textView.SetText("An error ocured while creating http request, Exiting in 5 secconds")
			})
			time.Sleep(5 * time.Second)
			app.Stop()
		}
		client := http.Client{
			Timeout: 5 * time.Second,
		}
		res, err := client.Do(req)
		if err != nil || res.StatusCode != 200 {
			app.QueueUpdateDraw(func ()  {
				textView.SetText("An error ocured while sending http request, Using default server...")
				time.Sleep(5 * time.Second)
				flex.RemoveItem(textView)
				flex.AddItem(serversList, 50, 1, false)
				serversList.AddItem(DEFAULT_SERVER, "", 0, nil)
				loginForm.RemoveFormItem(3)
				loginForm.AddTextView(SERVER, DEFAULT_SERVER, 30, 3, true, false)
				IS_SERVERS_LOADED = true
			})
			return
		}
		app.QueueUpdateDraw(func ()  {
			flex.RemoveItem(textView)
			flex.AddItem(serversList, 50, 1, false)
			IS_SERVERS_LOADED = true
		})
	}()
	return flex, nil
}