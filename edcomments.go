package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/crislord09/edcomments/commons"
	"github.com/crislord09/edcomments/migration"
	"github.com/crislord09/edcomments/routes"
	"github.com/urfave/negroni"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la BD")
	flag.IntVar(&commons.Port, "port", 8080, "Puerto para el servidor web")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Comenzo la migración")
		migration.Migrate()
		log.Println("Finalizo la migración")
	}

	//Inicia las rutas
	router := routes.InitRoutes()

	// Inicia los middlewares
	n := negroni.Classic()
	n.UseHandler(router)

	//Inicia el servidor
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", commons.Port),
		Handler: n,
	}

	log.Printf("Iniciado el servidor en http://localhost:%d", commons.Port)
	log.Println(server.ListenAndServe())
	log.Println("Finalizo la ejecución del programa")
}
