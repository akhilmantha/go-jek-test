package api

import (
  "testing"
  "strings"
)

func TestFileCommandApiEmptyFile(t *testing.T) {
  fca := &FileCommandApi{
    out: make(chan Command, 1),
    reader: strings.NewReader(""),
  }
  ch, err := fca.Chan()
  if err != nil {
    t.Log(err)
    t.FailNow()
  }
  go func() {
    _ = fca.Start()
  }()
  select {
  case cmd, ok := <-ch:
    if ok {
      t.Logf("channel should be closed, instead received %v", cmd)
      t.FailNow()
    }
  }
}

func TestFileCommandApiIncompleteCommand(t *testing.T) {
  input := `park abc GREEN
  par`
  fca := &FileCommandApi{
    out: make(chan Command, 1),
    reader: strings.NewReader(input),
  }
  ch, err := fca.Chan()
  if err != nil {
    t.Log(err)
    t.FailNow()
  }
  var parseError error
  go func() {
    parseError = fca.Start()
  }()
  select {
  case cmd := <- ch:
    if cmd.Cmd != ParkCommand {
      t.Logf("bad command")
      t.FailNow()
    }
    if len(cmd.Args) != 2 {
      t.Logf("should have 2 args")
      t.FailNow()
    }
  }
  select {
  case cmd, ok := <-ch:
    if ok {
      t.Logf("channel should be closed, instead received %v", cmd)
      t.FailNow()
    }
  }
  if parseError != ErrParseBadCommand {
    t.FailNow()
  }
}

func TestFileCommandApiMissingArguments(t *testing.T) {
  input := `park abc`
  fca := &FileCommandApi{
    out: make(chan Command, 1),
    reader: strings.NewReader(input),
  }
  ch, err := fca.Chan()
  if err != nil {
    t.Log(err)
    t.FailNow()
  }
  var parseError error
  go func() {
    parseError = fca.Start()
  }()
  select {
  case cmd, ok := <-ch:
    if ok {
      t.Logf("channel should be closed, instead received %v", cmd)
      t.FailNow()
    }
  }
  if parseError != ErrParseMissingArgs {
    t.Logf("bad error, expected %v, got %v", ErrParseMissingArgs, parseError)
    t.FailNow()
  }
}

func TestFileCommandApiBadCommand(t *testing.T) {
  input := `park abc GREEN
  abc 3 34 `
  fca := &FileCommandApi{
    out: make(chan Command, 1),
    reader: strings.NewReader(input),
  }
  ch, err := fca.Chan()
  if err != nil {
    t.Log(err)
    t.FailNow()
  }
  var parseError error
  go func() {
    parseError = fca.Start()
  }()
  select {
  case cmd := <- ch:
    if cmd.Cmd != ParkCommand {
      t.Logf("bad command")
      t.FailNow()
    }
    if len(cmd.Args) != 2 {
      t.Logf("should have 2 args")
      t.FailNow()
    }
  }
  select {
  case cmd, ok := <-ch:
    if ok {
      t.Logf("channel should be closed, instead received %v", cmd)
      t.FailNow()
    }
  }
  if parseError != ErrParseBadCommand {
    t.FailNow()
  }
}
