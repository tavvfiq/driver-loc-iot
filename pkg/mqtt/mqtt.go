package mymqtt

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MyMqtt interface {
	SetConnectionHandler(handler mqtt.OnConnectHandler)
	Connect() error
	Disconnect()
	Publish(topic string, payload []byte) error
	Subscribe(topic string, handler mqtt.MessageHandler) error
	Unsubscribe(topic string) error
}

type service struct {
	broker           string
	clientId         string
	client           mqtt.Client
	qos              byte
	connectedHandler mqtt.OnConnectHandler
	connLostHandler  mqtt.ConnectionLostHandler
}

func NewClient(clientId string, broker string, qos byte) MyMqtt {
	return &service{
		broker:   broker,
		clientId: clientId,
		qos:      qos,
	}
}

func (s *service) SetConnectionHandler(handler mqtt.OnConnectHandler) {
	s.connectedHandler = handler
}

func (s *service) Connect() error {
	options := mqtt.NewClientOptions()
	// broker IP and port
	options.AddBroker(s.broker)
	options.SetClientID(s.clientId)
	options.OnConnect = s.connectedHandler
	options.OnConnectionLost = s.connLostHandler
	client := mqtt.NewClient(options)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	s.client = client
	fmt.Printf("connected to broker: %s with clientId: %s\n", s.broker, s.clientId)
	return nil
}

func (s *service) Publish(topic string, payload []byte) error {
	token := s.client.Publish(topic, s.qos, false, payload)
	if token.Wait(); token.Error() != nil {
		return token.Error()
	}
	fmt.Printf("published to topic: %s\n", topic)
	return nil
}

func (s *service) Subscribe(topic string, handler mqtt.MessageHandler) error {
	token := s.client.Subscribe(topic, s.qos, handler)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	fmt.Printf("subscribed to topic: %s with qos: %d\n", topic, s.qos)
	return nil
}

func (s *service) Disconnect() {
	s.client.Disconnect(0)
	fmt.Printf("client with id: %s disconnected\n", s.clientId)
}

func (s *service) Unsubscribe(topic string) error {
	token := s.client.Unsubscribe(topic)
	if token.Wait(); token.Error() != nil {
		return token.Error()
	}
	fmt.Printf("client: %s unsubscribed from topic: %s\n", s.clientId, topic)
	return nil
}
