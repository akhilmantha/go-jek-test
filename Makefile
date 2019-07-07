TEST_MODULES:= parking_test

test:
	$(foreach mod, $(TEST_MODULES), go test -v ./$(mod);)
