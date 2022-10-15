package driver

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/tavvfiq/driver-loc-iot/pkg/model"
	mymqtt "github.com/tavvfiq/driver-loc-iot/pkg/mqtt"
)

var Topic = "driver-loc-iot"
var clientId = "driver-iot"
var BrokerUrl = "tcp://localhost:1883"
var Qos = byte(2)

type service struct {
	driverId string
	client   mymqtt.MyMqtt
}

// simulate location sending to mqtt
func NewDriver(driverId string) *service {
	myId := fmt.Sprintf("%s-%s", clientId, driverId)
	client := mymqtt.NewClient(myId, BrokerUrl, Qos)
	return &service{client: client, driverId: driverId}
}

func (s *service) Run(ctx context.Context) {
	err := s.client.Connect()
	if err != nil {
		panic(err)
	}

	for {
		select {
		case <-ctx.Done():
			s.client.Disconnect()
			fmt.Printf("%s shutdown gracefully\n", s.driverId)
			return
		case <-time.After(10 * time.Second):
			loc := model.RawLocation{
				Latitude:  1.23124323,
				Longitude: 110.23123123,
				Timestamp: time.Now(),
			}
			payload := model.MQTTPayload{
				Id:          "1",
				OrderNumber: "RB-123",
				DriverId:    s.driverId,
				Location:    loc,
			}
			b, _ := json.Marshal(payload)
			s.client.Publish(Topic, b)
		}
	}
}
