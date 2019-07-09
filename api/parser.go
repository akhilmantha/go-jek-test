package api

import (
  "strings"
  "strconv"
  // "log"
)

const defaultSeparator = " "

func parse(input string) (Command, error) {
  tokens := strings.Split(input, defaultSeparator)
  if len(tokens) == 0 {
    return newCommand("", []interface{}{}), ErrParse
  }
  switch tokens[0] {
  case CreateParkingCommand:
    if len(tokens) < 2 {
      return newCommand(CreateParkingCommand, []interface{}{}), ErrParseMissingArgs
    }
    size , err := strconv.Atoi(tokens[1])
    if err != nil {
      return newCommand(CreateParkingCommand, []interface{}{}), ErrParseBadArgument
    }
    return newCommand(
      CreateParkingCommand,
      []interface{}{
        size,
      },
    ), nil
  case ParkCommand:
    // log.Printf("tokens: %v\n",tokens)
    if len(tokens) < 3 {
      return newCommand(ParkCommand, []interface{}{}), ErrParseMissingArgs
    }
    return newCommand(
      ParkCommand,
      []interface{}{
        tokens[1],
        tokens[2],
      },
    ), nil
  case LeaveCommand:
    if len(tokens) < 2 {
      return newCommand(LeaveCommand, []interface{}{}), ErrParseMissingArgs
    }
    slot , err := strconv.Atoi(tokens[1])
    if err != nil {
      return newCommand(LeaveCommand, []interface{}{}), ErrParseBadArgument
    }
    return newCommand(
      LeaveCommand,
      []interface{}{
        slot,
      },
    ), nil
  case StatusCommand:
    return newCommand(StatusCommand, []interface{}{}), nil
  case QueryRegistrationSlotCommand:
    if len(tokens) < 2 {
      return newCommand(QueryRegistrationSlotCommand, []interface{}{}), ErrParseMissingArgs
    }
    return newCommand(
      QueryRegistrationSlotCommand,
      []interface{}{
        tokens[1],
      },
    ), nil
  case QueryColorSlotCommand:
    if len(tokens) < 2 {
      return newCommand(QueryColorSlotCommand, []interface{}{}), ErrParseMissingArgs
    }
    return newCommand(
      QueryColorSlotCommand,
      []interface{}{
        tokens[1],
      },
    ), nil
  case QueryColorRegistrationCommand:
    if len(tokens) < 2 {
      return newCommand(QueryColorRegistrationCommand, []interface{}{}), ErrParseMissingArgs
    }
    return newCommand(
      QueryColorRegistrationCommand,
      []interface{}{
        tokens[1],
      },
    ), nil
  }
  return newCommand("", []interface{}{}), ErrParseBadCommand
}
