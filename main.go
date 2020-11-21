package main

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"ri-co.cn/m2m/handlers"
)

func main() {
	app := iris.New()
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, //允许通过的主机名称
		AllowCredentials: true,
	})
	m2m := app.Party("/m2m", crs).AllowMethods(iris.MethodOptions)

	//前端
	front := m2m.Party("/front")

	front.Get("/getTemp", handlers.GetTempTest)
	front.Get("/getHumi", handlers.GetHumiTest)
	front.Get("/getLight", handlers.GetLightTest)

	app.Run(iris.Addr(":8090"))

	return
}
