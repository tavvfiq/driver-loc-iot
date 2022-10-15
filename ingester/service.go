package locingester

import (
	"context"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/tavvfiq/driver-loc-iot/driver"
	mymqtt "github.com/tavvfiq/driver-loc-iot/pkg/mqtt"
)

type service struct {
	id     string
	client mymqtt.MyMqtt
}

var clientId = "driver-iot-ingester"

func NewIngester(id string) *service {
	myId := fmt.Sprintf("%s-%s", clientId, id)
	client := mymqtt.NewClient(myId, driver.BrokerUrl, driver.Qos)
	return &service{
		id:     id,
		client: client,
	}
}

func (s *service) Run(ctx context.Context) {
	err := s.client.Connect()
	if err != nil {
		panic(err)
	}
	err = s.client.Subscribe(driver.Topic, handler)
	if err != nil {
		panic(err)
	}
	<-ctx.Done()
	err = s.client.Unsubscribe(driver.Topic)
	if err != nil {
		fmt.Printf("ingester: %s failed to unsubscribe to topic: %s with error: %v\n", s.id, driver.Topic, err)
	}
	fmt.Println("driver-loc-iot ingesters shutdown gracefully")
}

func handler(c mqtt.Client, m mqtt.Message) {
	fmt.Println(string(m.Payload()))
}
