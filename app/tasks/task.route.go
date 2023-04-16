package tasks

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type taskServer struct{
	store *TaskStore
}

func NewTaskServer() *taskServer{
	store := New()
	return &taskServer{store: store}
}

func (server *taskServer) TaskHandler(w http.ResponseWriter, request *http.Request){
	if(request.URL.Path == "/tasks"){
		switch request.Method {
			case http.MethodPost: {
				decoder := json.NewDecoder(request.Body)
				decoder.DisallowUnknownFields()
				var body Task
				if err := decoder.Decode(&body); err != nil{
					http.Error(w,err.Error(), http.StatusBadRequest)
					return
				}
				fmt.Println(body);
				result := server.store.Create(body.Name)
				response, err := json.Marshal(result)
				if(err != nil){
					http.Error(w,err.Error(), http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(response)
			}
			case http.MethodGet: {
				result := server.store.List()
				response, err := json.Marshal(result)
				if(err != nil){
					http.Error(w,err.Error(), http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(response)
			}
		}

	}
}