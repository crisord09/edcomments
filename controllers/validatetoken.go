package controllers

import (
	"context"
	"net/http"

	"github.com/crislord09/edcomments/commons"
	"github.com/crislord09/edcomments/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

// ValidateToken validar el token del cliente
func ValidateToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	//Con este va a retornar respuestas al cliente
	var m models.Message
	//el paquete request nos ayuda para traer el Token -> esto viene del jwt
	token, err := request.ParseFromRequestWithClaims(
		r,                       //el parametro del request
		request.OAuth2Extractor, //esto viene del paquete
		&models.Claim{},         //a la estructura de nuestro token
		//funcion anonima
		func(t *jwt.Token) (interface{}, error) {
			return commons.PublicKey, nil
		},
	)
	if err != nil {
		m.Code = http.StatusUnauthorized
		switch err.(type) {
		case *jwt.ValidationError:
			vError := err.(*jwt.ValidationError)
			switch vError.Errors {
			case jwt.ValidationErrorExpired:
				m.Message = "Su token ha expirado"
				commons.DisplayMessage(w, m)
				return
			case jwt.ValidationErrorSignatureInvalid:
				m.Message = "La firma del token no coincide"
				commons.DisplayMessage(w, m)
				return
			default:
				m.Message = "Su token no es válido"
				commons.DisplayMessage(w, m)
				return
			}
		}
	}

	if token.Valid {
		ctx := context.WithValue(r.Context(), "user", token.Claims.(*models.Claim).User)
		next(w, r.WithContext(ctx))
	} else {
		m.Code = http.StatusUnauthorized
		m.Message = "Su token no es válido"
		commons.DisplayMessage(w, m)
	}

}
