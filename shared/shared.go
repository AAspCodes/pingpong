package shared

import (
	"fmt"
)

type Message_struct struct {
	Sender         string
	SequenceNumber int
	Data           string
}

func (m *Message_struct) ToString() string {
	return fmt.Sprintf("Sender: %s, SequenceNumber: %d, Data: %s", m.Sender, m.SequenceNumber, m.Data)
}
