package parking

type parkingInMem struct {
  capacity int
  heap *parkingHeap
  container map[int]Slot
  colorIndex map[string](map[int]bool)
  regIndex map[string]int
}

func (p *parkingInMem) Name() string {
  return "parking.in_mem"
}

func (p *parkingInMem) Park(car *Car) (Slot, error) {
  if len(p.container) == p.capacity {
    return emptySlot(), ErrParkingFull
  }
  next := p.heap.Pop()
  slot := Slot{
    Idx: next,
    Car: car,
  }
  if _, ok := p.colorIndex[car.Color]; !ok {
    p.colorIndex[car.Color] = make(map[int]bool)
  }
  (p.colorIndex[car.Color])[next] = true
  p.regIndex[car.Registration] = next
  p.container[next] = slot;
  return slot, nil
}

func (p *parkingInMem) Leave(idx int) error {
  return ErrNotImplemented
}

func (p *parkingInMem) Status() ([]Slot, error) {
  return []Slot{}, ErrNotImplemented
}

func (p *parkingInMem) QueryColor(color string) ([]Slot, error) {
  return []Slot{}, ErrNotImplemented
}

func (p *parkingInMem) QueryRegistration(reg string) (Slot, error) {
  return emptySlot(), ErrNotImplemented
}
