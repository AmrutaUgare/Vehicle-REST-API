package vehicledata

import (
	vehiclestruct "vehicle/vehicle-struct"
)

var Vehicles = []vehiclestruct.Vehicle{
	{
		VehicleId:    100,
		VehicleName:  "Baleno",
		VehicleType:  "Car",
		Manufcrature: "Maruti Suzuki",
	},
	{
		VehicleId:    101,
		VehicleName:  "Activa",
		VehicleType:  "Bike",
		Manufcrature: "Honda",
	},
	{
		VehicleId:    102,
		VehicleName:  "Creta",
		VehicleType:  "Car",
		Manufcrature: "Hyundai",
	},
	{
		VehicleId:    103,
		VehicleName:  "Pulser",
		VehicleType:  "Bike",
		Manufcrature: "Bajaj",
	},
}
