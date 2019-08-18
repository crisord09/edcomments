package routes

import (
	"github.com/crislord09/edcomments/controllers"
	"github.com/gorilla/mux"
)

//SetLoginRouter estamos creando un router para login
func SetLoginRouter(router *mux.Router) {
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
}
