package omniterm

import (
	"strconv"
	"strings"

	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/diamondburned/gotkit/gtkutil"
)

type TerminalWindow struct {
	*adw.Window
	View      *adw.TabView
	TabBar    *adw.TabBar
	Page      *adw.TabPage
	HeaderBar *adw.HeaderBar
}

func (app *TerminalApplication) NewWindow() *TerminalWindow {
	window := TerminalWindow{
		Window:    adw.NewWindow(),
		HeaderBar: adw.NewHeaderBar(),
		TabBar:    adw.NewTabBar(),
		View:      adw.NewTabView(),
	}
	// *********************
	// Window actions
	// *********************
	gtkutil.BindActionMap(window, map[string]func(){
		"win.new-tab":     func() { window.NewTab() },
		"win.tab-config":  func() { window.TabConfigMode() },
		"win.fullscreen":  func() { window.FullscreenMode() },
		"win.zoom-in":     func() { window.ZoomIn() },
		"win.zoom-out":    func() { window.ZoomOut() },
		"win.zoom-normal": func() { window.ZoomNormal() },
	})

	// ***************************
	// Header bar
	// ***************************
	window.HeaderBar.SetShowEndTitleButtons(true)

	// New tab
	tabButton := gtk.NewButtonFromIconName("tab-new-symbolic")
	window.HeaderBar.PackStart(tabButton)
	tabButton.ConnectClicked(window.NewTab)

	// tab preferences
	prefButton := gtk.NewButtonFromIconName("document-edit-symbolic")
	window.HeaderBar.PackStart(prefButton)
	prefButton.ConnectClicked(window.TabConfigMode)

	// zoButton := gtk.NewButtonFromIconName("zoom-out-symbolic")
	// window.HeaderBar.PackStart(zoButton)
	// zoButton.ConnectClicked(window.ZoomOut)

	// ziButton := gtk.NewButtonFromIconName("zoom-in-symbolic")
	// window.HeaderBar.PackStart(ziButton)
	// ziButton.ConnectClicked(window.ZoomIn)

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
	searchButton.ConnectClicked(func() { window.SearchMode() })
	window.HeaderBar.PackEnd(searchButton)

	// ***************************
	// Window content
	// ***************************
	window.View.SetVExpand(true)
	window.View.SetMenuModel(gtkutil.MenuPair([][2]string{
		{"_Move to New Window", "tab.move-to-new-window"},
		{"D_uplicate", "tab.duplicate"},
		{"P_in Tab", "tab.pin"},
		{"Unp_in Tab", "tab.unpin"},
		{"Close Tabs to the _Left", "tab.close-before"},
		{"Close Tabs to the _Right", "tab.close-after"},
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

	window.SetDefaultSize(714, 478)

	window.NewTab()

	return &window
}

func (window *TerminalWindow) SearchMode() {
	page := window.View.SelectedPage()
	if page != nil {
		tab := page.Child().Cast().(*gtk.Box)
		content := tab.LastChild().Cast().(*gtk.Box)
		search := content.FirstChild().Cast().(*gtk.SearchBar)

		search.SetSearchMode(!search.SearchMode())
	}
}

func (window *TerminalWindow) GetTextView() *gtk.TextView {
	page := window.View.SelectedPage()
	if page != nil {
		tab := page.Child().Cast().(*gtk.Box)
		content := tab.LastChild().Cast().(*gtk.Box)
		win := content.LastChild().Cast().(*gtk.ScrolledWindow)
		text := win.Child().Cast().(*gtk.TextView)

		return text
	} else {
		return nil
	}
}

func (window *TerminalWindow) TabConfigMode() {
	page := window.View.SelectedPage()
	if page != nil {
		tab := page.Child().Cast().(*gtk.Box)
		settings := tab.FirstChild().Cast().(*gtk.Box)
		content := tab.LastChild().Cast().(*gtk.Box)
		v := settings.Visible()

		// make sure content and settings are not shown at the same time
		if v {
			settings.Hide()
			content.Show()
			window.GetTextView().GrabFocus()
		} else {
			content.Hide()
			settings.Show()
		}

	}
}

func (window *TerminalWindow) ZoomIn() {
	text := window.GetTextView()
	size, _ := strconv.ParseInt(strings.Split(text.Name(), "-")[1], 10, 0)
	if size < 400 {
		size += 10
		ssize := strconv.FormatInt(size, 10)
		text.SetName("TerminalTab-" + ssize)
	}
}

func (window *TerminalWindow) ZoomOut() {
	text := window.GetTextView()
	size, _ := strconv.ParseInt(strings.Split(text.Name(), "-")[1], 10, 0)
	if size > 60 {
		size -= 10
		ssize := strconv.FormatInt(size, 10)
		text.SetName("TerminalTab-" + ssize)

	}
}

func (window *TerminalWindow) ZoomNormal() {
	text := window.GetTextView()
	text.SetName("TerminalTab-100")
}

func (window *TerminalWindow) FullscreenMode() {
	if window.IsFullscreen() {
		window.Unfullscreen()
	} else {
		window.Fullscreen()
	}
}
