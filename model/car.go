package model

import "encoding/json"

type Car struct {
	OwnerName          string                `json:"ownerName"`
	SerialNumber       uint64                `json:"serialNumber"`
	ModelYear          uint64                `json:"modelYear"`
	Code               string                `json:"code"`
	VehicleCode        string                `json:"vehicleCode"`
	Engine             *CarEngine            `json:"engine"`
	FuelFigures        *CarFuelFigures       `json:"fuelFigures"`
	PerformanceFigures *CarPerfomanceFigures `json:"performanceFigures"`
	Manufacturer       string                `json:"manufacturer"`
	Model              string                `json:"model"`
	ActivationCode     string                `json:"activationCode"`
}

type CarEngine struct {
	capacity         uint16 `json:"capacity"`
	numCylinders     uint8  `json:"numCylinders"`
	maxRpm           uint16 `json:"maxRpm"`
	manufacturerCode string `json:"manufacturerCode"`
}

type CarFuelFigures struct {
	speed            uint16  `json:"speed"`
	mpg              float64 `json:"mpg"`
	usageDescription string  `json:"usageDescription"`
}

type CarPerfomanceFigures struct {
	octaneRating uint16                            `json:"octaneRating"`
	acceleration *CarPerfomanceFiguresAcceleration `json:"acceleration"`
}

type CarPerfomanceFiguresAcceleration struct {
	mph     uint16  `json:"mph"`
	seconds float64 `json:"seconds"`
}

func (c *Car) JSON() []byte {
	j, _ := json.Marshal(c)
	return j
}

type CarStorage interface {
	FindBySerialNumber(serialNumber uint64) (*Car, error)
}

type CarInteractor interface {
	FindBySerialNumber(serialNumber uint64) (*Car, error)
}
