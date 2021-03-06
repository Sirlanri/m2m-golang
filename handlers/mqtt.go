package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	uuid "github.com/satori/go.uuid"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("主题: %s\n", msg.Topic())
	fmt.Printf("消息内容: %s\n", msg.Payload())
}

var c mqtt.Client

//Createid 生成唯一名称
func Createid() string {
	// 创建 UUID v4
	u1 := uuid.Must(uuid.NewV4(), nil)
	id := u1.String()
	return id[:9]
}

func init() {
	mqtt.ERROR = log.New(os.Stdout, "", 0)

	opts := mqtt.NewClientOptions().AddBroker("tcp://mqtt.ri-co.cn:1883").SetClientID("emqx_golang_" + Createid())

	opts.SetKeepAlive(600 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	c = mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

//SendMqtt 通过mqtt发送消息
func SendMqtt(payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("json打包出错", err.Error())
	}
	go func() {
		token := c.Publish("golang", 0, false, data)
		err = token.Error()
		if err != nil {
			fmt.Println("mqtt出错", err.Error())
		}
		token.Wait()
	}()

}

//SendMqttString 通过mqtt发送消息，文本格式
func SendMqttString(payload string) {
	go func() {
		token := c.Publish("golang", 0, false, payload)
		err := token.Error()
		if err != nil {
			fmt.Println("mqtt出错", err.Error())
		}
		token.Wait()
	}()

}
