package fs

import (
	"context"
	"os"
	// "time"
)

// FileInfo extends os.FileInfo with additional metadata needed for the UI
type FileInfo interface {
	os.FileInfo
	MimeType() string
	IconName() string
	IsHidden() bool
	Path() string
}

// FileEvent represents a filesystem event (create, modify, delete)
type FileEvent struct {
	Path string
	Type EventType
}

type EventType int

const (
	EventCreate EventType = iota
	EventModify
	EventDelete
	EventRename
)

// FileSystem abstracts local and remote file systems
type FileSystem interface {
	// List returns the contents of a directory
	List(ctx context.Context, path string) ([]FileInfo, error)

	// Stat returns details for a single file
	Stat(ctx context.Context, path string) (FileInfo, error)

	// Watch starts monitoring a directory for changes
	Watch(ctx context.Context, path string) (<-chan FileEvent, error)

	// IsReadOnly returns true if the filesystem cannot be modified
	IsReadOnly() bool
}
