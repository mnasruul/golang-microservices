package services

import (
	"github.com/mnasruul/golang-microservices/mvc/domain"
	"github.com/mnasruul/golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
