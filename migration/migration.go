//Este paquete nos va a servir para poder crear las tablas de nuestros modelos en Nuestra BD
//Este paquete se va a ekecutar solo una ves
package migration

//va a necesitar 2 packetes
import (
	"github.com/crislord09/edcomments/configuration"
	"github.com/crislord09/edcomments/models"
)

func Migrate() {
	db := configuration.GetConnection()
	defer db.Close()

	db.CreateTable(&models.User{})
	db.CreateTable(&models.Comment{})
	db.CreateTable(&models.Vote{})
	db.Model(&models.Vote{}).AddUniqueIndex("comment_id_user_id_unique", "comment_id", "user_id")

}
