package domain

import "time"

type Sensor struct {
	SensorID string
	Name     string
	Code     string
	Symbol   string
	Status   string
	Metrics  []SensorMetric
}

type SensorMetric struct {
	Value     float64
	Symbol    string
	CreatedAt time.Time
}
