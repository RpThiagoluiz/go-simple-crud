package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/RpThiagoluiz/go-simple-crud/models"
)

func List( w http.ResponseWriter, r*http.Request){
	todos, err := models.List()

	if err !=nil {
		log.Printf("Error in find all todos: %v", err)
	}


	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(todos)


}