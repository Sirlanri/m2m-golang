/*
处理传感器传入的数据
*/

package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
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
	//温度写入data
	temp, _ := strconv.ParseFloat(reqData.M2m.Con, 32)
	SetTemp(float32(temp))
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
	//转换为float
	num, _ := strconv.ParseFloat(reqData.M2m.Con, 32)
	SetLight(float32(num))
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
	//写入data
	humi, _ := strconv.ParseFloat(reqData.M2m.Con, 32)
	SetHumi(float32(humi))
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
	num, _ := strconv.Atoi(reqData.M2m.Con)
	SetBody(num)
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

//Lighton 使用get手动操作测试开关灯
func Lighton(_ iris.Context) {
	LightToWifi("ON")
}

//Lightoff 使用get手动操作测试开关灯
func Lightoff(_ iris.Context) {
	LightToWifi("OFF")
}

//Buzzon 开启蜂鸣器
func Buzzon(_ iris.Context) {
	LightToWifi("xiang")
}

//LightToWifi 发送给wif模块的指令，ON||OFF
func LightToWifi(ins string) {
	//要发送的json数据
	var sourceData structs.WifiPostData
	sourceData.M2m.Con = ins //开灯
	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(sourceData)
	posturl := "http://v9v46x6k.shenzhuo.vip:10810"
	req, err := http.NewRequest("POST", posturl, requestBody)
	if err != nil {
		fmt.Println("初始化post出错", err.Error())
		return
	}

	req.Header.Set("content-type", "application/json")
	req.Close = true
	post := &http.Client{}
	res, _ := post.Do(req)

	body, _ := ioutil.ReadAll(res.Body)

	SendMqttString(string(body))

	defer res.Body.Close()
}

//以下为前端API

//GetTemp 获取当前温度
func GetTemp(con iris.Context) {
	temp := fmt.Sprintf("%.2f", Temp)
	data := map[string]string{
		"temp": temp,
	}
	con.JSON(data)
}

//GetHumi 获取当前湿度
func GetHumi(con iris.Context) {
	humi := fmt.Sprintf("%.2f", Humi)
	data := map[string]string{
		"humi": humi,
	}
	con.JSON(data)
}

//GetLight 获取当前光强
func GetLight(con iris.Context) {
	light := Light
	data := map[string]float32{
		"light": light,
	}
	con.JSON(data)
}

//GetBody 获取当前湿度
func GetBody(con iris.Context) {
	body := Body
	data := map[string]int{
		"body": body,
	}
	con.JSON(data)
}
