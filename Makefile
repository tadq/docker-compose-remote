
-include .env
export

setup:
	docker compose -f docker-compose-dev.yaml up

teardown:
	docker compose -f docker-compose-dev.yaml down

build:
	go build

run:
	go run .

deploy-dev: build
	docker --host "ssh://root@lts" compose up --build --force-recreate

deploy: build
	docker --host "ssh://root@lts" compose up -d --build --force-recreate

stop:
	docker --host "ssh://root@lts" compose down
