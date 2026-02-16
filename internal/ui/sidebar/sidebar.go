package sidebar

import (
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type Sidebar struct {
	*gtk.Box
}

func NewSidebar() *Sidebar {
	box := gtk.NewBox(gtk.OrientationVertical, 0)
	box.SetSizeRequest(200, -1)
	box.SetHExpand(false)
	// box.AddCSSClass("sidebar") // Commented out until we have CSS loader

	box.Append(gtk.NewLabel("PLACES"))
	box.Append(gtk.NewButtonWithLabel("Home"))
	box.Append(gtk.NewButtonWithLabel("Documents"))
	box.Append(gtk.NewButtonWithLabel("Downloads"))

	return &Sidebar{Box: box}
}
