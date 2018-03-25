package main

import (
	"fmt"

	"github.com/jannikhuels/ruuvi-go"
)

func onNewRuuviSensorData(rsp *ruuvi.RuuviSensorData) {
	fmt.Println(rsp)
}

func main() {
	ruuvi.Init(onNewRuuviSensorData)
}
