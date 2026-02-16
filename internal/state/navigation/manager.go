package navigation

import (
	"context"

	"github.com/user/finder-clone/internal/core/fs"
)

// ColumnState represents the state of a single column in the UI
type ColumnState struct {
	Path      string
	Selected  string
	ScrollPos float64
	Files     []fs.FileInfo
	Loading   bool
	Error     error
}

// ColumnManager manages the stack of columns for navigation
type ColumnManager interface {
	// Push adds a new column to the stack (navigating into a folder)
	Push(ctx context.Context, path string) error

	// Pop removes the last column (navigating back)
	Pop() error

	// NavigateTo resets the stack to a specific path
	NavigateTo(ctx context.Context, path string) error

	// CurrentPath returns the full path of the currently selected item
	CurrentPath() string

	// GetColumns returns the current stack of columns
	GetColumns() []ColumnState
}
