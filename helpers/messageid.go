package helpers

import (
	"chat/database"
	"sync"
)

type LastMessageID struct {
	mu        sync.RWMutex
	messageID int
}

func NewLastMessageID() (*LastMessageID, error) {
	mi := &LastMessageID{}
	var err error
	mi.messageID, err = database.GetLastMessageID()

	return mi, err
}

func (mi *LastMessageID) GetMessageID() int {
	mi.mu.Lock()
	defer mi.mu.Unlock()

	return mi.messageID
}

func (mi *LastMessageID) SetMessageID(id int) {
	mi.mu.Lock()
	defer mi.mu.Unlock()

	mi.messageID = id
}
