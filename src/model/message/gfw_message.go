package message

type GfwMessage struct {
	Data string
}

func (m *GfwMessage) GetMessageType() int {
	return GFW;
}

func (m *GfwMessage) GetData() string {
	return m.Data;
}
