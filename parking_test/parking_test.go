package parking_test

import (
  "testing"
  "math/rand"
	"encoding/base64"
  prk "parking_lot/parking"
)

const colorSize = 4
var defaultCapacity = 20;

var testColors []string = []string{
	"A",
	"B",
	"C",
	"D",
}

func genRegistration() string {
  buf := make([]byte, 9)
  _, _ = rand.Read(buf)
  return base64.StdEncoding.EncodeToString(buf)
}

func randomColor() string {
	c := rand.Intn(colorSize)
	return testColors[c]
}

func genCar() prk.Car {
	return prk.Car{
		Color: randomColor(),
		Registration: genRegistration(),
	}
}

func testParking(factory prk.ParkingFactory, config prk.ParkingConfig, t *testing.T) {
  config.SetCapacity(defaultCapacity)

  parking, err := factory.New(config)
  if err != nil {
    t.Fatal(err)
  }
  t.Run("test_empty", testParkEmpty(parking))

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

  parking, err = factory.New(config)
  if err != nil {
    t.Fatal(err)
  }
  t.Run("test_status", testStatus(parking))

  parking, err = factory.New(config)
  if err != nil {
    t.Fatal(err)
  }
  t.Run("test_query_registration", testQueryRegistration(parking))
}

func testParkEmpty(parking prk.Parking) func (t *testing.T) {
  return func (t *testing.T) {
    car := genCar()
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

func testQueryRegistration(parking prk.Parking) func (t *testing.T) {
  return func (t *testing.T) {
    saved := make(map[string]prk.Slot)
    for i := 1; i <= defaultCapacity; i++ {
      car := genCar()
      slot, err := parking.Park(&car)
      if err != nil {
        t.Log(err)
        t.FailNow()
      }
      saved[car.Registration] = slot
    }
    for k := range saved {
      slot, err := parking.QueryRegistration(k)
      if err != nil {
        t.Log(err)
        t.FailNow()
      }
      if slot.Idx != saved[k].Idx {
        t.FailNow()
      }
      if slot.Car.Registration != saved[k].Car.Registration {
        t.FailNow()
      }
    }
  }
}

func testParkAndLeave(parking prk.Parking) func (t *testing.T) {
  return func (t *testing.T) {
    car1 := genCar()
    car2 := genCar()
    slot, err := parking.Park(&car1)
    if err != nil {
      t.Log(err)
      t.FailNow()
    }
    if slot.Idx != 1 {
      t.Log("initial slot should be 1")
      t.Fail()
    }
    slot, err = parking.Park(&car2)
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
    slot, err = parking.Park(&car1)
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
  return func (t *testing.T) {
    saved := make(map[string][]prk.Slot)

    for i := 1; i <= 20; i++ {
      car := genCar()
      slot, err := parking.Park(&car)
      if err != nil {
        t.Log(err)
        t.FailNow()
      }
      if _, ok := saved[car.Color]; !ok {
        saved[car.Color] = []prk.Slot{}
      }
      saved[car.Color] = append(saved[car.Color], slot)
    }

    for k := range saved {
      slots, err := parking.QueryColor(k)
      if err != nil {
        t.Log(err)
        t.FailNow()
      }
      if len(slots) != len(saved[k]) {
        t.Logf("number of cars must be %d, got %d", len(saved[k]), len(slots))
        t.FailNow()
      }
      for i := 0; i < len(slots); i++ {
        if slots[i].Car.Registration != saved[k][i].Car.Registration {
          t.Logf(
            "should have registration %s, go %s",
            slots[i].Car.Registration,
            saved[k][i].Car.Registration,
          )
          t.FailNow()
        }
      }
    }
  }
}

func testStatus(parking prk.Parking) func (t *testing.T) {
  return func (t *testing.T) {
    for i := 1; i <= 20; i++ {
      car := genCar()
      slot, err := parking.Park(&car)
      if err != nil {
        t.Log(err)
        t.FailNow()
      }
      if slot.Idx != i {
        t.Logf("slot should be %d", i)
        t.FailNow()
      }
    }

    slots, statusErr := parking.Status()
    if statusErr != nil {
      t.Log(statusErr)
      t.FailNow()
    }
    if len(slots) != 20 {
      t.Log("all 20 slots must be filled")
      t.FailNow()
    }
    for i := 0; i < 20; i++ {
      if slots[i].Idx != i+1 {
        t.Logf("slot should be %d, instead got %d", i, slots[i].Idx)
        t.FailNow()
      }
    }
  }
}

func TestParkingInMem(t* testing.T) {
  config := &prk.BaseParkingConfig{}
  testParking(&prk.ParkingInMemFactory{}, config, t)
}
