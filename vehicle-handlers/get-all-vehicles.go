package vehiclehandlers

import (
	"encoding/json"
	"net/http"
	"vehicle/vehicle-data"
)

func GetAllVehicles(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(vehicledata.Vehicles)
}
