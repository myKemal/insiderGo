package services

import (
	"sync"

	"github.com/myKemal/insiderGo/app/dtos"
)

type MessageService struct {
	messages []dtos.MessageDTO
	mu       sync.RWMutex
}

var instance *MessageService
var once sync.Once

func GetMessageService() *MessageService {
	once.Do(func() {
		instance = &MessageService{}
	})
	return instance
}

func (s *MessageService) SetMessages(messages []dtos.MessageDTO) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.messages = messages
}

func (s *MessageService) GetMessages() []dtos.MessageDTO {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.messages
}
