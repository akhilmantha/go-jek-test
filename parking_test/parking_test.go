package parking_test

import (
  "testing"
  prk "parking_lot/parking"
)

func testParking(factory prk.ParkingFactory, config prk.ParkingConfig, t *testing.T) {
  parking, err := factory.New(config)
  if err != nil {
    t.Fatal(err)
  }
  t.Run("test_empty", testParkEmpty(parking))
}

func testParkEmpty(parking prk.Parking) func (t *testing.T) {
  car := prk.Car{
    Color: "Red",
    Registration: "ABC",
  }
  return func (t *testing.T) {
    slot, err := parking.Park(&car)
    if err != nil {
      t.Log(err)
      t.FailNow()
    }
    if slot.Idx != 1 {
      t.Log("initial slot should be 1")
      t.Fail()
    }
  }
}

func TestParkingInMem(t* testing.T) {
  config := prk.ParkingConfig{
    Capacity: 5,
  }
  testParking(&prk.ParkingInMemFactory{}, config, t)
}
