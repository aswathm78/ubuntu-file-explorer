package preview

import (
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/user/finder-clone/internal/core/fs"
)

type Panel struct {
	*gtk.Box
	InfoLabel *gtk.Label
	Image     *gtk.Picture
}

func NewPanel() *Panel {
	box := gtk.NewBox(gtk.OrientationVertical, 10)
	box.SetSizeRequest(300, -1)
	box.SetHExpand(false)
	box.SetVExpand(true)

	label := gtk.NewLabel("Select a file to preview")
	label.SetWrap(true)
	box.Append(label)

	img := gtk.NewPicture()
	box.Append(img)

	return &Panel{
		Box:       box,
		InfoLabel: label,
		Image:     img,
	}
}

func (p *Panel) Update(info fs.FileInfo) {
	p.InfoLabel.SetLabel(info.Name())
	// In production, load image async and update Picture
}
