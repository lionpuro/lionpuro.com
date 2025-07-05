include .env
export

.PHONY: build-image deploy build

build-image:
	docker build -t lionpuro/lionpuro.com .
	docker push lionpuro/lionpuro.com

deploy:
	export DOCKER_HOST=$(DOCKER_HOST)
	docker compose -f ./infra/compose.yaml pull
	docker compose -f ./infra/compose.yaml up --build -d

build:
	@npm run build
	@go tool templ generate views
	@go build -o bin/app .
