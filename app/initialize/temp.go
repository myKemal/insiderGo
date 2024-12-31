package initialize

import (
	"github.com/myKemal/mongoApi/app/services"
	"os"

	"github.com/myKemal/mongoApi/app/repository"
)

func Temp() (repository.TempMemoryRepo, error) {
	storageType := os.Getenv("TEMP_STORAGE") // Set to "REDIS" or "MEMORY"
	switch storageType {
	case "REDIS":
		redisHost := os.Getenv("REDIS_HOST")
		redisPort := os.Getenv("REDIS_PORT")
		return services.NewRedisClient(redisHost + ":" + redisPort), nil
	default:
		return services.NewInMemoryClient(), nil
	}
}
