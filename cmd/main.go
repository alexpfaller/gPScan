package main

import (
	"github.com/alexpfaller/gPScan/v2/internal/gpscan"
)

func main() {
	gPScan := gpscan.New()
	gPScan.DoPortScanning()
}
