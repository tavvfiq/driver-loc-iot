package model

import "time"

type MQTTPayload struct {
	Id          string      `json:"id"`
	DriverId    string      `json:"driver_id"`
	OrderNumber string      `json:"order_number"`
	Location    RawLocation `json:"location"`
}

type RawLocation struct {
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Timestamp time.Time `json:"time_stamp"`
}

type KafkaPayload struct {
	Id          string      `json:"id"`
	OrderNumber string      `json:"order_number"`
	DriverId    string      `json:"driver_id"`
	Location    RawLocation `json:"location"`
	Timestamp   time.Time   `json:"timestamp"`
}

type DriverLocation struct {
	Id            string      `json:"id"`
	OrderNumber   string      `json:"order_number"`
	DriverId      string      `json:"driver_id"`
	StartLocation RawLocation `json:"start_location"`
	LastLocation  RawLocation `json:"last_location"`
	EndLocation   RawLocation `json:"end_location"`
}
