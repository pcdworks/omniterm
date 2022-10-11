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

const tabCSS = `
#TerminalTab-50{
	font-size: 8px;
}
#TerminalTab-60{
	font-size: 10px;
}
#TerminalTab-70{
	font-size: 12px;
}
#TerminalTab-80{
	font-size: 14px;
}
#TerminalTab-90{
	font-size: 16px;
}
#TerminalTab-100{
	font-size: 18px;
}
#TerminalTab-110{
	font-size: 20px;
}
#TerminalTab-120{
	font-size: 22px;
}
#TerminalTab-130{
	font-size: 24px;
}
#TerminalTab-140{
	font-size: 26px;
}
#TerminalTab-150{
	font-size: 28px;
}
#TerminalTab-160{
	font-size: 30px;
}
#TerminalTab-170{
	font-size: 32px;
}
#TerminalTab-180{
	font-size: 34px;
}
#TerminalTab-190{
	font-size: 36px;
}
#TerminalTab-200{
	font-size: 38px;
}
#TerminalTab-210{
	font-size: 40px;
}
#TerminalTab-220{
	font-size: 42px;
}
#TerminalTab-230{
	font-size: 44px;
}
#TerminalTab-240{
	font-size: 46px;
}
#TerminalTab-250{
	font-size: 48px;
}
#TerminalTab-260{
	font-size: 50px;
}
#TerminalTab-270{
	font-size: 52px;
}
#TerminalTab-280{
	font-size: 54px;
}
#TerminalTab-290{
	font-size: 56px;
}
#TerminalTab-300{
	font-size: 58px;
}
#TerminalTab-310{
	font-size: 60px;
}
#TerminalTab-320{
	font-size: 62px;
}
#TerminalTab-330{
	font-size: 64px;
}
#TerminalTab-340{
	font-size: 66px;
}
#TerminalTab-350{
	font-size: 68px;
}
#TerminalTab-360{
	font-size: 70px;
}
#TerminalTab-370{
	font-size: 72px;
}
#TerminalTab-380{
	font-size: 74px;
}
#TerminalTab-390{
	font-size: 76px;
}
#TerminalTab-400{
	font-size: 78px;
}
`

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
	ts := gtk.NewCSSProvider()
	ts.LoadFromData(tabCSS)
	tv.SetName("TerminalTab-100")
	tv.StyleContext().AddProvider(ts, 1)

	// connect tab
	tab := window.View.AddPage(tt, nil)
	ico := gio.NewThemedIcon("utilities-terminal-symbolic")
	tab.SetIndicatorIcon(ico)
	window.View.SetSelectedPage(tab)

}
