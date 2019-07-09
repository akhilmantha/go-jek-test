TEST_MODULES:=parking_test api
DEPS:=parking api

build:
	go build -v

test:
	$(foreach mod, $(TEST_MODULES), go test -v ./$(mod);)
