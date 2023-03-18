// Package fcServer - Interact with flowercare BLE API's
package fcServer

/*
import (
	"errors"
	"fmt"

	utils "github.com/JOHNEPPILLAR/utility"
	"tinygo.org/x/bluetooth"
)

// Firmware data of device
type Firmware struct {
	Version string
	Battery byte
}

func getFirmwareData(services []bluetooth.DeviceService, device *bluetooth.Device, deviceAddress string) (firmware Firmware, err error) {

	// Get battery data
	utils.Logger.Debug(fmt.Sprint("[", deviceAddress, "] Getting battery data"))

	batteryServiceUUID := "00001204-0000-1000-8000-00805f9b34fb"
	batteryCharacteristicUUID := "00001a02-0000-1000-8000-00805f9b34fb"
	dataBuffer := make([]byte, 255)

	// Services
	for _, service := range services {
		if service.UUID().String() == batteryServiceUUID {
			utils.Logger.Debug(fmt.Sprint("[", deviceAddress, "] Found battery service"))

			characteristics, err := service.DiscoverCharacteristics(nil)
			if err != nil {
				utils.Logger.Error(fmt.Sprint("[", deviceAddress, "] Error getting battery characteristic: ", err.Error()))
				return Firmware{}, err
			}

			// Characteristics
			for _, characteristic := range characteristics {
				if characteristic.UUID().String() == batteryCharacteristicUUID {
					utils.Logger.Debug(fmt.Sprint("[", deviceAddress, "] Found battery characteristic"))

					_, err := characteristic.Read(dataBuffer)
					if err != nil {
						utils.Logger.Error(fmt.Sprint("[", deviceAddress, "] Error getting battery data: ", err.Error()))
						return Firmware{}, err
					}

					firmware := Firmware{
						Version: string(dataBuffer[2:]),
						Battery: dataBuffer[0],
					}

					utils.Logger.Debug(fmt.Sprint("[", deviceAddress, "] Firmware: ", firmware))

					return firmware, nil
				}
			}
		}
	}
	return Firmware{}, errors.New("No battery service found")
}
*/
