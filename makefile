LB_FILE = ./cmd/lb/main.go
MAIN_FILE = ./cmd/server/main.go
TEST_FILE = ./cmd/server/loadTest/main.go

run:
	go run $(MAIN_FILE)

run-lb:
	go run $(LB_FILE)

test:
	go run $(TEST_FILE)
