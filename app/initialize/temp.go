package initialize

import (
	"github.com/myKemal/insiderGo/app/services"
	"os"

	"github.com/myKemal/insiderGo/app/repository"
)

func Temp() (repository.TempMemoryRepository, error) {
	storageType := os.Getenv("TEMP_STORAGE") //  "REDIS" veya "INMEMORY"
	switch storageType {
	case "REDIS":
		redisHost := os.Getenv("REDIS_HOST")
		redisPort := os.Getenv("REDIS_PORT")
		client := services.NewRedisClient(redisHost + ":" + redisPort)
		return repository.NewTempMemory(client), nil
	case "INMEMORY":
		client := services.NewInMemoryClient()
		return repository.NewTempMemory(client), nil
	default:
		client := services.NewInMemoryClient()
		return repository.NewTempMemory(client), nil
	}
}
