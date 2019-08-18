package commons

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/crislord09/edcomments/models"
)

// DisplayMessage devuelve un mensaje al cliente
func DisplayMessage(w http.ResponseWriter, m models.Message) {
	j, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("Error al convertir el mensaje: %s", err)
	}

	//Vamos a escribir en el cabezado del ResponsWriter
	w.WriteHeader(m.Code)
	//Escribimos el mensaje en el cuerpo del responsewriter
	w.Write(j)

}
