package columns

import (
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/user/finder-clone/internal/core/fs"
)

type Column struct {
	*gtk.ScrolledWindow
	Path     string
	ListView *gtk.ListView
}

func NewColumn(path string, files []fs.FileInfo) *Column {
	scroll := gtk.NewScrolledWindow()
	scroll.SetSizeRequest(250, -1)
	scroll.SetPolicy(gtk.PolicyNever, gtk.PolicyAutomatic)

	// Create a ListStore to hold our data (simplification for now)
	// In a real production app with 100k+ files, we'd use a custom GListModel
	// that loads data on demand.
	store := gio.NewListStore(gio.GTypeObject) // This might need refinement for file objects

	// Factory for creating list items
	factory := gtk.NewSignalListItemFactory()
	factory.ConnectSetup(func(listItem *gtk.ListItem) {
		label := gtk.NewLabel("")
		listItem.SetChild(label)
	})

	factory.ConnectBind(func(listItem *gtk.ListItem) {
		// Bind data to label here
		// obj := listItem.Item()
		// label := listItem.Child().(*gtk.Label)
		// label.SetLabel(obj.Name())
	})

	// Selection model
	selModel := gtk.NewSingleSelection(store)

	listView := gtk.NewListView(selModel, factory)
	scroll.SetChild(listView)

	return &Column{
		ScrolledWindow: scroll,
		Path:           path,
		ListView:       listView,
	}
}
