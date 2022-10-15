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

	"pcdworks.com/omniterm/internal/omniterm"
)

func main() {
	tapp := omniterm.NewTerminalApplication()

	if code := tapp.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}
