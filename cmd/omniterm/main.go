package main

/*
 * libraries
 * https://pkg.go.dev/golang.org/x/term
 * https://pkg.go.dev/tinygo.org/x/bluetooth
 * https://pkg.go.dev/go.bug.st/serial
 * https://pkg.go.dev/github.com/diamondburned/gotkit
 * https://pkg.go.dev/github.com/diamondburned/gotk4/pkg
 * https://pkg.go.dev/github.com/diamondburned/gotk4-adwaita/pkg/adw
 */

import (
	"os"

	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func main() {
	app := gtk.NewApplication("com.pcdworks.omniterm", 0)
	app.Connect("activate", activate)

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

func newMainMenu() *gtk.MenuButton {
	// main menu
	mainMenu := gio.NewMenu()

	// size control
	sizeControl := gio.NewMenu()

	// zoom out
	zoomOut := gio.NewMenuItem("zoom-out", "win.zoom-out")
	zoomOut.SetAttributeValue("verb-icon", glib.NewVariantString("zoom-out-symbolic"))
	sizeControl.InsertItem(0, zoomOut)

	// 100% size
	normal := gio.NewMenuItem("100%", "win.zoom-normal")
	sizeControl.InsertItem(1, normal)

	// zoom in
	zoomIn := gio.NewMenuItem("zoom-in", "win.zoom-in")
	zoomIn.SetAttributeValue("verb-icon", glib.NewVariantString("zoom-in-symbolic"))
	sizeControl.InsertItem(2, zoomIn)

	// size control section
	sizeSection := gio.NewMenuItemSection("", sizeControl)
	sizeSection.SetAttributeValue("display-hint", glib.NewVariantString("horizontal-buttons"))
	mainMenu.InsertItem(0, sizeSection)

	// window control
	windowControl := gio.NewMenu()

	// New window menu entry
	mWindow := gio.NewMenuItem("New Window", "app.new-window")
	windowControl.InsertItem(0, mWindow)

	// Fullscreen menu entry
	mFullScreen := gio.NewMenuItem("Fullscreen", "window.fullscreen")
	windowControl.InsertItem(1, mFullScreen)

	windowSection := gio.NewMenuItemSection("", windowControl)
	mainMenu.InsertItem(1, windowSection)

	// Preferences menu entry
	mPref := gio.NewMenuItem("Preferences", "app.preferences")
	mainMenu.InsertItem(2, mPref)

	// About menu entry
	mAbout := gio.NewMenuItem("About", "app.about")
	mainMenu.InsertItem(3, mAbout)

	// menu button
	mainButton := gtk.NewMenuButton()
	mainButton.SetVAlign(gtk.AlignCenter)
	mainButton.SetIconName("open-menu-symbolic")
	mainButton.SetMenuModel(mainMenu)
	return mainButton
}

func newTabMenu() *adw.SplitButton {
	// tab button
	tabButton := adw.NewSplitButton()
	tabButton.SetIconName("tab-new-symbolic")

	// tab split menu
	tabMenu := gio.NewMenu()
	tabButton.SetMenuModel(tabMenu)
	tabButton.Popover().SetPosition(gtk.PosBottom)
	tabButton.Popover().SetHAlign(gtk.AlignStart)

	// serial tab menu entry
	serialTab := gio.NewMenuItem("New Serial", "win.new-serial")
	tabMenu.InsertItem(0, serialTab)

	// ble tab menu entry
	bleTab := gio.NewMenuItem("New Bluetooth", "win.new-ble")
	tabMenu.InsertItem(1, bleTab)

	return tabButton
}

func newSearchButton() *gtk.Button {
	// search button
	searchButton := gtk.NewButtonFromIconName("search-symbolic")
	return searchButton
}

func newHeaderBar() *adw.HeaderBar {
	header := adw.NewHeaderBar()
	header.SetShowEndTitleButtons(true)

	header.PackStart(newTabMenu())

	header.PackEnd(newMainMenu())

	header.PackEnd(newSearchButton())
	return header
}

func activate(app *gtk.Application) {

	header := newHeaderBar()
	box := gtk.NewBox(gtk.OrientationVertical, 12)

	window := gtk.NewApplicationWindow(app)
	window.SetChild(box)
	//window.SetTitle("/dev/ttyUSB")
	window.SetTitlebar(header)
	window.SetDefaultSize(808, 550)
	window.Show()
}
