package vehiclehandlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	vehicledata "vehicle/vehicle-data"
)

func GetSingleVehicle(response http.ResponseWriter, request *http.Request) {
	// read dynamic id parameter

	vehicleVars := mux.Vars(request)
	getId, _ := strconv.Atoi(vehicleVars["vehicle-id"])

	//iterate overall the vehicle-data vehicles
	for _, vehicle := range vehicledata.Vehicles {

		if vehicle.VehicleId == getId {
			// if ids are equal send book as response
			response.WriteHeader(http.StatusOK)
			response.Header().Add("Content-Type", "application/json")
			json.NewEncoder(response).Encode(vehicle)
			break
		}
	}

}
