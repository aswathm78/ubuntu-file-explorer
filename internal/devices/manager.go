package devices

import (
	"context"
	"fmt"
	"sync"

	"github.com/user/finder-clone/internal/core/event"
)

type Manager struct {
	watcher Watcher
	bus     event.Bus
	mounts  map[string]MountPoint
	mu      sync.RWMutex
}

func NewManager(watcher Watcher, bus event.Bus) *Manager {
	return &Manager{
		watcher: watcher,
		bus:     bus,
		mounts:  make(map[string]MountPoint),
	}
}

func (m *Manager) Start(ctx context.Context) error {
	if err := m.watcher.Start(ctx); err != nil {
		return err
	}

	go m.listen(ctx)
	return nil
}

func (m *Manager) listen(ctx context.Context) {
	ch := m.watcher.Events()
	for {
		select {
		case <-ctx.Done():
			return
		case evt := <-ch:
			m.handleEvent(evt)
		}
	}
}

func (m *Manager) handleEvent(evt DeviceEvent) {
	// Translate low-level device event to high-level application event
	// e.g., Filter out non-removable drives if configured
	m.bus.Publish(event.TopicDeviceAdded, evt)
}

func (m *Manager) ListMounts() ([]MountPoint, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	// Real implementation reads /proc/mounts or uses `mount` command
	// For now return dummy
	points := make([]MountPoint, 0, len(m.mounts))
	for _, mp := range m.mounts {
		points = append(points, mp)
	}
	return points, nil
}

func (m *Manager) Mount(devicePath string) error {
	// Use UDisks2 or 'mount' command
	return fmt.Errorf("not implemented")
}

func (m *Manager) Unmount(mountPath string) error {
	return fmt.Errorf("not implemented")
}

func (m *Manager) Format(devicePath string, fsType string) error {
	return fmt.Errorf("refusing to format in prototype")
}
