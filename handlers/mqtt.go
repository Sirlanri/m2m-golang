package handlers

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var client mqtt.Client

var deal mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func init() {
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker("tcp://mqtt.ri-co.cn:1883").SetClientID("golang_back")

	opts.SetKeepAlive(600 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(deal)
	opts.SetPingTimeout(5 * time.Second)

	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	//调用接收函数
	MqttRes()
}

//MqttRes 处理接收到的mqtt数据
func MqttRes() {
	var wg sync.WaitGroup
	wg.Add(1)

	//接收端
	go func() {
		if token := client.Subscribe("testtopic/#", 0, nil); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}
	}()

	wg.Wait()
}

//MqttSend 发送mqtt数据
func MqttSend(topic, msg string) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		token := client.Publish(topic, 0, false, msg)
		token.Wait()
		wg.Done()
	}()
	wg.Wait()
}

//MqttSendDefault 发送mqtt数据，使用默认topic：gosend
func MqttSendDefault(msg string) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		token := client.Publish("gosend", 0, false, msg)
		token.Wait()
		wg.Done()
	}()
	wg.Wait()
}
