package main

import (
	"os"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/user/finder-clone/internal/core/event"
	"github.com/user/finder-clone/internal/core/fs"
	"github.com/user/finder-clone/internal/state/navigation"
	"github.com/user/finder-clone/internal/ui"
)

func main() {
	app := gtk.NewApplication("com.user.finder-clone", gio.ApplicationFlagsNone)
	app.ConnectActivate(func() {
		activate(app)
	})

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

func activate(app *gtk.Application) {
	// Initialize core components
	bus := event.NewMemoryBus()
	fileSys := fs.NewLocalFileSystem()
	stack := navigation.NewStackManager(fileSys)

	// Create main window
	window := ui.NewMainWindow(app, bus, stack)
	window.Show()
}
