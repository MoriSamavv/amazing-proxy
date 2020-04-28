package message

type GfwMessage {
	data string
}

func (gfwMessage *GfwMessage) MessageType() string {
	return GFW;
}
