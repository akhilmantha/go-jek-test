package parking

import (
  "container/heap"
)

type parkingHeap struct {
  container *pHeap
}

func newParkingHeap(capacity int) *parkingHeap {
  p := &parkingHeap{
    container: &pHeap{},
  }
  for i := 1; i <= capacity; i++ {
    p.container.Push(i)
  }
  heap.Init(p.container)
  return p
}

func (p *parkingHeap) Push(x int) {
  heap.Push(p.container, x)
}

func (p *parkingHeap) Pop() int {
  x := heap.Pop(p.container)
  return x.(int)
}

func (p *parkingHeap) Len() int {
  return p.container.Len()
}

type pHeap []int

func (p pHeap) Len() int {
  return len(p)
}

func (p pHeap) Swap(i, j int) {
  p[i], p[j] = p[j], p[i]
}

func (p pHeap) Less(i, j int) bool {
  return p[i] < p[j]
}

func (p *pHeap) Push(x interface{}) {
  *p = append(*p, x.(int))
}

func (p *pHeap) Pop() interface{} {
  n := p.Len()
  old := *p
  elem := old[n-1]
  *p = old[0:n-1]
  return elem
}
