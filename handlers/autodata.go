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
	light := rand.Int31n(1000)
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

//GetWeekData 模拟接口，3D柱状图
func GetWeekData(con iris.Context) {
	var data [][]int
	//单个数据由单个列表、列表里有3个int
	for x := 0; x < 7; x++ {
		for y := 0; y < 24; y++ {
			randz := rand.Intn(20)
			single := []int{x, y, randz}
			data = append(data, single)
		}
	}
	resData := map[string]([][]int){
		"list": data,
	}
	con.JSON(resData)

}
