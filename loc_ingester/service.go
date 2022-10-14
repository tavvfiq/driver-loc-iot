package locingester

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/tavvfiq/driver-loc-iot/driver"
	mymqtt "github.com/tavvfiq/driver-loc-iot/pkg/mqtt"
)

type service struct {
	client mymqtt.MyMqtt
}

var clientId = "driver-iot-ingester"

func NewIngester() *service {
	client := mymqtt.NewClient(clientId, driver.BrokerUrl, 1)
	return &service{
		client: client,
	}
}

func (s *service) Run() {
	err := s.client.Connect()
	if err != nil {
		panic(err)
	}
	err = s.client.Subscribe(driver.Topic, handler)
	if err != nil {
		panic(err)
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
	err = s.client.Unsubscribe(driver.Topic)
	if err != nil {
		fmt.Printf("failed to unsubscribe to topic: %s with error: %v\n", driver.Topic, err)
	}
	fmt.Println("driver-loc-iot ingesters shutdown gracefully")
}

func handler(c mqtt.Client, m mqtt.Message) {
	fmt.Printf("got from topic: %s, payload: %v\n", driver.Topic, string(m.Payload()))
}
