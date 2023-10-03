package main

import (
	"fmt"
	"net/http"

	"github.com/RpThiagoluiz/go-simple-crud/configs"
	"github.com/RpThiagoluiz/go-simple-crud/handlers"
	"github.com/go-chi/chi/v5"
)

func main()  {
	err := configs.Load()

	if err!= nil{
		panic(err)
	}

	r := chi.NewRouter()
	r.Get("/", handlers.List)
	r.Post("/",handlers.Create)
	r.Get("/{id}", handlers.Get)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	

	http.ListenAndServe(fmt.Sprintf(":%s",configs.GetServerPort()), r)

}