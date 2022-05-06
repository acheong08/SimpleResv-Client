package GUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	data "github.com/acheong08/SimpleResv-Client/data"
	network "github.com/acheong08/SimpleResv-Client/utilities/network"

	"strconv"
)

var a fyne.App
var mainWindow fyne.Window

var savedEmail string
var savedPassword string

var isAdmin bool

func Run() {
	//Define app and window
	a = app.NewWithID("tech.duti.SimpleResv")
	a.SetIcon(theme.AccountIcon())
	mainWindow = a.NewWindow("SimpleResv")
	//Set main window as master
	mainWindow.SetMaster()
	// Start authentication
	auth(0)
}

// Window handlers

func auth(num int) {
	//Set window title
	mainWindow.SetTitle("SimpleResv | Login")
	//Set window content
	mainWindow.SetContent(authCanvas())
	//Set size
	fixedResize(640, 200)
	//Show window. Run only if first
	if num == 0 {
		mainWindow.ShowAndRun()
	} else {
		mainWindow.Show()
	}
}

func itemIndex() {
	// Reconfigure window
	mainWindow.SetTitle("SimpleResv | Available")
	// Resize window
	fixedResize(640, 460)
	// Reconfigure content
	mainWindow.SetContent(itemIndexCanvas()) // To be replaced by welcome screen
	// Show window again
	mainWindow.Show()
}

func reservedIndex() {
	// Reconfigure window
	mainWindow.SetTitle("SimpleResv | Reserved by you")
	// Resize window
	fixedResize(640, 460)
	// Reconfigure content
	mainWindow.SetContent(reservedIndexCanvas()) // To be replaced by welcome screen
	// Show window again
	mainWindow.Show()
}

func userList() {
	// Make new window
	w := a.NewWindow("SimpleResv | User List")
	// Resize window
	w.Resize(fyne.NewSize(100, 400))
	// Configure content
	w.SetContent(userListCanvas())
	// Show window
	w.Show()
}

