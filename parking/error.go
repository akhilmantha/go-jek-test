package parking

import (
  "fmt"
)

var ErrNotImplemented = fmt.Errorf("not implemented")
var ErrParkingFull = fmt.Errorf("parking full")
var ErrBadSlot = fmt.Errorf("bad slot")
var ErrEmptySlot = fmt.Errorf("empty slot")
var ErrNotFound = fmt.Errorf("not found")
