package parking

type Slot struct {
  Idx int
  Car *Car
}

func emptySlot() Slot {
  return Slot{
    Idx: 0,
    Car: nil,
  }
}
