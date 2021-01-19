package app

import (
	"net/http"

	"github.com/mnasruul/golang-microservices/mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser())
}
