package parking

import (
  "sort"
//  "log"
)

type parkingInMem struct {
  capacity int
  heap *parkingHeap
  container map[int]*Car
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
  p.container[next] = car;
  return slot, nil
}

func (p *parkingInMem) Leave(idx int) error {
  if idx == 0 || idx > p.capacity {
    return ErrBadSlot
  }
  car, ok := p.container[idx]
  if !ok {
    return ErrNotFound
  }
  delete(p.regIndex, car.Registration)
  delete(p.colorIndex[car.Color], idx)
  delete(p.container, idx)
  p.heap.Push(idx)
  return nil
}

func (p *parkingInMem) Status() ([]Slot, error) {
  result := []Slot{}
  // Order of iteration of a map is now randomized
  // so order of keys must be manually tracked
  keys := make([]int, len(p.container))
  i := 0
  for k := range p.container {
    keys[i] = k
    i++
  }
  sort.Ints(keys)
  for _, k := range keys {
    result = append(result, Slot{
      Idx: k,
      Car: p.container[k],
    })
  }
  return result, nil
}

func (p *parkingInMem) QueryColor(color string) ([]Slot, error) {
  if _, ok := p.colorIndex[color]; !ok {
    return []Slot{}, nil
  }
  keys := []int{}
  result := []Slot{}
  for k := range p.colorIndex[color] {
    keys = append(keys, k)
  }
  sort.Ints(keys)
  for _, k := range keys {
    result = append(
      result,
      Slot{
        Idx: k,
        Car: p.container[k],
      },
    )
  }
  return result, nil
}

func (p *parkingInMem) QueryRegistration(reg string) (Slot, error) {
  if _, ok := p.regIndex[reg]; ! ok {
    return emptySlot(), ErrNotFound
  }
  idx, _ := p.regIndex[reg]
  return Slot{
    Idx: idx,
    Car: p.container[idx],
  }, nil
}
