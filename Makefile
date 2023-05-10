-include .env

.PHONY: all
all: setup build run

.PHONY: setup
setup:
	@echo "Loading environment variables..."
	@set -o allexport; source .env; set +o allexport

.PHONY: build
build:
	@echo "Running build..."
	docker-compose build

.PHONY: run
run:
	@echo "Running Run..."
	docker-compose up
