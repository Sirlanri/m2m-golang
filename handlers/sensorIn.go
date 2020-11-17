/*
处理传感器传入的数据
*/

package handlers

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"ri-co.cn/m2m/sqls"
)

/*BodyRes -handler 人体传感器
传入的数据为0或1，表示是否有人 api- params {status:0}
*/
func BodyRes(con iris.Context) {
	resflag, err := con.URLParamInt("status")
	if err != nil {
		fmt.Println("人体传感器 传入数据有误", err.Error())
		con.StatusCode(201)
		return
	}
	result := sqls.BodyRes(resflag)
	if !result {
		con.StatusCode(210)
		con.WriteString("写入数据库失败")
		return
	}

}
