package ruuvi

import (
	"testing"

	"github.com/paypal/gatt"
)

func getAdvertisement(om RuuviOperationMode) gatt.Advertisement {
	if om == RAW {
		return gatt.Advertisement{ManufacturerData: []byte{153, 4, 3, 64, 23, 50, 194, 106, 255, 232, 255, 248, 4, 16, 12, 67, 0, 0, 0, 0}}
	} else {
		return gatt.Advertisement{}
	}
}

func TestRuuviOperationMode(t *testing.T) {
	var aRAW gatt.Advertisement = getAdvertisement(RAW)
	var aNONE gatt.Advertisement = getAdvertisement(NONE)

	var om = getRuuviOperationMode(&aRAW)
	if om != RAW {
		t.Errorf("Operation Mode should be %s, but is %s instead.", RAW, om)
	}

	om = getRuuviOperationMode(&aNONE)
	if om != NONE {
		t.Errorf("Operation Mode should be %s, but is %s instead.", RAW, om)
	}
}

func TestRuuviFormat(t *testing.T) {
	var aRAW gatt.Advertisement = getAdvertisement(RAW)
	var aNONE gatt.Advertisement = getAdvertisement(NONE)

	var r RuuviSensorFormat
	var err error
	r, err = getRuuviSensorFormat(&aRAW)
	if err != nil && r.Version() != FORMAT3 {
		t.Errorf("Format Version should be %d, but is %d instead.", FORMAT3, r.Version())
	}

	r, err = getRuuviSensorFormat(&aNONE)
	if r != nil && r.Version() != FORMATUNKNOWN {
		t.Errorf("Format Version should be %d, but is %d instead.", FORMATUNKNOWN, r.Version())
	}
}
