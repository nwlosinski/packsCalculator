APP_NAME = packscalculator
BIN_NAME = server
PKG = github.com/nwlosinski/packsCalculator

## Run the app locally
run:
	go run main.go

## Build local binary
build:
	go build -o $(BIN_NAME) main.go

## Run tests
test:
	go test ./...

## Clean build artifacts
clean:
	rm -f $(BIN_NAME)

## Build Docker image
docker-build:
	docker build -t $(APP_NAME):latest .

## Run Docker container
docker-run:
	docker run -p 80:80 $(APP_NAME):latest

## Up docker compose
up:
	docker compose up --build

## Down docker compose
down:
	docker compose down

