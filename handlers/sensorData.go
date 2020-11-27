package handlers

import (
	"fmt"

	"ri-co.cn/m2m/sqls"
)

var (
	Humi  float32
	Temp  float32
	Body  int
	Light int

	humiindex  int
	tempindex  int
	bodyindex  int
	lightindex int
)

//SetTemp 更改温度/写入数据库
func SetTemp(num float32) {
	Temp = num
	tempindex++
	fmt.Println("温度数据已更新", num)
	SendMqttString("温度数据已更新 " + fmt.Sprintf("%f", num))

	//测试环境，暂定2次写入数据库
	if tempindex == 1 {
		sqls.TempRes(num)
		tempindex = 0
		fmt.Println("温度已写入数据库", num)
		SendMqttString("温度已写入数据库 " + fmt.Sprintf("%f", num))
	}
}
