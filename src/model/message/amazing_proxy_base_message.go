package message

// all message type
const (
	GFW = 0
)

type AmazingProxyBaseMessage interface {
	GetMessageType() int
	GetData() string
}
