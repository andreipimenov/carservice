package storage

import (
	"github.com/andreipimenov/carservice/model"
)

// Car implements model.CarStorage
type Car struct {
	Driver CarDriver
	Source string
}

// CarDriver represents injected driver for concrete DB
type CarDriver interface {
	Get(c string, q interface{}, r interface{}) error
}

func NewCar(d CarDriver, s string) *Car {
	return &Car{
		Driver: d,
		Source: s,
	}
}

func (c *Car) FindBySerialNumber(serialNumber uint64) (*model.Car, error) {
	v := &model.Car{}
	err := c.Driver.Get(c.Source, map[string]interface{}{"serialNumber": serialNumber}, &v)
	return v, err
}
