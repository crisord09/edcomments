package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/crislord09/edcomments/commons"
	"github.com/crislord09/edcomments/configuration"
	"github.com/crislord09/edcomments/models"
)

// VoteRegister controlador para registrar un voto
func VoteRegister(w http.ResponseWriter, r *http.Request) {
	vote := models.Vote{}
	user := models.User{}
	currentVote := models.Vote{}
	m := models.Message{}

	user, _ = r.Context().Value("user").(models.User)
	err := json.NewDecoder(r.Body).Decode(&vote)
	if err != nil {
		// StatusBadRequest el error fue de la peticion se envio mal la estructura del json
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error al leer el voto a registrar: %s", err)
		commons.DisplayMessage(w, m)
		return
	}
	vote.UserID = user.ID

	db := configuration.GetConnection()
	defer db.Close()

	db.Where("comment_id = ? AND user_id = ?", vote.CommentID, vote.UserID).First(&currentVote)
	// Si no existe
	if currentVote.ID == 0 {
		db.Create(&vote)
		err := updateComments(vote.CommentID, vote.Value, false)
		if err != nil {
			m.Message = err.Error()
			m.Code = http.StatusBadRequest
			commons.DisplayMessage(w, m)
			return
		}
		m.Message = "Voto registrado"
		m.Code = http.StatusCreated
		commons.DisplayMessage(w, m)
		return
	} else if currentVote.Value != vote.Value {
		currentVote.Value = vote.Value
		db.Save(&currentVote)
		err := updateComments(vote.CommentID, vote.Value, true)
		if err != nil {
			m.Message = err.Error()
			m.Code = http.StatusBadRequest
			commons.DisplayMessage(w, m)
			return
		}
		m.Message = "Voto Actualizado"
		m.Code = http.StatusOK
		commons.DisplayMessage(w, m)
	}
	m.Message = "Este voto ya está registrado"
	m.Code = http.StatusBadRequest
	commons.DisplayMessage(w, m)
}

// updateComments Actualiza la cantidad de votos en la talbla comentarios
// IsUpdate indica si es un voto para actualziar
func updateComments(commentID uint, vote bool, isUpdate bool) (err error) {
	comment := models.Comment{}

	db := configuration.GetConnection()
	defer db.Close()

	rows := db.First(&comment, commentID).RowsAffected

	if rows > 0 {
		if vote {
			comment.Votes++
			if isUpdate {
				comment.Votes++
			}
		} else {
			comment.Votes--
			if isUpdate {
				comment.Votes--
			}
		}
		db.Save(&comment)
	} else {
		err = errors.New("No se encontro un registro de comentario para asignarle el voto")
	}
	return
}
