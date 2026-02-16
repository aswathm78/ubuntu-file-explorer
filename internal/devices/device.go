package devices

import "context"

// DeviceEvent represents a hardware event (USB inserted/removed)
type DeviceEvent struct {
	DevicePath string
	Action     DeviceAction
	Type       DeviceType
}

type DeviceAction int
type DeviceType int

const (
	ActionAdd DeviceAction = iota
	ActionRemove
)

const (
	TypeUSB DeviceType = iota
	TypeSD
	TypeMTP
)

// Watcher monitors hardware changes (udev)
type Watcher interface {
	Start(ctx context.Context) error
	Events() <-chan DeviceEvent
}

// MountPoint represents a mounted filesystem
type MountPoint struct {
	DevicePath string
	MountPath  string
	FSType     string
	Options    []string
}

// MountManager handles mounting and unmounting of devices
type MountManager interface {
	ListMounts() ([]MountPoint, error)
	Mount(devicePath string) error
	Unmount(mountPath string) error
	Format(devicePath string, fsType string) error
}
