package main

import (
	"flag"
	"log"

	"github.com/crislord09/edcomments/migration"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la BD")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Comenzo la migración")
		migration.Migrate()
		log.Println("Finalizo la migración")
	}
}
