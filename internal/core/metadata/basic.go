package metadata

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/user/finder-clone/internal/core/fs"
)

type BasicMetadataService struct{}

func NewBasicMetadataService() *BasicMetadataService {
	return &BasicMetadataService{}
}

// Enrich currently just passes through, but will later extract video/image data
func (s *BasicMetadataService) Enrich(ctx context.Context, info fs.FileInfo) (fs.FileInfo, error) {
	// For now, no enrichment. Return as is.
	return info, nil
}

// GetIconName returns a standard icon name based on file extension/type
func (s *BasicMetadataService) GetIconName(info fs.FileInfo) string {
	if info.IsDir() {
		return "folder"
	}

	ext := strings.ToLower(filepath.Ext(info.Name()))
	switch ext {
	// Images
	case ".png", ".jpg", ".jpeg", ".gif", ".webp":
		return "image-x-generic"
	// Videos
	case ".mp4", ".mkv", ".avi", ".mov":
		return "video-x-generic"
	// Audio
	case ".mp3", ".wav", ".flac", ".ogg":
		return "audio-x-generic"
	// Documents
	case ".pdf":
		return "application-pdf"
	case ".txt", ".md":
		return "text-x-generic"
	case ".go", ".c", ".cpp", ".py", ".js", ".ts", ".html", ".css":
		return "text-x-script"
	case ".zip", ".tar", ".gz", ".7z", ".rar":
		return "package-x-generic"
	case ".iso":
		return "application-x-cd-image"
	}

	return "text-x-generic" // Default fallback
}