// Make menu
func makeMenu() *fyne.MainMenu {
	// Menu items
	// Account stuff
	logout := fyne.NewMenuItem("Logout", func() {
		// Hide current window
		mainWindow.Hide()
		// Clear saved credentials
		savedEmail = ""
		savedPassword = ""
		// Remove admin permission
		isAdmin = false
		// Refresh menu
		mainWindow.SetMainMenu(makeMenu())
		// Go back to authentication window
		auth(1)
	})
	// Window stuff
	available := fyne.NewMenuItem("Available", func() {
		mainWindow.Hide()
		itemIndex()
	})
	reserved := fyne.NewMenuItem("Reserved", func() {
		mainWindow.Hide()
		reservedIndex()
	})
	// Admin stuff
	if isAdmin {
		// Item management
		addItem := fyne.NewMenuItem("Add Item", func() {
			// Form items
			itemName := widget.NewEntry()
			itemDesc := widget.NewMultiLineEntry()
			// Form item list
			formItems := []*widget.FormItem{
				widget.NewFormItem("Name", itemName),
				widget.NewFormItem("Description", itemDesc),
			}
			// Create form
			form := dialog.NewForm("Add Item", "Add", "Cancel", formItems, func(b bool) {
				if !b {
					return
				}
				addItem(itemName.Text, itemDesc.Text)
			}, mainWindow)
			// Set size and show
			form.Resize(fyne.NewSize(400, 250))
			form.Show()
		})
		// Deletion
		delItem := fyne.NewMenuItem("Delete Item", func() {
			// Form items
			itemID := widget.NewEntry()
			// Form item list
			formItems := []*widget.FormItem{
				widget.NewFormItem("ID: ", itemID),
			}
			// Show form
			dialog.ShowForm("Delete Item", "Delete", "Cancel", formItems, func(b bool) {
				if !b {
					return
				}
				intItemID, _ := strconv.Atoi(itemID.Text)
				delItem(intItemID)
			}, mainWindow)
		})
		// User management
		userAdd := fyne.NewMenuItem("Add User", func() {
			// Form items
			email := widget.NewEntry()
			email.SetPlaceHolder("example@example.com")
			email.Validator = validation.NewRegexp(`\w{1,}@\w{1,}\.\w{1,4}`, "not a valid email")
			password := widget.NewPasswordEntry()
			password.SetPlaceHolder("Password")
			// Create list for form items
			formItems := []*widget.FormItem{
				widget.NewFormItem("Email", email),
				widget.NewFormItem("Password", password),
			}
			//Create and show dialog form
			form := dialog.NewForm("Add User", "Add", "Cancel", formItems, func(b bool) {
				if !b {
					return
				}
				addUser(email.Text, password.Text)
			}, mainWindow)
			// Set size and show
			form.Resize(fyne.NewSize(400, 250))
			form.Show()
		})
		userDel := fyne.NewMenuItem("Delete User", func() {
			// Form items
			email := widget.NewEntry()
			email.SetPlaceHolder("example@example.com")
			email.Validator = validation.NewRegexp(`\w{1,}@\w{1,}\.\w{1,4}`, "not a valid email")
			// Create list for form items
			formItems := []*widget.FormItem{
				widget.NewFormItem("Email", email),
			}
			//Create and show dialog form
			form := dialog.NewForm("Delete User", "Delete", "Cancel", formItems, func(b bool) {
				if !b {
					return
				}
				delUser(email.Text)
			}, mainWindow)
			// Set size and show
			form.Resize(fyne.NewSize(400, 250))
			form.Show()
		})
		listUser := fyne.NewMenuItem("User list", func() {
			userList()
		})
		// Settings
		settingsItem := fyne.NewMenuItem("Settings", func() {
			w := a.NewWindow("Fyne Settings")
			w.SetContent(settings.NewSettings().LoadAppearanceScreen(mainWindow))
			w.Resize(fyne.NewSize(480, 480))
			w.Show()
		})
		// Make menu with logout as item
		account := fyne.NewMenu("Account", logout)
		// Make menu for window stuff
		windows := fyne.NewMenu("Windows", available, reserved)
		// Settings only in desktop
		if !fyne.CurrentDevice().IsMobile() {
			account.Items = append(account.Items, fyne.NewMenuItemSeparator(), settingsItem)
		}
		admin := fyne.NewMenu("Admin", addItem, delItem, userAdd, userDel, listUser)
		return fyne.NewMainMenu(
			account,
			windows,
			admin,
		)

	} else {
		// Settings
		settingsItem := fyne.NewMenuItem("Settings", func() {
			w := a.NewWindow("Fyne Settings")
			w.SetContent(settings.NewSettings().LoadAppearanceScreen(mainWindow))
			w.Resize(fyne.NewSize(480, 480))
			w.Show()
		})
		// Make menu with logout as item
		account := fyne.NewMenu("Account", logout)
		// Make menu for window stuff
		windows := fyne.NewMenu("Windows", available, reserved)
		// Settings only in desktop
		if !fyne.CurrentDevice().IsMobile() {
			account.Items = append(account.Items, fyne.NewMenuItemSeparator(), settingsItem)
		}
		if savedEmail == "" {
			return fyne.NewMainMenu(
				account,
			)
		} else {
			return fyne.NewMainMenu(
				account,
				windows,
			)
		}
	}
}

