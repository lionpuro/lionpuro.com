#!/bin/bash

cleanup() {
    echo "Cleaning up..."
    exit
}

trap cleanup EXIT

go tool wgo -debounce 100ms -xfile _templ.go -xdir assets/public \
	npm run build \
	:: wgo -xfile _templ.go -xdir assets/public templ generate views \
	:: go run .
