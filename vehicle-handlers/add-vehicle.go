package vehiclehandlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	vehicledata "vehicle/vehicle-data"
	vehiclestruct "vehicle/vehicle-struct"
)

func AddVehicle(response http.ResponseWriter, request *http.Request) {

	// read to request body
	defer request.Body.Close()

	vehicleBody, err := ioutil.ReadAll(request.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var vehicle vehiclestruct.Vehicle
	json.Unmarshal(vehicleBody, &vehicle)

	// append to the vehicle vehicle-data

	vehicledata.Vehicles = append(vehicledata.Vehicles, vehicle)

	// send a 201 created response
	response.WriteHeader(http.StatusCreated)
	response.Header().Add("Content-Type", "application/json")
	json.NewEncoder(response).Encode("Created")

}
