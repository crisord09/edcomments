package models

import "github.com/jinzhu/gorm"

/*====  Esta estructura nos sirve para tener listo la BD 👈😉👇  ====*/
type User struct {
	//Gorm crea 3 campos
	gorm.Model

	//Estamos creando una API por eso las primeras letras en Json deben de ir en minusculas
	//Al ORM debemos de indicarle que este campo debe de ser unico y no nulo
	Username string `json:"username" gorm:"not null;unique"`
	Email    string `json:"email" gorm:"not null;unique"` //unico para que no haya 2 usuarios con el mismo email
	Fullname string `json:"fullname" gorm:"not null"`
	//No es necesario el password a la persona por que no lo necesita(omitempty)
	Password string `json:"password,omitempty" gorm:"not null;type:varchar(256)"`
	//Con (-) el ORM entiende que este campo no lo va a crear
	ConfirmPassword string    `json:"confirmPassword,omitempty" gorm:"-"`
	Picture         string    `json:"picture"`
	Comments        []Comment `json:"comments,omitempty"`
}
