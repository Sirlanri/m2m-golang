package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {
	//mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:1882").SetClientID("emqx_test_client")

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
	//time.Sleep(time.Hour)
	// 发布消息
	//token := c.Publish("testtopic/1", 0, false, "Hello World")
	//token.Wait()

	//time.Sleep(60 * time.Second)

}

func mysend() {
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:1882").SetClientID("emqx_test_client")

	opts.SetKeepAlive(600 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(f)
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	// 订阅主题
	if token := c.Subscribe("testtopic/#", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// 发布消息
	token := c.Publish("testtopic/1", 0, false, "Send by rico from go")
	token.Wait()

	time.Sleep(60 * time.Second)
}

func MyWait() {
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:1882").SetClientID("emqx_test_client")

	opts.SetKeepAlive(600 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	// 订阅主题
	if token := c.Subscribe("testtopic/#", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	time.Sleep(60 * time.Second)
}
