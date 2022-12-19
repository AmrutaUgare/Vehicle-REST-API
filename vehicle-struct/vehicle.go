package vehiclestruct

type Vehicle struct {
	VehicleId    int    `json:"vehicle-id,omitempty"` // int32
	VehicleName  string `json:"vehicle-name,omitempty"`
	VehicleType  string `json:"vehicle-type,omitempty"`
	Manufcrature string `json:"manufacture,omitempty"`
}
