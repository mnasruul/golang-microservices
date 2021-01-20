package domain

import (
	"fmt"
	"github.com/mnasruul/golang-microservices/mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		123: &User{Id: 123, FirstName: "Nasrul", LastName: "Muhammad", Email: "mnasruul@gmail.com"},
		124: &User{Id: 124, FirstName: "Nasrul2", LastName: "Muhammad2", Email: "mnasruul@gmail.com2"},
		125: &User{Id: 125, FirstName: "Nasrul2", LastName: "Muhammad2", Email: "mnasruul@gmail.com2"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
