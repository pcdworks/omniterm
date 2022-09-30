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

type TerminalApplication struct {
	*gtk.Application
	Windows []*TerminalWindow
}

type TerminalWindow struct {
	*adw.Window
	View           *adw.TabView
	TabBar         *adw.TabBar
	TabActionGroup *gio.ActionMap
	Page           *adw.TabPage
	HeaderBar      *adw.HeaderBar
}

func NewTerminalApplication() *TerminalApplication {
	app := TerminalApplication{
		Application: gtk.NewApplication("com.pcdworks.omniterm", 0),
	}
	app.Connect("activate", app.activate)
	return &app
}

func (app *TerminalApplication) NewWindow() *TerminalWindow {
	window := TerminalWindow{
		Window:    adw.NewWindow(),
		HeaderBar: adw.NewHeaderBar(),
		TabBar:    adw.NewTabBar(),
		View:      adw.NewTabView(),
	}

	// ***************************
	// Header bar
	// ***************************
	window.HeaderBar.SetShowEndTitleButtons(true)

	// *************************************
	// tab button
	// *************************************
	tabButton := adw.NewSplitButton()
	tabButton.SetIconName("tab-new-symbolic")
	tabButton.ConnectClicked(func() { window.NewTab() })

	// tab split menu
	tabMenu := gio.NewMenu()
	tabButton.SetMenuModel(tabMenu)
	tabButton.Popover().SetPosition(gtk.PosBottom)
	//tabButton.Popover().SetHAlign(gtk.AlignStart)

	// serial tab menu entry
	serialTab := gio.NewMenuItem("New Serial", "win.new-serial")
	tabMenu.InsertItem(0, serialTab)

	// ble tab menu entry
	bleTab := gio.NewMenuItem("New Bluetooth", "win.new-ble")
	tabMenu.InsertItem(1, bleTab)
	window.HeaderBar.PackStart(tabButton)

	// *************************************
	// main menu button
	// *************************************
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
	window.HeaderBar.PackEnd(mainButton)

	// **************************************
	// search button
	// **************************************
	searchButton := gtk.NewButtonFromIconName("search-symbolic")
	window.HeaderBar.PackEnd(searchButton)

	// ***********************
	// Tab context
	// ***********************
	tabContext := gio.NewMenu()

	// Preferences menu entry
	tabClose := gio.NewMenuItem("Close", "tab.close")
	tabContext.InsertItem(1, tabClose)

	// ***************************
	// Window content
	// ***************************
	window.View.SetVExpand(true)
	window.View.SetMenuModel(tabContext)
	window.SetApplication(app.Application)
	window.SetTitle("OmniTerm")
	window.TabBar.SetView(window.View)
	box := gtk.NewBox(gtk.OrientationVertical, 0)
	box.Append(window.HeaderBar)
	box.Append(window.TabBar)
	box.Append(window.View)
	window.SetContent(box)

	window.SetDefaultSize(808, 550)

	window.NewTab()
	window.NewTab()

	return &window
}

func main() {
	app := NewTerminalApplication()

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

func (app *TerminalApplication) activate(self *gtk.Application) {
	win := app.NewWindow()
	win.Show()
}

func (window *TerminalWindow) NewTab() {
	content := gtk.NewBox(gtk.OrientationHorizontal, 0)
	tab := window.View.AddPage(content, &adw.TabPage{})
	//tab.SetTitle("/dev/ttyUSB0")
	//ico := gio.NewThemedIcon("bluetooth-symbolic")
	ico := gio.NewThemedIcon("utilities-terminal-symbolic")
	tab.SetIndicatorIcon(ico)
	window.View.SetSelectedPage(tab)
}
