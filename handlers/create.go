package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/RpThiagoluiz/go-simple-crud/models"
)

func Create(w http.ResponseWriter, r*http.Request) {
	var todo models.Todo


	err :=json.NewDecoder(r.Body).Decode(&todo)

	if err !=nil{
		log.Printf("Error in decode JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error": true,
			"Message": fmt.Sprintf("Error in when try insert data: %v", err),
		} 
	}else {
		resp = map[string]any{
			"Error": false,
			"Message": fmt.Sprintf("Success in insert data, ID:%d", id),
		}
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)
}