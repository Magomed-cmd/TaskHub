run:
	go run ./cmd/taskhub/
build:
	go build cmd/taskhub/main.go
air:
	air -c .air.toml