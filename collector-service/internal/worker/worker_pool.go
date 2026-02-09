package worker

import (
	"context"
	"log"
	"sync"

	"github.com/Chintukr2004/collector/internal/checker"
	"github.com/Chintukr2004/collector/internal/models"
	"github.com/Chintukr2004/collector/internal/repository"
)

func StartWorkerPool(
	services []models.Service,
	workerCount int,
	healthRepo *repository.HealthRepository,
	incidentRepo *repository.IncidentRepository,
) {

	var wg sync.WaitGroup

	jobs := make(chan models.Service)

	// Start workers
	for i := 0; i < workerCount; i++ {

		wg.Add(1)

		go func(workerID int) {
			defer wg.Done()

			for service := range jobs {

				result := checker.CheckHealth(service)
				//
				log.Println("Saving result for service ID:", result.ServiceID)

				err := healthRepo.SaveResult(context.Background(), result)
				if err != nil {
					log.Println("Failed to save health check:", err)
				}
				active, err := incidentRepo.HasActiveIncident(context.Background(), result.ServiceID)
				if err != nil {
					log.Println("Incident check failed:", err)
					continue
				}

				if result.Status == "DOWN" && !active {

					log.Println("🚨 INCIDENT STARTED for", result.ServiceName)

					err := incidentRepo.CreateIncident(context.Background(), result.ServiceID)
					if err != nil {
						log.Println("Failed to create incident:", err)
					}

				}

				if result.Status == "UP" && active {

					log.Println("✅ INCIDENT RESOLVED for", result.ServiceName)

					err := incidentRepo.ResolveIncident(context.Background(), result.ServiceID)
					if err != nil {
						log.Println("Failed to resolve incident:", err)
					}
				}

				log.Printf(
					"[Worker %d] %s is %s (latency: %v)",
					workerID,
					result.ServiceName,
					result.Status,
					result.Latency,
				)
			}

		}(i + 1)
	}

	// Send jobs
	for _, service := range services {
		jobs <- service
	}

	close(jobs)

	wg.Wait()
}
