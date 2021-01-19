package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hellow World"))
	})
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
