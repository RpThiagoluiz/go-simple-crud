package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/RpThiagoluiz/go-simple-crud/models"
	"github.com/go-chi/chi/v5"
)

func Get(w http.ResponseWriter, r*http.Request){
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil{
		log.Printf("Error in decode JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
		}
	
	todo, err:= models.Get(int64(id))

	if err != nil {
		log.Printf("Error when try get todo: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(todo)



}