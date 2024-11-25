fake_api:
	go build -o ./bin/fake ./cmd/fake_api/main.go

run_fake: fake_api
	./bin/fake

build:
	go build -o ./bin/app ./cmd/music_lib/main.go

run: build
	./bin/app

up:
	docker compose -f ./docker-compose.yml up --build

down:
	docker compose down

swag:
	swag init -g cmd/music_lib/main.go

clean_db:
	docker-compose down --volumes
