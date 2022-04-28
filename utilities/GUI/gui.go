package GUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"

	network "github.com/acheong08/SimpleResv-Client/utilities/network"

	"fmt"
)

var a fyne.App
var mainWindow fyne.Window

func Run() {
	//Define app and window
	a = app.NewWithID("tech.duti.SimpleResv")
	mainWindow = a.NewWindow("SimpleResv Login")
	//Set main window as master
	mainWindow.SetMaster()
	//Set window content
	mainWindow.SetContent(auth())
	//Set size
	mainWindow.Resize(fyne.NewSize(640, 200))
	//Show window
	mainWindow.ShowAndRun()
}

func auth() fyne.CanvasObject {
	// Enter email
	email := widget.NewEntry()
	email.SetPlaceHolder("example@example.com")
	email.Validator = validation.NewRegexp(`\w{1,}@\w{1,}\.\w{1,4}`, "not a valid email")
	// Enter password
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")
	// Create form
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Email", Widget: email, HintText: "Email address must be valid"},
		},
		OnCancel: func() {
			//Debug output
			fmt.Println("Cancelled")
			//Close app
			mainWindow.Close()
		},
		OnSubmit: func() {
			//Debug output
			fmt.Println("Form submitted")
			// Send requests
			if network.AuthUser(email.Text, password.Text) {
				if network.CheckAdmin(email.Text, password.Text) {
					// Send notification if success
					fyne.CurrentApp().SendNotification(&fyne.Notification{
						Title: "Authenticated as Admin!",
					})
				} else {
					// Send notification if success
					fyne.CurrentApp().SendNotification(&fyne.Notification{
						Title: "Authenticated as User!",
					})
				}
				// Hide done authentication
				mainWindow.Hide()
				// Go to main window
				master()
			} else {
				// Send notification and do nothing
				fyne.CurrentApp().SendNotification(&fyne.Notification{
					Title: "Authentication failed",
				})
			}
		},
	}
	// Append password entry to form
	form.Append("Password", password)
	//Return form
	return form
}

func master() {
	// Reconfigure window
	mainWindow.SetTitle("SimpleResv")
	mainWindow.Resize(fyne.NewSize(640, 460))
	// This fixes window size not updating
	mainWindow.SetFixedSize(true)
	mainWindow.Show()
	mainWindow.Hide()
	mainWindow.SetFixedSize(false)
	// Reconfigure content
	mainWindow.SetContent(widget.NewLabel("Hello world"))
	// Show window again
	mainWindow.Show()
}
