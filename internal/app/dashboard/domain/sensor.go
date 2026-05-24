package domain

import "time"

type Sensor struct {
	SensorID     string
	Name         string
	Code         string
	Symbol       string
	Status       string
	MetricTypes  []MetricType
	Measurements []Measurement
}

type MetricType struct {
	MetricTypeID string
	SensorID     string
	Code         string
	Name         string
	Symbol       string
	MinValue     float64
	MaxValue     float64
}

type Measurement struct {
	SensorID   string
	PatientID  string
	CreatedAt  time.Time
	Components []MeasurementComponent
}

type MeasurementComponent struct {
	MetricTypeID string
	Code         string
	Name         string
	Value        float64
	Symbol       string
}
