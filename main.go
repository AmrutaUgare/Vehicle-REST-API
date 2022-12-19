package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"vehicle/vehicle-handlers"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/vehicles", vehiclehandlers.GetAllVehicles).Methods(http.MethodGet)
	router.HandleFunc("/vehicles/{vehicle-id}", vehiclehandlers.GetSingleVehicle).Methods(http.MethodGet)
	router.HandleFunc("/vehicles", vehiclehandlers.AddVehicle).Methods(http.MethodPost)
	router.HandleFunc("/vehicles/{vehicle-id}", vehiclehandlers.UpdateVehicleAllColumn).Methods(http.MethodPut)
	router.HandleFunc("/vehicles/{vehicle-id}", vehiclehandlers.DeleteVehicle).Methods(http.MethodDelete)

	log.Println("API is running....")
	http.ListenAndServe(":4000", router)
}
