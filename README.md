# Parking Lot

Implementation of automated parking lot.

Language : GoLang

Version : go1.11.5


### Experience with GoLang :
This is the first time I've used go to build something. And I've seen a lot of support for Concurrency and Parallelism. Writing parallelised solutions becomes an interesting design challenge and the implementation goes smoothly.

And I have found that go is adequately mature and stable.

## Architecture

Project is made up of following main components :

## 1. Parking :

   This is the main component. Parking Lot is responsible for managing all the parking related operations such as Create_Parking_Lot, Park Vehicle, Remove Vehicle etc.

     Properties:
      1. parking capacity
      2. List of Spots

     Functions for :  
      1. Create Parking
      2. Park Vehicle
      3. Remove Vehicle
      4. Check Parking Status
      5. Get all Cars by Colour
      6. Get all Slots by Registration Number

## 2. Slot :

  Slot is the place in the Parking Lot where a Vehicle can be parked. Slot have a unique slot number. There can only be one Vehicle at a spot at a time.

    Properties:
      1. Slot Number
      2. Status
      3. Vehicle

    Functions:
      1. Add Vehicle
      2. Remove Vehicle

## 3. Vehicle :

  Vehicle can be parked at the Spot. Vehicle have a Registration Number and Colour

    Properties:
      1. Registration Number
      2. Colour

## 4. Makefile :

  Instead of writing commands manually you can execute them from a bash script or a makefile.
  You can keep all your common tasks together at makefile.

##5. Runner :

Runner accepts an instance of CommandApi, ParkingFactory, and ParkingConfig, and runs the parsed commands against
instances of Parking created using the factory.

## Build

Navigate to 'parking_lot' directory
```
parking_lot$ ./bin/setup
```
run 'make build'

## Run

Navigate to 'parking_lot' directory

## Test

Tests are important to implement to have code coverage and check the working of your solution for different scenarios.

I also implemented a hook where to commit, your tests need to pass.
 - Hook for tests to run before committing the code
 - Functional Tests
 - File command api tests

 run 'make test' and 'make ft' for functional tests

```
parking_lot$ ./bin/run_functional_tests
```


## Features :

- [x] Configurable Commands, easy to create new commands
- [x] Indexing for registration number and colour
- [x] File Support
- [x] Command Helper
- [x] Modular Design
- [x] Functional tests and api tests
