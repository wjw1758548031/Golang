package routers

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"jwt_go/server/jwt/controllers"
)

func SetRouter(router iris.Party, path string) iris.Party {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	})
	router = router.Party("/", crs).AllowMethods(iris.MethodOptions)
	r := router.Party(path)
	r.Post("/create_token", controllers.CreateToken)
	r.Post("/query_token", controllers.VerifyToken, controllers.QueryToken)
	return router

}
