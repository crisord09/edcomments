//Esto vamos a inicialiozaro desde el main de la app
package routes

import (
	"github.com/gorilla/mux"
)

//InitRoutes Inicia todas las rutas
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	//Como estamos dentro de nuestor packete no es necesario importarlo
	SetLoginRouter(router)
	SetUserRouter(router)
	SetCommentRouter(router)
	SetVoteRouter(router)
	SetPublicRouter(router)
	SetRealtimeRouter(router)

	return router

}
