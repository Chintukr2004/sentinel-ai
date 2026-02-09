package main

import (
	"context"
	"log"
	"time"

	"github.com/Chintukr2004/collector/internal/worker"

	"github.com/Chintukr2004/collector/internal/repository"
	"github.com/Chintukr2004/collector/pkg/db"
)

func main() {

	database := db.NewDB()
	repo := repository.NewServiceRepository(database)
	healthRepo := repository.NewHealthRepository(database)
	incidentRepo := repository.NewIncidentRepository(database)

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		log.Println("Starting health checks.....")

		services, err := repo.GetAllServices(context.Background())
		if err != nil {
			log.Fatal(err)
			continue
		}
		worker.StartWorkerPool(services, 3, healthRepo, incidentRepo)

		<-ticker.C
	}

}
