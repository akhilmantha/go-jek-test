package api

import (
  "fmt"
)

var ErrParseBadCommand = fmt.Errorf("bad command")
var ErrParseMissingArgs = fmt.Errorf("missing args")
var ErrParse = fmt.Errorf("parse error")
var ErrParseBadArgument = fmt.Errorf("bad argument") 
