package ruuvi

import (
	"testing"

	"github.com/paypal/gatt"
)

func TestVersion(t *testing.T) {
	var v RuuviSensorFormatVersion = RuuviSensorFormat3{}.Version()
	if v != FORMAT3 {
		t.Errorf("RuuviFormatVersion should be 3 but is %d instead", v)
	}
}

func TestSensorData(t *testing.T) {
	var r = RuuviSensorFormat3{}
	var aRAW gatt.Advertisement = getAdvertisement(RAW)

	sd := r.GetSensorData(&aRAW)
	if sd.Temperature != 23.5 {
		t.Errorf("Wrong temperature, expected 23.5 and got %f", sd.Temperature)
	}
	if sd.Humidity != 32.0 {
		t.Errorf("Wrong humidity, expected 32 and got %f", sd.Humidity)
	}
	// TODO: Extend
}
