package navigation

import (
	"context"
	"fmt"
	"path/filepath"
	"sync"

	"github.com/user/finder-clone/internal/core/fs"
)

type StackManager struct {
	fs      fs.FileSystem
	columns []ColumnState
	mu      sync.RWMutex
}

func NewStackManager(fileSys fs.FileSystem) *StackManager {
	return &StackManager{
		fs:      fileSys,
		columns: []ColumnState{},
	}
}

func (m *StackManager) Push(ctx context.Context, path string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Load directory contents
	files, err := m.fs.List(ctx, path)
	if err != nil {
		return fmt.Errorf("failed to list directory %s: %w", path, err)
	}

	newCol := ColumnState{
		Path:    path,
		Files:   files,
		Loading: false,
	}

	m.columns = append(m.columns, newCol)
	return nil
}

func (m *StackManager) Pop() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.columns) == 0 {
		return fmt.Errorf("stack is empty")
	}

	m.columns = m.columns[:len(m.columns)-1]
	return nil
}

func (m *StackManager) NavigateTo(ctx context.Context, path string) error {
	m.mu.Lock()
	m.columns = []ColumnState{}
	m.mu.Unlock()

	// Build the column stack by walking up the path
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	// This is a simplified version; in production we'd build the stack properly
	return m.Push(ctx, absPath)
}

func (m *StackManager) CurrentPath() string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if len(m.columns) == 0 {
		return "/"
	}

	last := m.columns[len(m.columns)-1]
	if last.Selected != "" {
		return filepath.Join(last.Path, last.Selected)
	}
	return last.Path
}

func (m *StackManager) GetColumns() []ColumnState {
	m.mu.RLock()
	defer m.mu.RUnlock()

	// Return a copy to avoid data races
	cols := make([]ColumnState, len(m.columns))
	copy(cols, m.columns)
	return cols
}

func (m *StackManager) Select(index int, fileName string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if index < 0 || index >= len(m.columns) {
		return
	}

	m.columns[index].Selected = fileName
	// Truncate stack after this index if navigating into a new folder
	if index < len(m.columns)-1 {
		m.columns = m.columns[:index+1]
	}
}
