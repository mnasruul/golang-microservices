package services

import (
	"github.com/mnasruul/golang-microservices/mvc/domain"
)

func GetUser(userId int64) (domain.User, error) {
	return domain.GetUser(userId)
}
