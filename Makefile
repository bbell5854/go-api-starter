default: build clean run

test:
	go test -v -cover ./...

build:
	docker build -t go-api-starter .

clean:
	docker system prune -f
	docker volume prune -f

run:
	docker-compose up
