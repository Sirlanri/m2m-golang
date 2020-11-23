/*
处理传感器传入的数据
*/

package handlers

import (
	"fmt"

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
	response.M2mcin.Con = "Rico: " + req.M2mcin.Con
	_, err = con.JSON(response)
	if err != nil {
		fmt.Println("返回数据出错", err.Error())
	}

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
