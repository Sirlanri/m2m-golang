package handlers

import (
	"math/rand"

	"github.com/kataras/iris/v12"
)

//GetTempTest 模拟接口 温度
func GetTempTest(con iris.Context) {
	temp := rand.Int31n(50)
	data := map[string]int32{
		"temp": temp,
	}
	con.JSON(data)
}

//GetHumiTest 模拟接口 湿度
func GetHumiTest(con iris.Context) {
	humi := rand.Int31n(100)
	data := map[string]int32{
		"humi": humi,
	}
	con.JSON(data)
}

//GetLightTest 模拟接口 光照
func GetLightTest(con iris.Context) {
	light := rand.Int31n(100)
	data := map[string]int32{
		"light": light,
	}
	con.JSON(data)
}

//GetVoiceTest 模拟接口 是否有声音
func GetVoiceTest(con iris.Context) {
	voice := rand.Intn(2)
	data := map[string]int{
		"voice": voice,
	}
	con.JSON(data)
}
