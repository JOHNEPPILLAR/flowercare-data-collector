// Package fcServer - Interact with flowercare BLE API's
package fcServer

/*
import (
	"context"
	"fmt"

	utils "github.com/JOHNEPPILLAR/utility"
)

func saveDeviceData(dataToSave FlowercareData) {

	utils.Logger.Debug(fmt.Sprint("[", dataToSave.Address, "] Saving data..."))

	// Connect to DB
	ctx := context.Background()
	dbClient, dbConnection, err := utils.ConnectToDB(ctx, "flowercare", "sensor_data")
	if err != nil {
		utils.Logger.Error(fmt.Sprint("[", dataToSave.Address, "] Failed to save data"))
		return
	}

	// Disconnect from DB once operation is complete
	defer utils.DisconnectFromDB(ctx, dbClient)

	// Save data
	_, err = dbConnection.InsertOne(ctx, dataToSave)
	if err != nil {
		utils.Logger.Error(fmt.Sprint("[", dataToSave.Address, "] Failed to save data"))
		return
	}

	utils.Logger.Info(fmt.Sprint("[", dataToSave.Address, "] Saved data"))
}
*/