// Canvases
func authCanvas() fyne.CanvasObject {
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
			//Close app
			mainWindow.Close()
		},
		OnSubmit: func() {
			// Send requests
			if network.AuthUser(email.Text, password.Text) {
				if network.CheckAdmin(email.Text, password.Text) {
					// Send notification if success
					fyne.CurrentApp().SendNotification(&fyne.Notification{
						Title: "Authenticated as Admin!",
					})
					// Save status
					isAdmin = true
				} else {
					// Send notification if success
					fyne.CurrentApp().SendNotification(&fyne.Notification{
						Title: "Authenticated as User!",
					})
					// Save status
					isAdmin = false
				}
				// Save authentication to global variable
				savedEmail = email.Text
				savedPassword = password.Text
				// Hide done authentication
				mainWindow.Hide()
				// Make a top menu
				mainWindow.SetMainMenu(makeMenu())
				// Go to main window
				itemIndex()
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

func itemIndexCanvas() fyne.CanvasObject {
	// Get items from server and set to list of item struct
	var items []data.Item = network.GetItems()
	// Create grid
	grid := container.New(layout.NewAdaptiveGridLayout(4))
	// Loop through items
	for _, itemVals := range items {
		// Only show if not reserved
		if itemVals.Available {
			// Set a new id variable to pass by value
			id := itemVals.Id
			// Define anonymous takeItem function
			anonTake := func() {
				takeItem(id)
			}
			// Create card for each item
			button := widget.NewButton("Take", anonTake)
			strId := "ID: " + strconv.Itoa(itemVals.Id)
			label := widget.NewLabel(string("Status: " + itemVals.Status + "\nDetails: " + itemVals.Details))
			submit := container.New(layout.NewGridLayoutWithRows(2), label, button)
			grid.Add(widget.NewCard(itemVals.Name, strId, submit))
		}
	}
	// Create a button to refresh the data
	refreshBtn := widget.NewButton("Refresh", func() { itemIndex() })
	grid.Add(refreshBtn)
	// Make scroll
	scrollCont := container.NewVScroll(grid)
	// Return grids
	return scrollCont
}

func reservedIndexCanvas() fyne.CanvasObject {
	// Get items from server and set to list of item struct
	var items []data.Item = network.GetItems()
	// Create grid
	grid := container.New(layout.NewAdaptiveGridLayout(4))
	// Loop through items
	for _, itemVals := range items {
		// Only show if not reserved
		if itemVals.Status == savedEmail || isAdmin {
			// Set a new id variable to pass by value
			id := itemVals.Id
			// Define anonymous takeItem function
			anonReturn := func() {
				returnItem(id)
			}
			// Create card for each item
			button := widget.NewButton("Return", anonReturn)
			strId := "ID: " + strconv.Itoa(itemVals.Id)
			label := widget.NewLabel(string("Status: " + itemVals.Status + "\nDetails: " + itemVals.Details))
			submit := container.New(layout.NewGridLayoutWithRows(2), label, button)
			grid.Add(widget.NewCard(itemVals.Name, strId, submit))
		}
	}
	// Create a button to refresh the data
	refreshBtn := widget.NewButton("Refresh", func() { reservedIndex() })
	grid.Add(refreshBtn)
	// Make scroll
	scrollCont := container.NewVScroll(grid)
	// Return grids
	return scrollCont
}

func userListCanvas() fyne.CanvasObject {
	// Make a container
	cont := container.New(layout.NewVBoxLayout())
	// Get list
	accounts := network.GetUserList(savedEmail, savedPassword)
	// Put into string list
	var accountlist string = ""
	for _, accountVals := range accounts{
		accountlist = accountlist + "\n" + accountVals.Email
	}
	// Put list into labels
	list := widget.NewLabel(accountlist)
	// Put label into container
	cont.Add(list)
	// Make container scrollable
	scrollCont := container.NewVScroll(cont)
	// Return
	return scrollCont
}
// Utilities
func fixedResize(width float32, height float32) {
	// Set size
	mainWindow.Resize(fyne.NewSize(width, height))
	// This fixes window size not updating
	mainWindow.SetFixedSize(true)
	mainWindow.Show()
	mainWindow.Hide()
	mainWindow.SetFixedSize(false)
}

//User items
func takeItem(id int) {
	// Send request
	network.TakeItem(savedEmail, savedPassword, id)
	// Refresh content
	mainWindow.SetContent(itemIndexCanvas())
}
func returnItem(id int) {
	// Send request
	network.ReturnItem(savedEmail, savedPassword, id)
	// Refresh content
	mainWindow.SetContent(reservedIndexCanvas())
}

//Admin items
func addItem(name string, details string) {
	notifyMessage(network.AddItem(savedEmail, savedPassword, name, details))
}
func delItem(id int) {
	notifyMessage(network.DeleteItem(savedEmail, savedPassword, id))
}

//Admin Users
func addUser(email string, password string) {
	notifyMessage(network.AddUser(savedEmail, savedPassword, email, password))
}
func delUser(email string) {
	notifyMessage(network.DeleteUser(savedEmail, savedPassword, email))
}

// Notify status
func notifyMessage(text string) {
	fyne.CurrentApp().SendNotification(&fyne.Notification{
		Title:   "SimpleResv Action",
		Content: text,
	})
}
