#!/bin/bash

COMPOSE_FILE="compose.dev.yaml"

# Helpers

usage() {
	echo "Usage: $0 [COMMAND] [ARGUMENTS]"
	echo "Commands:"
	echo "  start     Start up the development containers"
	echo "  stop      Shut down the development containers"
	echo "  cmd       Run a command in the workspace container"
	echo "  shell     Open a shell into the workspace container"
	echo "  fmt       Format templ and go files"
}

fn_exists() {
    type $1 2>/dev/null | grep -q 'is a function'
}

COMMAND=$1
shift
ARGUMENTS=${@}

# Commands

start() {
	docker compose -f $COMPOSE_FILE up
}

stop() {
	docker compose -f $COMPOSE_FILE stop
}

cmd() {
	docker compose -f $COMPOSE_FILE exec workspace ${@}
}

shell() {
	docker compose -f $COMPOSE_FILE exec workspace bash
}

fmt() {
	docker compose -f $COMPOSE_FILE exec workspace sh -c \
		"go tool templ fmt /app/views && gofmt -l -s -w /app"
}

# Exec

fn_exists $COMMAND
if [ $? -eq 0 ]; then
	$COMMAND $ARGUMENTS
else
	usage
fi
