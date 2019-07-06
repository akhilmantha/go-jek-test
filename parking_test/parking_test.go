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
  config.Capacity = 2
  parking, err = factory.New(config)
  if err != nil {
    t.Fatal(err)
  }
  t.Run("test_park_and_leave", testParkAndLeave(parking))
  parking, err = factory.New(config)
  if err != nil {
    t.Fatal(err)
  }
  t.Run("test_query_color", testQueryColor(parking))
  config.Capacity = 10
  parking, err = factory.New(config)
  if err != nil {
    t.Fatal(err)
  }
  t.Run("test_status", testStatus(parking))
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

func testParkAndLeave(parking prk.Parking) func (t *testing.T) {
  return func (t *testing.T) {
    car1 := &prk.Car{
      Color: "Red",
      Registration: "A",
    }
    car2 := &prk.Car{
      Color: "Green",
      Registration: "B",
    }
    slot, err := parking.Park(car1)
    if err != nil {
      t.Log(err)
      t.FailNow()
    }
    if slot.Idx != 1 {
      t.Log("initial slot should be 1")
      t.Fail()
    }
    slot, err = parking.Park(car2)
    if err != nil {
      t.Log(err)
      t.FailNow()
    }
    if slot.Idx != 2 {
      t.Log("slot should be 2")
      t.Fail()
    }
    err = parking.Leave(1)
    if err != nil {
      t.Log(err)
      t.FailNow()
    }
    slot, err = parking.Park(car2)
    if err != nil {
      t.Log(err)
      t.FailNow()
    }
    if slot.Idx != 1 {
      t.Log("slot should be 1")
      t.Fail()
    }
  }
}

func testQueryColor(parking prk.Parking) func (t *testing.T) {
  return func (t *testing.T) {}
}

func testStatus(parking prk.Parking) func (t *testing.T) {
  return func (t *testing.T) {}
}

func TestParkingInMem(t* testing.T) {
  config := prk.ParkingConfig{}
  testParking(&prk.ParkingInMemFactory{}, config, t)
}
