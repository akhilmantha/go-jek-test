TEST_MODULES:=parking_test api
DEPS:=parking api

build:
	go build -v; mv ./parking_lot bin/parking_lot_exec

test:
	$(foreach mod, $(TEST_MODULES), go test -v ./$(mod);)

ft: build
	bin/run_functional_tests
