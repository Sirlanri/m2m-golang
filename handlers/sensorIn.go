/*
处理传感器传入的数据
*/

package handlers

import (
	"fmt"
	"time"

	"github.com/kataras/iris/v12"
	"ri-co.cn/m2m/sqls"
	"ri-co.cn/m2m/structs"
)

//SendTemp 接收传感器发来的信息
func SendTemp(con iris.Context) {
	var req structs.ReqData
	err := con.ReadJSON(&req)
	if err != nil {
		fmt.Println("前端传入格式错误", err.Error())
		return
	}
	var response = req

	//写入当前时间
	timenow := time.Now().Format("2006-01-02 15:04:05")
	response.M2mcin.Lt = timenow
	response.M2mcin.Con = "Rico: " + req.M2mcin.Con
	_, err = con.JSON(response)
	if err != nil {
		fmt.Println("返回数据出错", err.Error())
	}

	//通过mqtt发送json
	SendMqtt(response)

}

//GetTimePer 数据统计页面，获取有无人的统计次数
func GetTimePer(con iris.Context) {
	have, no := sqls.GetTimePer()
	data := map[string]int{
		"have": have,
		"no":   no,
	}
	con.JSON(data)
}

//GetWeekTempHumi 获取一周温度湿度的平均值列表
func GetWeekTempHumi(con iris.Context) {
	data := sqls.GetWeekTempHumi()
	con.JSON(data)
}
