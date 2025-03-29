include .env
export

docker-build:
	DOCKER_HOST=$(DOCKER_HOST) docker compose build

docker-stop:
	DOCKER_HOST=$(DOCKER_HOST) docker compose stop

docker-up:
	DOCKER_HOST=$(DOCKER_HOST) docker compose up -d

docker-deploy:
	export DOCKER_HOST=$(DOCKER_HOST) docker compose build && docker compose up -d --build

build:
	@npx @tailwindcss/cli -i ./input.css -o ./static/global.css --minify
	@templ generate views
	@go build -o bin/app .

run: build
	@./bin/app

fmt:
	@templ fmt views
	@gofmt -l -s -w .

dev:
	@air -c .air.toml
