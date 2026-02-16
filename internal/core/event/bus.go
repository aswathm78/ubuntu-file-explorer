package event

// Bus defines the interface for the application event bus
type Bus interface {
	// Publish sends an event to all subscribers of the given topic
	Publish(topic string, payload interface{})

	// Subscribe returns a channel that receives events for the given topic
	Subscribe(topic string) <-chan interface{}
}

// Common topics
const (
	TopicDeviceAdded   = "device.added"
	TopicDeviceRemoved = "device.removed"
	TopicMountChanged  = "mount.changed"
	TopicFileChanged   = "file.changed"
	TopicSelection     = "ui.selection"
)
