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
