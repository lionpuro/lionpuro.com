include .env
export

docker-stop:
	DOCKER_HOST=$(DOCKER_HOST) docker compose stop

docker-up:
	export DOCKER_HOST=$(DOCKER_HOST) docker compose pull && docker compose up -d

build-image:
	docker build -t lionpuro/lionpuro.com .
	docker push lionpuro/lionpuro.com

deploy: build-image
	export DOCKER_HOST=$(DOCKER_HOST) docker compose pull && docker compose up -d


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
