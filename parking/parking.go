package parking

/*
Parking describes the interface to interact
with a parking lot.
*/
type Parking interface {
  /*
  Name should return the name of the Parking implementation
  */
  Name() string
  /*
  Park should be used to park a car in the parking lot. It
  should return the slot object with the proper index when
  parking is successful. When the parking is full it should
  return an ErrParkingFull as the error.
  */
  Park(car *Car) (Slot, error)
  /*
  Leave unparks a car from the parking lot. It must return
  ErrBadSlot if an invalid slot is passed, or ErrNotFound
  if the required slot is empty, otherwise nil.
  */
  Leave(int) error
  /*
  Status returns an ordered list of slots which have parked
  cars.
  */
  Status() ([]Slot, error)
  /*
  QueryColor must return an ordered list of slots which have
  cars of a given color.
  */
  QueryColor(color string) ([]Slot, error)
  /*
  QueryRegistration must return the slot which contains
  the car with the given registration, otherwise ErrNotFound.
  */
  QueryRegistration(reg string) (Slot, error)
}

/*
ParkingFactory must be able to create an instance of Parking
from the supplied ParkingConfig
*/
type ParkingFactory interface {
  New(ParkingConfig) (Parking, error)
}

/*
ParkingConfig must be implemented by types that are passed as
configuration to ParkingFactory.
*/
type ParkingConfig interface {
  /*
  Capacity should return the capacity described by the config
  */
  Capacity() int
  /*
  SetCapacity should modify the capacity of the current config
  */
  SetCapacity(int)
}

/*
BaseParkingConfig describes a struct which implements
ParkingConfig. It can be embedded into other struct for more
complex configuration descriptions.
*/
type BaseParkingConfig struct {
  capacity int
}

func (b *BaseParkingConfig) Capacity() int {
  return b.capacity
}

func (b *BaseParkingConfig) SetCapacity(capacity int) {
  b.capacity = capacity
}

/*
ParkingInMemFactory contructs an in memory Parking object.
*/
type ParkingInMemFactory struct {}

func (p *ParkingInMemFactory) New(config ParkingConfig) (Parking, error) {
  return &parkingInMem{
    capacity: config.Capacity(),
    heap: newParkingHeap(config.Capacity()),
    container: make(map[int]*Car),
    colorIndex: make(map[string](map[int]bool)),
    regIndex: make(map[string]int),
  }, nil
}
