package parking

type Parking interface {
  Name() string
  Park(car *Car) (Slot, error)
  Leave(int) error
  Status() ([]Slot, error)
  QueryColor(color string) ([]Slot, error)
  QueryRegistration(reg string) (Slot, error)
}

type ParkingFactory interface {
  New(ParkingConfig) (Parking, error)
}

type ParkingConfig struct {
  Capacity int
}

type ParkingInMemFactory struct {}

func (p *ParkingInMemFactory) New(config ParkingConfig) (Parking, error) {
  return &parkingInMem{
    capacity: config.Capacity,
    heap: newParkingHeap(config.Capacity),
    container: make(map[int]Slot),
    colorIndex: make(map[string](map[int]bool)),
    regIndex: make(map[string]int),
  }, nil
}
