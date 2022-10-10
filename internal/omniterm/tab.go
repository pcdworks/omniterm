package omniterm

import (
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/diamondburned/gotkit/components/autoscroll"
)

type TerminalTab struct {
	*gtk.Box
	Content   *gtk.Box
	SearchBar *gtk.SearchBar
	Settings  *gtk.Box
}

func (window *TerminalWindow) NewSerialTab() {
	tt := &TerminalTab{
		Box:      gtk.NewBox(gtk.OrientationVertical, 0),
		Content:  gtk.NewBox(gtk.OrientationVertical, 0),
		Settings: gtk.NewBox(gtk.OrientationVertical, 0),
	}

	as := autoscroll.NewWindow()
	as.SetVExpand(true)
	as.SetMarginStart(2)
	as.SetMarginEnd(2)
	tv := gtk.NewTextView()
	tv.SetVExpand(true)
	as.SetChild(tv)
	tt.Content.Append(as)

	tx := gtk.NewBox(gtk.OrientationVertical, 0)
	tx.Append(tt)
	tab := window.View.AddPage(tx, nil)
	//tab.SetTitle("/dev/ttyUSB0")
	ico := gio.NewThemedIcon("utilities-terminal-symbolic")
	tab.SetIndicatorIcon(ico)
	window.View.SetSelectedPage(tab)
	tv.GrabFocus()

}

// func (tt *TerminalTab) BaseWidget() *gtk.Widget {
// 	return &tt.Content.Widget
// }

// func NewTerminalTab() *TerminalTab {
// 	tt := &TerminalTab{}
// 	tt.Content = gtk.NewBox(gtk.OrientationVertical, 0)

// 	// Search
// 	tt.SearchBar = gtk.NewSearchBar()
// 	tt.SearchBar.SetHExpand(true)
// 	tt.SearchBar.SetSearchMode(false)
// 	clamp := adw.NewClamp()
// 	clamp.SetHExpand(true)
// 	searchBox := gtk.NewBox(gtk.OrientationHorizontal, 0)
// 	searchBox.AddCSSClass("linked")
// 	sentry := gtk.NewSearchEntry()
// 	sentry.SetHExpand(true)
// 	searchBox.Append(sentry)
// 	next := gtk.NewButtonFromIconName("go-up-symbolic")
// 	prev := gtk.NewButtonFromIconName("go-down-symbolic")
// 	searchBox.Append(next)
// 	searchBox.Append(prev)
// 	clamp.SetChild(searchBox)
// 	tt.SearchBar.SetChild(clamp)
// 	tt.Content.Append(tt.SearchBar)

// 	// Scroll bar
// 	as := autoscroll.NewWindow()
// 	as.SetVExpand(true)
// 	as.SetMarginStart(2)
// 	as.SetMarginEnd(2)
// 	tt.Content.Append(as)

// 	// Text view
// 	tv := gtk.NewTextView()
// 	tv.SetVExpand(true)
// 	as.SetChild(tv)

// 	return tt
// }
