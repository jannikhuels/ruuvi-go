package ruuvi

import (
	"bytes"
	"encoding/binary"
	"log"

	"github.com/paypal/gatt"
)

type RuuviSensorFormat3 struct {
	ManufacturerID      uint16
	DataFormat          uint8
	Humidity            uint8
	Temperature         uint8
	TemperatureFraction uint8
	Pressure            uint16
	AccelerationX       int16
	AccelerationY       int16
	AccelerationZ       int16
	BatteryVoltageMv    uint16
}

func (s RuuviSensorFormat3) Version() RuuviSensorFormatVersion {
	return 3
}

func parseTemperature(t uint8, f uint8) float64 {
	var mask uint8
	mask = (1 << 7)
	isNegative := (t & mask) > 0
	temp := float64(t&^mask) + float64(f)/100.0
	if isNegative {
		temp *= -1
	}
	return temp
}

func (s RuuviSensorFormat3) GetSensorData(a *gatt.Advertisement) *RuuviSensorData {
	reader := bytes.NewReader(a.ManufacturerData)
	result := RuuviSensorFormat3{}
	err := binary.Read(reader, binary.BigEndian, &result)

	if err != nil {
		log.Println("Error reading sensor data: ", err)
		return nil
	}

	sensorData := RuuviSensorData{}
	sensorData.Temperature = parseTemperature(result.Temperature, result.TemperatureFraction)
	sensorData.Humidity = float64(result.Humidity) / 2.0
	sensorData.Pressure = uint32(result.Pressure) + 50000
	sensorData.Battery = result.BatteryVoltageMv
	sensorData.AccelerationX = result.AccelerationX
	sensorData.AccelerationY = result.AccelerationY
	sensorData.AccelerationZ = result.AccelerationZ

	return &sensorData
}
