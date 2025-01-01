package services

import (
	"github.com/myKemal/insiderGo/app/model"
	"github.com/myKemal/insiderGo/app/repository"
	"log"
	"sync"
	"time"
)

type PeriodicTaskService struct {
	TempRepository  repository.TempMemoryRepository
	WebhookService  *WebhookService
	InstanceMessage *MessageService
	Interval        time.Duration
	stopChannel     chan bool
	isRunning       bool
	mu              sync.Mutex
}

func NewPeriodicTaskService(instanceMessage *MessageService, tempRepo repository.TempMemoryRepository, webhookService *WebhookService, interval time.Duration) *PeriodicTaskService {
	return &PeriodicTaskService{
		TempRepository:  tempRepo,
		WebhookService:  webhookService,
		InstanceMessage: instanceMessage,
		Interval:        interval,
		stopChannel:     make(chan bool),
		isRunning:       false,
	}
}

func (p *PeriodicTaskService) Start() (bool, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.isRunning {
		log.Println("PeriodicTaskService is already running.")
		return false, nil
	}

	p.isRunning = true
	log.Println("PeriodicTaskService started...")
	p.processMessages()

	go func() {
		ticker := time.NewTicker(p.Interval)
		defer func() {
			ticker.Stop()
			p.mu.Lock()
			p.isRunning = false
			p.mu.Unlock()
		}()

		for {
			select {
			case <-ticker.C:
				log.Println("Executing periodic webhook sender...")
				p.processMessages()
			case <-p.stopChannel:
				log.Println("PeriodicTaskService stopped.")
				return
			}
		}
	}()

	return true, nil
}

func (p *PeriodicTaskService) Stop() (bool, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if !p.isRunning {
		log.Println("PeriodicTaskService is not running.")
		return false, nil
	}

	log.Println("Stopping PeriodicTaskService...")
	p.stopChannel <- true
	return true, nil
}

func (p *PeriodicTaskService) processMessages() {
	log.Println("Processing messages...")
	if len(p.InstanceMessage.GetMessages()) < 2 {
		log.Println("Not enough messages to process.")
		//can call FetchUnsentMessages after mongo repo dependency added
		return
	}

	waitingMessage := p.InstanceMessage.GetMessages()
	toSend := waitingMessage[:2]
	for _, message := range toSend {
		payload := model.WebHookPayload{
			To:      message.RecipientPhone,
			Content: message.MessageContent,
		}

		response, err := p.WebhookService.SendPost(payload)
		if err != nil {
			log.Printf("Failed to send message to %s: %v", message.RecipientPhone, err)
			continue
		}

		if err := p.TempRepository.Set(response.MessageID, time.Now().Format(time.RFC3339), 10*time.Minute); err != nil {
			log.Printf("Failed to store messageId %s in TempRepository: %v", response.MessageID, err)
		} else {
			log.Printf("Message sent successfully: ID=%s, Time=%s", response.MessageID, time.Now().Format(time.RFC3339))
		}
	}

	p.InstanceMessage.SetMessages(waitingMessage[2:])
}
