package vehiclehandlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	vehicledata "vehicle/vehicle-data"
)

func DeleteVehicle(response http.ResponseWriter, request *http.Request) {

	//read the dynamic id parameter
	vehicleVars := mux.Vars(request)
	id, _ := strconv.Atoi(vehicleVars["vehicle-id"])

	// iterate over all mock vehicles
	for index, vehicle := range vehicledata.Vehicles {
		if vehicle.VehicleId == id {
			vehicledata.Vehicles = append(vehicledata.Vehicles[:index], vehicledata.Vehicles[index+1:]...)

			response.WriteHeader(http.StatusOK)
			response.Header().Add("Content-Type", "application/json")
			json.NewEncoder(response).Encode("Deleted")
			break
		}
		// delete vehicle and send response if the vehicle id matches dynamic id

	}
}
