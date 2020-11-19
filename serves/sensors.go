package serves

import (
	"fmt"

	"ri-co.cn/m2m/sqls"
)

// BodyRes -serves 人体传感器
// 传入的数据为0或1，表示是否有人
func BodyRes(resflag int) {
	result := sqls.BodyRes(resflag)
	if !result {
		fmt.Println("写入数据库失败")
		return
	}
}

// TempRes -serves 温度传感器
// 传入温度值 float
func TempRes(temp float64) {
	if temp < 100.0 && temp > -20.0 {
		sqls.TempRes(temp)
	} else {
		fmt.Println("温度传感器，传入数据不合法")
	}
}

// HumiRes -serves 湿度传感器
// 传入湿度百分比
func HumiRes(humi float64) {
	if humi < 100 && humi > 0 {
		sqls.HumiRes(humi)
	} else {
		fmt.Println("湿度传感器，传入数据不合法")
	}
}
