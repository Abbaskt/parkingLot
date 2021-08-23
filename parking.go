package parkingLot

type Vehicle struct {
	vehicleType string
	number string
}

func (v Vehicle) Park() string {
	return "Parked"
}
