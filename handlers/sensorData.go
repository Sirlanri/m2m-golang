package handlers

import (
	"fmt"
	"strconv"

	"ri-co.cn/m2m/sqls"
)

var (
	Humi  float32
	Temp  float32
	Body  int
	Light float32

	humiindex  int
	tempindex  int
	bodyindex  int
	lightindex int

	bodyTimes int

	index = 2
)

//SetTemp 更改温度/写入数据库
func SetTemp(num float32) {
	Temp = num
	tempindex++
	fmt.Println("温度数据已更新", num)
	SendMqttString("温度数据已更新 " + fmt.Sprintf("%f", num))

	//测试环境，暂定index次写入数据库
	if tempindex == index {
		sqls.TempRes(num)
		tempindex = 0
		fmt.Println("温度已写入数据库", num)
		SendMqttString("温度已写入数据库 " + fmt.Sprintf("%f", num))
	}
}

//SetHumi 更改温度/写入数据库
func SetHumi(num float32) {
	Humi = num
	humiindex++
	fmt.Println("湿度数据已更新", num)
	SendMqttString("湿度数据已更新 " + fmt.Sprintf("%f", num))

	//测试环境，暂定index次写入数据库
	if humiindex == index {
		sqls.HumiRes(num)
		humiindex = 0
		fmt.Println("湿度已写入数据库", num)
		SendMqttString("湿度已写入数据库 " + fmt.Sprintf("%f", num))
	}
}

//SetLight 更改光照/写入数据库
func SetLight(num float32) {
	Light = num
	lightindex++
	fmt.Println("光照数据已更新", num)
	SendMqttString("光照数据已更新 " + fmt.Sprintf("%f", num))

	//测试环境，暂定index次写入数据库
	if lightindex == index {
		sqls.LightRes(num)
		lightindex = 0
		fmt.Println("光照已写入数据库", num)
		SendMqttString("光照已写入数据库 " + fmt.Sprintf("%f", num))
	}

	//亮度<300 开灯
	if num <= 100 {
		LightToWifi("ON")
	}
	if num > 700 {
		LightToWifi("OFF")
	}
}

//SetBody 更改人体/写入数据库
func SetBody(num int) {
	Body = num
	bodyindex++
	fmt.Println("人体数据已更新", num)
	SendMqttString("人体数据已更新 " + strconv.Itoa(num))

	//测试环境，暂定index次写入数据库
	if bodyindex == index {
		sqls.BodyRes(num)
		bodyindex = 0
		fmt.Println("人体已写入数据库", num)
		SendMqttString("人体已写入数据库 " + strconv.Itoa(num))
	}

	//达到一定次数，蜂鸣器响
	if num == 1 {
		bodyTimes++
	} else {
		bodyTimes = 0
	}

	if bodyTimes == 3 {
		LightToWifi("xiang")
		bodyTimes = 0
	}

}
