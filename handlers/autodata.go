package handlers

import (
	"math/rand"

	"github.com/kataras/iris/v12"
)

func GetTempTest(con iris.Context) {
	temp := rand.Int31n(50)
	data := map[string]int32{
		"temp": temp,
	}
	con.JSON(data)
}

func GetHumiTest(con iris.Context) {
	humi := rand.Int31n(100)
	data := map[string]int32{
		"humi": humi,
	}
	con.JSON(data)
}

func GetLightTest(con iris.Context) {
	light := rand.Int31n(1000)
	data := map[string]int32{
		"light": light,
	}
	con.JSON(data)
}
