docker-build:
	docker build --no-cache -t osscameroon/sammy:latest -f Dockerfile .

run:
	go run main.go $(ARGS)

build:
	go build -o sammy

exec: build
	./sammy $(ARGS)
