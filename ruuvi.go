package ruuvi

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log"

	"github.com/paypal/gatt"
	"github.com/paypal/gatt/examples/option"
)

var newDataHandler func(*RuuviSensorData)

type RuuviOperationMode string

const (
	URL  RuuviOperationMode = "URL"
	RAW  RuuviOperationMode = "RAW"
	NONE RuuviOperationMode = "NoRuuviDevice"
)

func getRuuviOperationMode(a *gatt.Advertisement) RuuviOperationMode {
	if len(a.ManufacturerData) == 20 && binary.LittleEndian.Uint16(a.ManufacturerData[0:2]) == 0x0499 {
		return RAW
	} else {
		// TODO: Test for URL Operation Mode
		return NONE
	}
}

func getRuuviSensorFormat(a *gatt.Advertisement) (RuuviSensorFormat, error) {
	var r RuuviSensorFormat
	if getRuuviOperationMode(a) == RAW {
		switch int(a.ManufacturerData[2]) {
		case 3:
			r = RuuviSensorFormat3{}
		}
	} else {
		return r, errors.New("unknown Ruuvi Operation Mode")
	}
	return r, nil
}

func getRuuviSensorData(a *gatt.Advertisement) (*RuuviSensorData, error) {
	rsf, err := getRuuviSensorFormat(a)
	if err != nil {
		return rsf.GetSensorData(a), nil
	}
	return nil, err
}

func onStateChanged(d gatt.Device, s gatt.State) {
	fmt.Println("State:", s)
	switch s {
	case gatt.StatePoweredOn:
		fmt.Println("scanning...")
		d.Scan([]gatt.UUID{}, false)
		return
	default:
		d.StopScanning()
	}
}

func onPeriphDiscovered(p gatt.Peripheral, a *gatt.Advertisement, rssi int) {
	data, err := getRuuviSensorData(a)
	if err != nil {
		newDataHandler(data)
	}
}

func Init(fp func(*RuuviSensorData)) {
	d, err := gatt.NewDevice(option.DefaultClientOptions...)
	if err != nil {
		log.Fatalf("Failed to open device, err: %s\n", err)
		return
	}

	newDataHandler = fp

	d.Handle(gatt.PeripheralDiscovered(onPeriphDiscovered))
	d.Init(onStateChanged)
	select {}
}
