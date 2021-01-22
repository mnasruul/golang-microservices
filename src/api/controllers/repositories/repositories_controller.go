package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/mnasruul/golang-microservices/src/api/domain/repositories"
	"github.com/mnasruul/golang-microservices/src/api/services"
	"github.com/mnasruul/golang-microservices/src/api/utils/errors"
)

func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	result, err := services.RepositoryService.CreateRepo(request)
}
