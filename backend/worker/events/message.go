package events

type Message struct {
	Body          string
	EventType     string
	ReceiptHandle string
}
