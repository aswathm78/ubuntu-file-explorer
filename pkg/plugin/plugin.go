package plugin

import (
	"context"
)

// Plugin represents an external extension for the file explorer
type Plugin interface {
	ID() string
	Name() string
	Version() string
}

// MenuExtension allows plugins to add items to the context menu
type MenuExtension interface {
	GetMenuItems(paths []string) []MenuItem
}

type MenuItem struct {
	Label    string
	IconName string
	Action   func(ctx context.Context, paths []string) error
}

// PreviewExtension allows plugins to provide custom preview rendering
type PreviewExtension interface {
	CanPreview(mimeType string) bool
	RenderPreview(ctx context.Context, path string) (interface{}, error)
}
