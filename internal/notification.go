package internal

import (
	"github.com/gen2brain/beeep"
)

var (
	beepApp error
)

// ===================================================
// Beep config for showing notification
// ===================================================

func SendNotification(
	Title string,
	MessageBody string,
) error {
	var icon []byte
	beeep.AppName = "ENTG-Application"
	beepApp = beeep.Notify(Title, MessageBody, icon)
	if beepApp != nil {
		panic(beepApp)
	}
	return nil
}