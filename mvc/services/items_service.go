package services

import (
	"github.com/mnasruul/golang-microservices/mvc/domain"
	"github.com/mnasruul/golang-microservices/mvc/utils"
	"net/http"
)

type itemsService struct{}

var (
	ItemsService itemsService
)

func (s *itemsService) GetItem(itemId string) (*domain.User, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
