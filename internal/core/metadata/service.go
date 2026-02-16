package metadata

import (
	"context"

	"github.com/user/finder-clone/internal/core/fs"
)

// Service enriches file information with additional metadata
type Service interface {
	// Enrich adds extra metadata to a FileInfo struct (e.g., video duration, image dimensions)
	Enrich(ctx context.Context, info fs.FileInfo) (fs.FileInfo, error)

	// GetIconName returns the icon name for a given file type
	GetIconName(info fs.FileInfo) string
}
