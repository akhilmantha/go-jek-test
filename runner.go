package main

import (
  "parking_lot/api"
  prk "parking_lot/parking"
  "fmt"
  "text/tabwriter"
  "os"
)

type Runner struct {
  api api.CommandApi
  parkingFactory prk.ParkingFactory
  config prk.ParkingConfig
  parking prk.Parking
}

func (r *Runner) Run() error {
  inputChan, err := r.api.Chan()
  if err != nil {
    return err
  }
  for cmd := range inputChan {
    switch cmd.Cmd {
    case api.ExitCommand:
      return nil
    case api.CreateParkingCommand:
      r.config.Capacity = cmd.Args[0].(int)
      r.parking, err = r.parkingFactory.New(r.config)
      if err != nil {
        return err
      }
      fmt.Printf("Created a parking lot with %d slots\n", r.config.Capacity)
    case api.ParkCommand:
      registration, color := cmd.Args[0].(string), cmd.Args[1].(string)
      car := prk.Car {
        Color: color,
        Registration: registration,
      }
      slot, err := r.parking.Park(&car)
      if err != nil {
        if err == prk.ErrParkingFull {
          fmt.Printf("Sorry, parking lot is full\n")
          continue
        }
        return err
      }
      fmt.Printf("Allocated slot number: %d\n", slot.Idx)
    case api.LeaveCommand:
      slotIdx := cmd.Args[0].(int)
      err := r.parking.Leave(slotIdx)
      if err != nil {
        if err == prk.ErrBadSlot {
          fmt.Printf("Invalid slot number: %d\n", slotIdx)
          continue
        }
        if err == prk.ErrNotFound {
          fmt.Printf("Not Found\n")
          continue
        }
        return err
      }
      fmt.Printf("Slot number %d is free\n", slotIdx)
    case api.StatusCommand:
      slots, err := r.parking.Status()
      if err != nil {
        return err
      }
      if len(slots) == 0 {
        fmt.Printf("Not Found\n")
        continue
      }
       w := new(tabwriter.Writer)
      // w := os.Stdout
       w.Init(os.Stdout, 0, 8, 0, '\t', 0)
      fmt.Fprintf(w, "Slot No.\tRegistration No\tColour")
      for _, slot := range slots {
        fmt.Fprintf(
          w,
          "\n%d\t%s\t%s",
          slot.Idx,
          slot.Car.Registration,
          slot.Car.Color,
        )
      }
       w.Flush()
      fmt.Printf("\n")
    case api.QueryColorRegistrationCommand:
      slots, err := r.parking.QueryColor(cmd.Args[0].(string))
      if err != nil {
        return err
      }
      if len(slots) == 0 {
        fmt.Printf("Not Found\n")
        continue
      }
      first := true
      for _, slot := range slots {
        if first {
          fmt.Printf("%s", slot.Car.Registration)
          first = false
        } else {
          fmt.Printf(", %s", slot.Car.Registration)
        }
      }
      fmt.Printf("\n")
    case api.QueryColorSlotCommand:
      slots, err := r.parking.QueryColor(cmd.Args[0].(string))
      if err != nil {
        return err
      }
      if len(slots) == 0 {
        fmt.Printf("Not Found\n")
        continue
      }
      first := true
      for _, slot := range slots {
        if first {
          fmt.Printf("%d", slot.Idx)
          first = false
        } else {
          fmt.Printf(", %d", slot.Idx)
        }
      }
      fmt.Printf("\n")
    case api.QueryRegistrationSlotCommand:
      slot, err := r.parking.QueryRegistration(cmd.Args[0].(string))
      if err != nil {
        if err == prk.ErrNotFound {
          fmt.Printf("Not Found\n")
          continue
        }
        return err
      }
      fmt.Printf("%d", slot.Idx)
      fmt.Printf("\n")
    }
  }
  return nil
}
