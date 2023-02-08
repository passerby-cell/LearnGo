package mock

type Receiver struct {
	Content string
}

func (receiver Receiver) GetRes(s string) string {
	return receiver.Content
}
