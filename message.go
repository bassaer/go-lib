package mylib

import (
	"fmt"
)

type Message struct {
	content string
}

func NewMessage(c string) *Message {
	return &Message{c}
}

func (m *Message) Send() {
	fmt.Println("send: ", m.content)
}
