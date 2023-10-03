package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/RpThiagoluiz/go-simple-crud/models"
	"github.com/go-chi/chi/v5"
)

func Delete(w http.ResponseWriter, r*http.Request)  {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil{
		log.Printf("Error in decode JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
		}
	



	rows, err:= models.Delete(int64(id))

	if err != nil {
		log.Printf("Error when try delete todo: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}


	if(rows > 1){
		log.Printf("Error when try delete todo %d this update more than one", rows)
	} 

	resp := map[string]any{
		"Error": false,
		"Message":"Success in delete todo",
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)
}




