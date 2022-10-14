package driver

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tavvfiq/driver-loc-iot/pkg/model"
	mymqtt "github.com/tavvfiq/driver-loc-iot/pkg/mqtt"
)

var Topic = "driver-loc-iot"
var clientId = "driver-iot"
var BrokerUrl = "tcp://localhost:1883"

type service struct {
	client mymqtt.MyMqtt
}

// simulate location sending to mqtt
func NewDriver() *service {
	client := mymqtt.NewClient(clientId, BrokerUrl, 1)
	return &service{client: client}
}

func (s *service) Run() {
	err := s.client.Connect()
	if err != nil {
		panic(err)
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-sig:
			s.client.Disconnect()
			fmt.Println("driver-iot shutdown gracefully")
			return
		case <-time.After(10 * time.Second):
			data := model.RawLocation{
				Latitude:  1.23124323,
				Longitude: 110.23123123,
				Timestamp: time.Now(),
			}
			b, _ := json.Marshal(data)
			s.client.Publish(Topic, b)
		}
	}
}
