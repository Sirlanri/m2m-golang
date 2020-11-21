/*
处理传感器传入的数据
*/

package handlers

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"ri-co.cn/m2m/structs"
)

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
