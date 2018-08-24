package interactor

import (
	"github.com/andreipimenov/carservice/model"
)

// Car implements model.CarInteractor
type Car struct {
	Storage model.CarStorage
}

func NewCar(s model.CarStorage) *Car {
	return &Car{
		Storage: s,
	}
}

func (c *Car) FindBySerialNumber(serialNumber uint64) (*model.Car, error) {
	return c.Storage.FindBySerialNumber(serialNumber)
}
