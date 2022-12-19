package vehiclehandlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	vehicledata "vehicle/vehicle-data"
	vehiclestruct "vehicle/vehicle-struct"
)

func UpdateVehicleAllColumn(response http.ResponseWriter, request *http.Request) {
	// read dynamic id parameter
	vehicleVars := mux.Vars(request)

	getId, _ := strconv.Atoi(vehicleVars["vehicle-id"])

	// read request body
	defer request.Body.Close()
	vehicleBody, err := ioutil.ReadAll(request.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updateVehicle vehiclestruct.Vehicle
	json.Unmarshal(vehicleBody, &updateVehicle)

	// iterate over all the vehicle-data vehicle
	for index, vehicle := range vehicledata.Vehicles {
		if vehicle.VehicleId == getId {
			vehicle.VehicleName = updateVehicle.VehicleName
			vehicle.VehicleType = updateVehicle.VehicleType
			vehicle.Manufcrature = updateVehicle.Manufcrature

			vehicledata.Vehicles[index] = vehicle

			response.WriteHeader(http.StatusOK)
			response.Header().Add("Content-Type", "application/json")
			json.NewEncoder(response).Encode("Updated")
			break
		}
	}
	// update and send response when vehicle-id matches dynamic id

}
