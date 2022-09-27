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

func newHeaderBar() *adw.HeaderBar {
	header := adw.NewHeaderBar()
	header.SetShowEndTitleButtons(true)

	// tab button
	tabButton := adw.NewSplitButton()
	tabButton.SetIconName("tab-new-symbolic")
	//tabMenu := gtk.NewBox(gtk.OrientationVertical, 6)
	//pop := gtk.NewPopover()
	//pop.SetChild(tabMenu)
	//tabButton.SetPopover(pop)
	tMenu := gio.NewMenu()
	serTab := gio.NewMenuItem("New Serial", "window.new_tty")
	bleTab := gio.NewMenuItem("New Bluetooth", "window.new_ble")
	tMenu.InsertItem(0, serTab)
	tMenu.InsertItem(1, bleTab)
	tabButton.SetMenuModel(tMenu)
	tabButton.Popover().SetPosition(gtk.PosBottom)
	tabButton.Popover().SetHAlign(gtk.AlignStart)

	header.PackStart(tabButton)

	// main menu button
	mainMenu := gtk.NewPopover()
	mainMenu.SetVisible(false)
	// menuContent := gtk.NewBox(gtk.OrientationVertical, 0)

	// ctrls := gtk.NewBox(gtk.OrientationHorizontal, 0)
	// p := gtk.NewButtonFromIconName("value-increase-symbolic")
	// v := gtk.NewButton()
	// v.SetLabel("100%")
	// m := gtk.NewButtonFromIconName("value-decrease-symbolic")
	// ctrls.Append(m)
	// ctrls.Append(v)
	// ctrls.Append(p)
	// menuContent.Append(ctrls)

	a := gio.NewMenuItem("zoom-out", "app.about")
	a.SetAttributeValue("verb-icon", glib.NewVariantString("zoom-out-symbolic"))
	b := gio.NewMenuItem("100%", "app.about")
	c := gio.NewMenuItem("zoom-in", "app.about")
	c.SetAttributeValue("verb-icon", glib.NewVariantString("zoom-in-symbolic"))

	abc := gio.NewMenu()
	abc.InsertItem(0, a)
	abc.InsertItem(1, b)
	abc.InsertItem(2, c)
	ms := gio.NewMenuItemSection("", abc)
	ms.SetAttributeValue("display-hint", glib.NewVariantString("horizontal-buttons"))

	mm := gio.NewMenu()

	mm.InsertItem(0, ms)

	// Preferences menu entry
	mFullScreen := gio.NewMenuItem("Fullscreen", "window.fullscreen")
	mm.InsertItem(1, mFullScreen)

	// Preferences menu entry
	mPref := gio.NewMenuItem("Preferences", "app.preferences")
	mm.InsertItem(2, mPref)

	// About menu entry
	mAbout := gio.NewMenuItem("About", "app.about")
	mm.InsertItem(3, mAbout)

	//mainMenu.SetChild(menuContent)
	mainButton := gtk.NewMenuButton()
	mainButton.SetVAlign(gtk.AlignCenter)
	mainButton.SetIconName("open-menu-symbolic")
	//mainButton.SetPopover(mainMenu)
	mainButton.SetMenuModel(mm)
	header.PackEnd(mainButton)

	// search button
	searchButton := gtk.NewButtonFromIconName("search-symbolic")
	header.PackEnd(searchButton)
	return header
}

func activate(app *gtk.Application) {

	header := newHeaderBar()
	box := gtk.NewBox(gtk.OrientationVertical, 12)

	window := gtk.NewApplicationWindow(app)
	window.SetChild(box)
	//window.SetTitle("Adwaita Example")
	window.SetTitlebar(header)
	window.SetDefaultSize(808, 550)
	window.Show()
}
