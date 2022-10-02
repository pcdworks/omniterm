package omniterm

import (
	"context"

	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/diamondburned/gotkit/app"
)

type TerminalApplication struct {
	*app.Application
	Windows []*TerminalWindow
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

func (app *TerminalApplication) activate(self *gtk.Application) {
	win := app.NewWindow()
	win.Show()
}
