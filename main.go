package main

import (
  "os"
  prk "parking_lot/parking"
  "parking_lot/api"
  "log"
)

func main() {
  factory := &prk.ParkingInMemFactory{}
  input := api.NewStdinApi()
  if len(os.Args) > 1 {
    var err error
    input, err = api.NewFileCommandApi(os.Args[1])
    if err != nil {
      log.Fatal(err)
    }
  }
  config := prk.ParkingConfig{}
  go input.Start()
  (&Runner{
    api: input,
    parkingFactory: factory,
    config: config,
    parking: nil,
  }).Run()
}
