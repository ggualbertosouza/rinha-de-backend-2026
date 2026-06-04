MAIN_FILE = ./cmd/server/main.go
TEST_FILE = ./cmd/server/loadTest/main.go

run:
	go run $(MAIN_FILE)

test:
	go run $(TEST_FILE)