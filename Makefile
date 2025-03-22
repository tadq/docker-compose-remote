
build:
	go build

run:
	go run .

deploy: build
	docker --host "ssh://root@lts" compose up --build --force-recreate 
	

