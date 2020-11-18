package handlers

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

//MqttRes 处理接收到的mqtt数据
func MqttRes() {
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker("tcp://39.96.49.63:1883").SetClientID("emqx_test_client")

	opts.SetKeepAlive(600 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	var wg sync.WaitGroup
	wg.Add(1)

	//接收端
	go func() {
		if token := c.Subscribe("testtopic/#", 0, nil); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}
	}()
	//发送端
	go func() {
		token := c.Publish("testtopic/gosend", 0, false, "嗯哼~从go发出的")
		token.Wait()
	}()
	wg.Wait()
}
