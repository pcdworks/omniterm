package omniterm

import (
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"go.bug.st/serial/enumerator"
)

func NewBaudSelector() *gtk.ComboBoxText {
	bauds := []string{
		"300", "600", "1200", "2400", "4800", "9600",
		"19200", "38400", "57600", "115200", "230400",
		"460800", "576000", "921600", "1000000", "2000000",
	}
	bb := gtk.NewComboBoxText()
	for _, b := range bauds {
		bb.Append(b, b)
	}
	bb.SetActiveID("115200")
	return bb
}

func NewParity() *gtk.ComboBoxText {
	parity := []string{"none", "odd", "even"}
	pb := gtk.NewComboBoxText()
	for _, p := range parity {
		pb.Append(p, p)
	}
	pb.SetActiveID("none")
	return pb
}

func NewBits() *gtk.ComboBoxText {
	bits := []string{"5", "6", "7", "8"}
	bb := gtk.NewComboBoxText()
	for _, b := range bits {
		bb.Append(b, b)
	}
	bb.SetActiveID("8")
	return bb
}

func NewStopBits() *gtk.ComboBoxText {
	stopBits := []string{"1", "2"}
	sb := gtk.NewComboBoxText()
	for _, s := range stopBits {
		sb.Append(s, s)
	}
	sb.SetActiveID("1")
	return sb
}

func NewPorts() *gtk.ComboBoxText {
	pb := gtk.NewComboBoxText()
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		return pb
	}
	for _, p := range ports {
		pb.Append(p.Name, p.Name)
	}
	return pb
}
