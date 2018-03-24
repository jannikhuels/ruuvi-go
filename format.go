package ruuvi

import (
	"time"

	"github.com/paypal/gatt"
)

type RuuviSensorFormatVersion int

const (
	FORMAT1       RuuviSensorFormatVersion = 1
	FORMAT2       RuuviSensorFormatVersion = 2
	FORMAT3       RuuviSensorFormatVersion = 3
	FORMAT4       RuuviSensorFormatVersion = 4
	FORMAT5       RuuviSensorFormatVersion = 5
	FORMAT6       RuuviSensorFormatVersion = 6
	FORMATUNKNOWN RuuviSensorFormatVersion = -1
)

type RuuviSensorData struct {
	Temperature   float64
	Humidity      float64
	Pressure      uint32
	Battery       uint16
	Address       string
	AccelerationX int16
	AccelerationY int16
	AccelerationZ int16
	TimeStamp     time.Time
}

type RuuviSensorFormat interface {
	Version() RuuviSensorFormatVersion

	GetSensorData(a *gatt.Advertisement) *RuuviSensorData
}
