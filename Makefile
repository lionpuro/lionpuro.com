include .env
export

build-image:
	docker build -t lionpuro/lionpuro.com .
	docker push lionpuro/lionpuro.com

deploy:
	export DOCKER_HOST=$(DOCKER_HOST)
	docker compose -f ./infra/compose.yaml pull
	docker compose -f ./infra/compose.yaml up --build -d

build:
	@npx @tailwindcss/cli -i ./input.css -o ./static/global.css --minify
	@templ generate views
	@go build -o bin/app .

fmt:
	@templ fmt views
	@gofmt -l -s -w .

run-dev:
	@./run_dev.sh
