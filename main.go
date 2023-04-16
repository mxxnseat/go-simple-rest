package main

import (
	"net/http"

	"tasks"
)

func main(){
	mux := http.NewServeMux()
	server := tasks.NewTaskServer()
	mux.HandleFunc("/tasks", server.TaskHandler)
	http.ListenAndServe("localhost:3000",mux)

}