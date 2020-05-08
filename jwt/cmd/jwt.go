package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"jwt_go/server/jwt/routers"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	routers.SetRouter(app, "/api/v1")

	app.Run(iris.Addr(":4011"), iris.WithoutPathCorrection)
}
