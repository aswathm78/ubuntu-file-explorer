package ui

import (
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/user/finder-clone/internal/core/event"
	"github.com/user/finder-clone/internal/state/navigation"
	"github.com/user/finder-clone/internal/ui/columns"
	"github.com/user/finder-clone/internal/ui/preview"
	"github.com/user/finder-clone/internal/ui/sidebar"
)

type MainWindow struct {
	*gtk.ApplicationWindow
	bus       event.Bus
	stack     navigation.ColumnManager
	columnBox *gtk.Box
}

func NewMainWindow(app *gtk.Application, bus event.Bus, stack navigation.ColumnManager) *MainWindow {
	win := gtk.NewApplicationWindow(app)
	win.SetTitle("Finder Clone")
	win.SetDefaultSize(1000, 600)

	mw := &MainWindow{
		ApplicationWindow: win,
		bus:               bus,
		stack:             stack,
	}

	mw.setupUI()
	mw.renderColumns()
	return mw
}

func (mw *MainWindow) setupUI() {
	// Root Layout
	box := gtk.NewBox(gtk.OrientationHorizontal, 0)
	mw.SetChild(box)

	// Sidebar
	sb := sidebar.NewSidebar()
	box.Append(sb)

	// Separator
	box.Append(gtk.NewSeparator(gtk.OrientationVertical))

	// Main Content (Horizontal Scrollable Column View)
	scroll := gtk.NewScrolledWindow()
	scroll.SetHExpand(true)
	scroll.SetVExpand(true)
	scroll.SetPolicy(gtk.PolicyAlways, gtk.PolicyNever)
	box.Append(scroll)

	// Horizontal box to hold columns
	mw.columnBox = gtk.NewBox(gtk.OrientationHorizontal, 0)
	scroll.SetChild(mw.columnBox)

	// Separator
	box.Append(gtk.NewSeparator(gtk.OrientationVertical))

	// Preview Panel
	prev := preview.NewPanel()
	box.Append(prev)

	// TODO: Connect navigation events to columnView
}

func (mw *MainWindow) renderColumns() {
	// Clear existing columns
	for child := mw.columnBox.FirstChild(); child != nil; {
		w := child.(interface{ NextSibling() *gtk.Widget }).NextSibling()
		mw.columnBox.Remove(child)
		child = w
	}

	// Get columns from stack
	cols := mw.stack.GetColumns()
	for _, colState := range cols {
		col := columns.NewColumn(colState.Path, colState.Files)
		mw.columnBox.Append(col)
	}
}
