package fs

import (
	"context"
	"os"
	"path/filepath"
	"strings"
)

type LocalFileSystem struct{}

func NewLocalFileSystem() *LocalFileSystem {
	return &LocalFileSystem{}
}

type localFileInfo struct {
	os.FileInfo
	path string
}

func (l *localFileInfo) MimeType() string {
	// Simple extension based mime detection for now
	return "application/octet-stream"
}

func (l *localFileInfo) IconName() string {
	if l.IsDir() {
		return "folder"
	}
	return "text-x-generic"
}

func (l *localFileInfo) IsHidden() bool {
	return strings.HasPrefix(l.Name(), ".")
}

func (l *localFileInfo) Path() string {
	return l.path
}

func (l *LocalFileSystem) List(ctx context.Context, path string) ([]FileInfo, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var infos []FileInfo
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}
		infos = append(infos, &localFileInfo{FileInfo: info, path: filepath.Join(path, info.Name())})
	}
	return infos, nil
}

func (l *LocalFileSystem) Stat(ctx context.Context, path string) (FileInfo, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	return &localFileInfo{FileInfo: info, path: path}, nil
}

func (l *LocalFileSystem) Watch(ctx context.Context, path string) (<-chan FileEvent, error) {
	// In-memory watch placeholder
	// In production, use fsnotify
	return make(<-chan FileEvent), nil
}

func (l *LocalFileSystem) IsReadOnly() bool {
	return false
}
