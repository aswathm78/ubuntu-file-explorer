package ops

import (
	"context"
)

// Progress reports the status of a long-running operation
type Progress struct {
	CurrentFile     string
	TotalBytes      int64
	BytesDone       int64
	TotalFiles      int
	FilesDone       int
	Error           error
	IsIndeterminate bool
}

// Operator handles file manipulation operations
type Operator interface {
	// Copy copies a file or directory recursively
	Copy(ctx context.Context, src, dest string) (<-chan Progress, error)

	// Move moves a file or directory
	Move(ctx context.Context, src, dest string) (<-chan Progress, error)

	// Delete permanently deletes files
	Delete(ctx context.Context, paths []string) error

	// Trash moves files to the trash can
	Trash(ctx context.Context, paths []string) error
}
