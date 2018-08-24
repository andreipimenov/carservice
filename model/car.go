package model

import "encoding/json"

type Car struct {
	OwnerName          string                `json:"ownerName" bson:"ownerName"`
	SerialNumber       uint64                `json:"serialNumber" bson:"serialNumber"`
	ModelYear          uint64                `json:"modelYear" bson:"modelYear"`
	Code               string                `json:"code" bson:"code"`
	VehicleCode        string                `json:"vehicleCode" bson:"vehicleCode"`
	Engine             *CarEngine            `json:"engine" bson:"engine"`
	FuelFigures        *CarFuelFigures       `json:"fuelFigures" bson:"fuelFigures"`
	PerformanceFigures *CarPerfomanceFigures `json:"performanceFigures" bson:"performanceFigures"`
	Manufacturer       string                `json:"manufacturer" bson:"manufacturer"`
	Model              string                `json:"model" bson:"model"`
	ActivationCode     string                `json:"activationCode" bson:"activationCode"`
}

type CarEngine struct {
	Capacity         uint16 `json:"capacity" bson:"capacity"`
	NumCylinders     uint8  `json:"numCylinders" bson:"numCylinders"`
	MaxRpm           uint16 `json:"maxRpm" bson:"maxRpm"`
	ManufacturerCode string `json:"manufacturerCode" bson:"manufacturerCode"`
}

type CarFuelFigures struct {
	Speed            uint16  `json:"speed" bson:"speed"`
	Mpg              float64 `json:"mpg" bson:"mpg"`
	UsageDescription string  `json:"usageDescription" bson:"usageDescription"`
}

type CarPerfomanceFigures struct {
	OctaneRating uint16                            `json:"octaneRating" bson:"octaneRating"`
	Acceleration *CarPerfomanceFiguresAcceleration `json:"acceleration" bson:"acceleration"`
}

type CarPerfomanceFiguresAcceleration struct {
	Mph     uint16  `json:"mph" bson:"mph"`
	Seconds float64 `json:"seconds" bson:"seconds"`
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
