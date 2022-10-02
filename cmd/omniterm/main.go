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
	"context"
	"os"

	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/diamondburned/gotkit/app"
	"github.com/diamondburned/gotkit/gtkutil"
)

type TerminalApplication struct {
	*app.Application
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
	tapp := TerminalApplication{
		Application: app.New(context.Background(), "com.pcdworks.omniterm", "omniterm"),
	}
	tapp.Connect("activate", tapp.activate)
	tapp.AddActions(map[string]func(){
		"app.new-window":  func() { tapp.NewWindow().Show() },
		"app.preferences": func() {},
		"app.about":       func() {},
	})
	return &tapp
}

func (app *TerminalApplication) NewWindow() *TerminalWindow {
	window := TerminalWindow{
		Window:         adw.NewWindow(),
		HeaderBar:      adw.NewHeaderBar(),
		TabBar:         adw.NewTabBar(),
		View:           adw.NewTabView(),
		TabActionGroup: &gio.NewSimpleActionGroup().ActionMap,
	}

	// *********************
	// Window actions
	// *********************
	gtkutil.BindActionMap(window, map[string]func(){
		"win.new-serial-tab": func() { window.NewSerialTab() },
		"win.new-ble-tab":    func() { window.NewBLETab() },
		"win.fullscreen":     func() { window.FullscreenCB() },
		"win.zoom-in":        func() { window.ZoomIn() },
		"win.zoom-out":       func() { window.ZoomOut() },
		"win.zoom-normal":    func() { window.ZoomNormal() },
	})

	// ***************************
	// Header bar
	// ***************************
	window.HeaderBar.SetShowEndTitleButtons(true)

	// *************************************
	// tab button
	// *************************************
	tabButton := adw.NewSplitButton()
	tabButton.SetIconName("tab-new-symbolic")
	tabButton.ConnectClicked(func() {
		window.NewTab()
	})

	// tab split menu
	tabButton.SetMenuModel(gtkutil.MenuPair([][2]string{
		{"New Serial", "win.new-serial-tab"},
		{"New Bluetooth", "win.new-ble-tab"},
	}))
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

	// menu button
	mainButton := gtk.NewMenuButton()
	mainButton.SetVAlign(gtk.AlignCenter)
	mainButton.SetIconName("open-menu-symbolic")
	list := gtkutil.MenuPair([][2]string{
		{"New Window", "app.new-window"},
		{"Fullscreen", "win.fullscreen"},
		{"Preferences", "app.preferences"},
		{"About", "app.about"},
	})
	mainMenu.InsertSection(1, "", list)
	mainButton.SetMenuModel(mainMenu)
	window.HeaderBar.PackEnd(mainButton)

	// **************************************
	// search button
	// **************************************
	searchButton := gtk.NewButtonFromIconName("search-symbolic")
	window.HeaderBar.PackEnd(searchButton)

	// ***************************
	// Window content
	// ***************************
	window.View.SetVExpand(true)
	window.View.SetMenuModel(gtkutil.MenuPair([][2]string{
		{"Close", "tab.close"},
	}))
	window.SetApplication(app.Application.Application)
	window.SetTitle("OmniTerm")
	window.TabBar.SetView(window.View)
	box := gtk.NewBox(gtk.OrientationVertical, 0)
	box.Append(window.HeaderBar)
	box.Append(window.TabBar)
	box.Append(window.View)
	window.SetContent(box)

	window.SetDefaultSize(808, 550)

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

func (window *TerminalWindow) ZoomIn() {
}

func (window *TerminalWindow) ZoomOut() {
}

func (window *TerminalWindow) ZoomNormal() {
}

func (window *TerminalWindow) FullscreenCB() {
	if window.IsFullscreen() {
		window.Unfullscreen()
	} else {
		window.Fullscreen()
	}
}

func (window *TerminalWindow) NewTab() {
	pageType := "utilities-terminal-symbolic"
	if window.View.SelectedPage() != nil {
		pageType = window.View.SelectedPage().IndicatorIcon().String()
	}

	if pageType == "bluetooth-symbolic" {
		window.NewBLETab()
	} else if pageType == "utilities-terminal-symbolic" {
		window.NewSerialTab()
	}
}

func (window *TerminalWindow) NewSerialTab() {
	content := gtk.NewBox(gtk.OrientationHorizontal, 0)
	tab := window.View.AddPage(content, &adw.TabPage{})
	//tab.SetTitle("/dev/ttyUSB0")
	ico := gio.NewThemedIcon("utilities-terminal-symbolic")
	tab.SetIndicatorIcon(ico)
	window.View.SetSelectedPage(tab)
}

func (window *TerminalWindow) NewBLETab() {
	content := gtk.NewBox(gtk.OrientationHorizontal, 0)
	tab := window.View.AddPage(content, &adw.TabPage{})
	//tab.SetTitle("/dev/ttyUSB0")
	ico := gio.NewThemedIcon("bluetooth-symbolic")
	tab.SetIndicatorIcon(ico)
	window.View.SetSelectedPage(tab)
}
