//TOKENS claim nos va a permitir Tener la información del usuario
//Es como para guardar los datos en la session
//El codigo de este bloque me sirve para saber si el usuario se autentifico o no
package models

import jwt "github.com/dgrijalva/jwt-go"

//Claim lo entenderemos como solicitud o reclamo
// Claim Token del usuario
type Claim struct {
	User `json:"user"`
	jwt.StandardClaims
}
