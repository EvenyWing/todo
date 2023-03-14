PATHGO = ./cmd/main.go

run:
	go run $(PATHGO)

migrate:
	go run $(PATHGO) -migrate