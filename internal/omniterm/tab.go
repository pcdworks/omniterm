package omniterm

import (
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
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

func (window *TerminalWindow) NewBLETab() {
	tt := gtk.NewBox(gtk.OrientationVertical, 0)

	// connect tab
	tab := window.View.AddPage(tt, nil)
	ico := gio.NewThemedIcon("bluetooth-symbolic")
	tab.SetIndicatorIcon(ico)
	window.View.SetSelectedPage(tab)
}

func (window *TerminalWindow) NewSerialTab() {
	tt := gtk.NewBox(gtk.OrientationVertical, 0)

	// settings area
	settings := gtk.NewBox(gtk.OrientationVertical, 0)
	settings.SetVExpand(true)
	tt.Append(settings)
	settings.Append(gtk.NewButtonWithLabel("Settings"))
	settings.Hide()

	// content area
	content := gtk.NewBox(gtk.OrientationVertical, 0)
	content.SetVExpand(true)
	tt.Append(content)

	// Search
	sb := gtk.NewSearchBar()
	sb.SetHExpand(true)
	sb.SetSearchMode(false)
	clamp := adw.NewClamp()
	clamp.SetHExpand(true)
	searchBox := gtk.NewBox(gtk.OrientationHorizontal, 0)
	searchBox.AddCSSClass("linked")
	sentry := gtk.NewSearchEntry()
	sentry.SetHExpand(true)
	searchBox.Append(sentry)
	next := gtk.NewButtonFromIconName("go-up-symbolic")
	prev := gtk.NewButtonFromIconName("go-down-symbolic")
	searchBox.Append(next)
	searchBox.Append(prev)
	clamp.SetChild(searchBox)
	sb.SetChild(clamp)
	content.Append(sb)

	// Scroll bar
	as := autoscroll.NewWindow()
	as.SetVExpand(true)
	as.SetMarginStart(2)
	as.SetMarginEnd(2)
	content.Append(as)

	// Text view
	tv := gtk.NewTextView()
	tv.SetVExpand(true)
	as.SetChild(tv)

	// connect tab
	tab := window.View.AddPage(tt, nil)
	ico := gio.NewThemedIcon("utilities-terminal-symbolic")
	tab.SetIndicatorIcon(ico)
	window.View.SetSelectedPage(tab)

}
