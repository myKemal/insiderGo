package initialize

import (
	"github.com/myKemal/insiderGo/app/daos"
	_ "github.com/myKemal/insiderGo/app/daos"
	"github.com/myKemal/insiderGo/app/repository"
	"github.com/myKemal/insiderGo/app/services"
	"log"
)

func Start(mongoRepo repository.MongoRepository, tempRepo repository.TempMemoryRepository, messageService *services.MessageService) error {
	// Fetch unsent messages
	messageDAOList, err := mongoRepo.FetchUnsentMessages()
	if err != nil {
		log.Printf("Failed to fetch unsent messages: %v", err)
		return err
	}

	messageService.SetMessages(daos.ConvertDAOsToDTOs(messageDAOList))

	return nil
}
