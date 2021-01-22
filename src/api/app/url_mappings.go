package app

import (
	"github.com/mnasruul/golang-microservices/src/api/controllers/polo"
	"github.com/mnasruul/golang-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)
	router.GET("/repositories/:user_id", repositories.CreateRepo)
}
