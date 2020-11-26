/*
处理传感器传入的数据
*/

package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"net/http"

	"github.com/kataras/iris/v12"
	"ri-co.cn/m2m/sqls"
	"ri-co.cn/m2m/structs"
)

//SendTemp 接收传感器发来的信息
func SendTemp(con iris.Context) {
	var reqData structs.SensorData
	err := con.ReadJSON(&reqData)
	if err != nil {
		fmt.Println("温度传感器，传入数据出错", err.Error())
		SendMqttString("温度传感器接口，解析json失败\n" + err.Error())
		con.StatusCode(iris.StatusBadRequest)
		return
	}
	//当前时间
	timenow := time.Now().Format("2006-01-02 15:04:05")
	data := map[string]string{
		"time":    timenow,
		"RecData": reqData.M2m.Con,
	}
	SendMqtt(data)
	con.JSON(data)
}

//SendLight 接收光照传感器 int
func SendLight(con iris.Context) {
	var reqData structs.SensorData
	err := con.ReadJSON(&reqData)
	if err != nil {
		fmt.Println("光照传感器，传入数据出错", err.Error())
		SendMqttString("光照传感器接口，解析json失败\n" + err.Error())
		con.StatusCode(iris.StatusBadRequest)
		return
	}
	//当前时间
	timenow := time.Now().Format("2006-01-02 15:04:05")
	data := map[string]string{
		"time":    timenow,
		"RecData": reqData.M2m.Con,
	}
	SendMqtt(data)
	con.JSON(data)
}

//SendHumi 接收湿度数据
func SendHumi(con iris.Context) {
	var reqData structs.SensorData
	err := con.ReadJSON(&reqData)
	if err != nil {
		fmt.Println("湿度传感器，传入数据出错", err.Error())
		SendMqttString("湿度传感器接口，解析json失败\n" + err.Error())
		con.StatusCode(iris.StatusBadRequest)
		return
	}
	//当前时间
	timenow := time.Now().Format("2006-01-02 15:04:05")
	data := map[string]string{
		"time":    timenow,
		"RecData": reqData.M2m.Con,
	}
	SendMqtt(data)
	con.JSON(data)
}

//SendBody 接收人体数据
func SendBody(con iris.Context) {
	var reqData structs.SensorData
	err := con.ReadJSON(&reqData)
	if err != nil {
		fmt.Println("人体传感器，传入数据出错", err.Error())
		SendMqttString("人体传感器接口，解析json失败\n" + err.Error())
		con.StatusCode(iris.StatusBadRequest)
		return
	}
	//当前时间
	timenow := time.Now().Format("2006-01-02 15:04:05")
	data := map[string]string{
		"time":    timenow,
		"RecData": reqData.M2m.Con,
	}
	SendMqtt(data)
	con.JSON(data)
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

func ToWifi() {

	//要发送的json数据
	var sourceData structs.WifiPostData
	sourceData.M2m.Con = "ON"
	_, err := json.Marshal(sourceData)
	if err != nil {
		fmt.Println("发送wifi数据，json打包出错", err.Error())
		return
	}

	requestBody := new(bytes.Buffer)
	posturl := "http://v9v46x6k.shenzhuo.vip:10810"
	req, err := http.NewRequest("POST", posturl, requestBody)
	if err != nil {
		fmt.Println("初始化post出错", err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")
	post := &http.Client{}
	res, err := post.Do(req)
	if err != nil {
		fmt.Println("发送数据出错", err.Error())
		panic(err)
	}
	body, _ := ioutil.ReadAll(res.Body)

	SendMqttString(string(body))

	defer res.Body.Close()
}
