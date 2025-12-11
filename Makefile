# Autoload .env
ifneq (,$(wildcard .env))
include .env
export $(shell sed 's/=.*//' .env)
endif

run:
	go run ./cmd/auth-service

build:
	go build -o auth-service ./cmd/auth-service