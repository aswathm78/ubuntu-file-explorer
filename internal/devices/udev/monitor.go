package udisks2

import (
	"context"
	"fmt"
	"sync"

	"github.com/godbus/dbus/v5"
	"github.com/user/finder-clone/internal/devices"
)

const (
	uDisks2Interface = "org.freedesktop.UDisks2"
	uDisks2Path      = "/org/freedesktop/UDisks2"
)

type Monitor struct {
	conn   *dbus.Conn
	events chan devices.DeviceEvent
	stop   chan struct{}
	wg     sync.WaitGroup
}

func NewMonitor() *Monitor {
	return &Monitor{
		events: make(chan devices.DeviceEvent, 10),
		stop:   make(chan struct{}),
	}
}

func (m *Monitor) Start(ctx context.Context) error {
	conn, err := dbus.SystemBus()
	if err != nil {
		return fmt.Errorf("failed to connect to system bus: %v", err)
	}
	m.conn = conn

	// Subscribe to detailed signals if needed, but for now we look for InterfaceAdded/Removed
	// on the ObjectManager interface which UDisks2 implements.
	if err := m.conn.AddMatchSignal(
		dbus.WithMatchInterface("org.freedesktop.DBus.ObjectManager"),
		dbus.WithMatchSender(uDisks2Interface),
	); err != nil {
		conn.Close()
		return fmt.Errorf("failed to add dbus match: %v", err)
	}

	c := make(chan *dbus.Signal, 10)
	m.conn.Signal(c)

	m.wg.Add(1)
	go m.worker(ctx, c)

	// Initial scan (simplified for this implementation)
	// In production, we would query the ObjectManager for existing devices here.
	go m.scanExisting()

	return nil
}

func (m *Monitor) Events() <-chan devices.DeviceEvent {
	return m.events
}

func (m *Monitor) worker(ctx context.Context, signals chan *dbus.Signal) {
	defer m.wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case <-m.stop:
			return
		case sig := <-signals:
			m.handleSignal(sig)
		}
	}
}

func (m *Monitor) handleSignal(sig *dbus.Signal) {
	switch sig.Name {
	case "org.freedesktop.DBus.ObjectManager.InterfacesAdded":
		// Provide real parsing logic here for production
		// Body[0] is ObjectPath, Body[1] is map[string]map[string]Variant
		if len(sig.Body) >= 2 {
			path, ok := sig.Body[0].(dbus.ObjectPath)
			if ok {
				// Filter for Block devices or Filesystems
				m.events <- devices.DeviceEvent{
					DevicePath: string(path),
					Action:     devices.ActionAdd,
					Type:       devices.TypeUSB, // Simplification
				}
			}
		}
	case "org.freedesktop.DBus.ObjectManager.InterfacesRemoved":
		if len(sig.Body) >= 1 {
			path, ok := sig.Body[0].(dbus.ObjectPath)
			if ok {
				m.events <- devices.DeviceEvent{
					DevicePath: string(path),
					Action:     devices.ActionRemove,
				}
			}
		}
	}
}

func (m *Monitor) scanExisting() {
	// Query UDisks2 Mananger for GetManagedObjects
	// This would populate initial state
}

func (m *Monitor) Stop() {
	close(m.stop)
	m.wg.Wait()
	if m.conn != nil {
		m.conn.Close()
	}
}
