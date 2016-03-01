all: build_and_start
build_and_start: build docker up
build:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o datafetch
docker:
	docker-compose build datafetch
up:
	docker-compose up datafetch
