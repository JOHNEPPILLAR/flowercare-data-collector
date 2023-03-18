// Package fcServer - Interact with flowercare BLE API's
package fcServer

/*
import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"time"

	utils "github.com/JOHNEPPILLAR/utility"
	"tinygo.org/x/bluetooth"
)

// SensorData data of device
type SensorData struct {
	Timestamp    int64
	Temperature  float64
	Moisture     byte
	Brightness   uint16
	Conductivity uint16
}

func getSensorData(services []bluetooth.DeviceService, device *bluetooth.Device, deviceAddress string) (sensors SensorData, err error) {

	// Get sensor data
	utils.Logger.Debug(fmt.Sprint("[", deviceAddress, "] Getting sensor data..."))

	// Enable real time data mode on device
	if !enableRealTimeDataMode(services, deviceAddress) {
		return SensorData{}, err
	}

	// Get data from device
	sensorData, err := readSensorData(services, deviceAddress)
	if err != nil {
		return SensorData{}, err
	}

	return sensorData, nil
}

func enableRealTimeDataMode(services []bluetooth.DeviceService, deviceAddress string) (proceed bool) {

	realTimeDataModeServiceUUID := "00001204-0000-1000-8000-00805f9b34fb"
	realTimeDataModeCharacteristicUUID := "00001a00-0000-1000-8000-00805f9b34fb"

	utils.Logger.Debug(fmt.Sprint("[", deviceAddress, "] Enable realtime data mode"))

	for _, service := range services {
		if service.UUID().String() == realTimeDataModeServiceUUID {
			utils.Logger.Debug(fmt.Sprint("[", deviceAddress, "] Found realtime data mode service"))

			characteristics, err := service.DiscoverCharacteristics(nil)
			if err != nil {
				utils.Logger.Error(fmt.Sprint("[", deviceAddress, "] Error getting realtime data mode characteristics: ", err.Error()))
				return false
			}

			// Characteristics
			for _, characteristic := range characteristics {
				if characteristic.UUID().String() == realTimeDataModeCharacteristicUUID {
					utils.Logger.Debug(fmt.Sprint("[", deviceAddress, "] Found realtime data mode characteristic"))

					dataBuffer := []byte{0xa0, 0x1f}
					_, err := characteristic.WriteWithoutResponse(dataBuffer)
					if err != nil {
						utils.Logger.Error(fmt.Sprint("[", deviceAddress, "] Error setting realtime data mode: ", err.Error()))
						return false
					}

					utils.Logger.Debug(fmt.Sprint("[", deviceAddress, "] Enabled realtime data mode"))
					return true
				}
			}
		}
	}
	utils.Logger.Error(fmt.Sprint("[", deviceAddress, "] Unable to enable realtime data mode"))
	return false
}

func readSensorData(services []bluetooth.DeviceService, deviceAddress string) (sensorData SensorData, err error) {

	realTimeDataServiceUUID := "00001204-0000-1000-8000-00805f9b34fb"
	realTimeDataCharacteristicUUID := "00001a01-0000-1000-8000-00805f9b34fb"

	dataBuffer := make([]byte, 255)

	for _, service := range services {
		if service.UUID().String() == realTimeDataServiceUUID {
			utils.Logger.Debug(fmt.Sprint("[", deviceAddress, "] Found realtime data service"))

			characteristics, err := service.DiscoverCharacteristics(nil)
			if err != nil {
				utils.Logger.Error(fmt.Sprint("[", deviceAddress, "] Error getting realtime data mode characteristics: ", err.Error()))
				return SensorData{}, err
			}

			// Characteristics
			for _, characteristic := range characteristics {
				if characteristic.UUID().String() == realTimeDataCharacteristicUUID {
					utils.Logger.Debug(fmt.Sprint("[", deviceAddress, "] Found realtime data mode characteristic"))

					_, err := characteristic.Read(dataBuffer)
					if err != nil {
						utils.Logger.Error(fmt.Sprint("[", deviceAddress, "] Error getting sensor data: ", err.Error()))
						return SensorData{}, err
					}

					var temperature int16
					var brightness uint16
					var moisture uint8
					var conductivity uint16

					p := bytes.NewBuffer(dataBuffer)
					// TT TT ?? LL LL ?? ?? MM CC CC
					binary.Read(p, binary.LittleEndian, &temperature)
					p.Next(1)
					binary.Read(p, binary.LittleEndian, &brightness)
					p.Next(2)
					binary.Read(p, binary.LittleEndian, &moisture)
					binary.Read(p, binary.LittleEndian, &conductivity)

					sensorData := SensorData{
						Timestamp:    time.Now().Unix(),
						Temperature:  float64(temperature) / 10,
						Brightness:   brightness,
						Moisture:     moisture,
						Conductivity: conductivity,
					}

					utils.Logger.Debug(fmt.Sprint("[", deviceAddress, "] Sensor data: ", sensorData))

					return sensorData, nil
				}
			}
		}
	}

	return SensorData{}, errors.New("No sensor data service found")
}
*/
