// Package fcServer - Interact with flowercare BLE API's
package fcServer

import (
	"fmt"
	"time"

	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter
var devices []bluetooth.ScanResult

func (fcs *FlowerCare) scan() {

	// Empty any stored devices and reset error count
	devices = nil
	fcs.errorCount = 0

	// Enable BLE interface
	fcs.logger.Debug("Enable BLE stack")

	err := adapter.Enable()
	if err != nil {
		fcs.logger.Error(err.Error())
		fcs.errorCount++
		return
	}

	// Set scan stop timer
	stopScanTimer := time.NewTimer(15 * time.Second)

	// Stop scanning when timer expires
	go func() {
		<-stopScanTimer.C

		// Stop scanning
		err := adapter.StopScan()
		if err != nil {
			fcs.logger.Error(err.Error())
		}

		fcs.logger.Debug("Scan complete")

		if len(devices) == 0 {
			fcs.logger.Debug("No flowercare devices found")
			return
		}

		// Process devices
		// fcs.processDevices()
	}()

	// Start scanning
	fcs.logger.Debug("Scanning...")
	err = adapter.Scan(fcs.scanHandler)
	if err != nil {
		fcs.logger.Error(err.Error())
		return
	}
}

func (fcs *FlowerCare) scanHandler(adapter *bluetooth.Adapter, bleDevice bluetooth.ScanResult) {

	flowerCareUUIDString := "0000fe95-0000-1000-8000-00805f9b34fb"
	flowerCareUUID, err := bluetooth.ParseUUID(flowerCareUUIDString)
	if err != nil {
		fcs.logger.Error(err.Error())
		return
	}

	if bleDevice.AdvertisementPayload.HasServiceUUID(flowerCareUUID) {
		deviceFound := false

		for _, device := range devices {
			if device.Address.String() == bleDevice.Address.String() {
				deviceFound = true
				break
			}
		}

		if !deviceFound {
			fcs.logger.Info(fmt.Sprint("[", bleDevice.Address.String(), "] Found flowercare device (RSSI: ", bleDevice.RSSI, ")"))
			devices = append(devices, bleDevice)
		}
	}
}
