#!/bin/bash

cleanup() {
    echo "Cleaning up..."
    exit
}

trap cleanup EXIT

go tool wgo -debounce 100ms -xfile static/global.css \
	npx @tailwindcss/cli -i ./input.css -o ./static/global.css --minify \
	:: wgo -xfile _templ.go -xfile static/global.css templ generate views :: go run .
