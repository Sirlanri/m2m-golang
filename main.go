package main

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, //允许通过的主机名称
		AllowCredentials: true,
	})

	sensor := app.Party("/sensor", crs).AllowMethods(iris.MethodOptions)

	app.Run(iris.Addr(":8090"))

	return
}
