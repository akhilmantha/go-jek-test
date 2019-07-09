TEST_MODULES:=parking_test api
DEPS:=parking api

build:
	$(foreach dep, $(DEPS), go build ./$(dep);)

test:
	$(foreach mod, $(TEST_MODULES), go test -v ./$(mod);)
