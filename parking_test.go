package parkingLot

import (
	"testing"
)

func TestParkingAvailability(t *testing.T) {
	t.Run("Expect to park a vehicle", func(t *testing.T) {
		vehicle := Vehicle{vehicleType: "Car",number: "KA05HF2541"}
		result := vehicle.Park()

		if result != "Parked"{
			t.Fail()
			t.Fatalf("Vehicle was not parked")
		}
	})
}
