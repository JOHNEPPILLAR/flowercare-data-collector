// Package fcServer - Interact with flowercare BLE API's
package fcServer

/*
import (
	"fmt"

	utils "github.com/JOHNEPPILLAR/utility"
	"tinygo.org/x/bluetooth"
)

func processDevices() {

	for _, device := range devices {

		// Connect to device
		utils.Logger.Debug(fmt.Sprint("[", device.Address.String(), "] Connecting..."))
		connectedDevice, err := adapter.Connect(device.Address, bluetooth.ConnectionParams{})
		if err != nil {
			utils.Logger.Error(fmt.Sprint("[", device.Address.String(), "] ", err.Error()))
			continue
		}

		utils.Logger.Debug(fmt.Sprint("[", device.Address.String(), "] Connected"))

		// Get all services
		utils.Logger.Debug(fmt.Sprint("[", device.Address.String(), "] Discovering services/characteristics"))
		services, err := connectedDevice.DiscoverServices(nil)
		if err != nil || len(services) == 0 {
			utils.Logger.Error(fmt.Sprint("[", device.Address.String(), "] Error getting services: ", err.Error()))
			continue
		}

		// Setup data to save structure
		var dataToSave FlowercareData
		dataToSave.Address = device.Address.String()

		// Get battery data
		firmwareData, err := getFirmwareData(services, connectedDevice, device.Address.String())
		if err != nil {
			utils.Logger.Error(fmt.Sprint("[", device.Address.String(), "] Error getting battery data: ", err.Error()))
			continue
		}
		dataToSave.Battery = uint16(firmwareData.Battery)

		// Get sensor data
		sensorData, err := getSensorData(services, connectedDevice, device.Address.String())
		if err != nil {
			utils.Logger.Error(fmt.Sprint("[", device.Address.String(), "] Error getting battery data: ", err.Error()))
			continue
		}
		dataToSave.Timestamp = sensorData.Timestamp
		dataToSave.Brightness = sensorData.Brightness
		dataToSave.Conductivity = sensorData.Conductivity
		dataToSave.Moisture = sensorData.Moisture
		dataToSave.Temperature = sensorData.Temperature

		// Disconnect device
		err = connectedDevice.Disconnect()
		if err != nil {
			utils.Logger.Error(err.Error())
		}

		// Save data
		saveDeviceData(dataToSave)
	}
}
*/
