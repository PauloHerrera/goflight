test:
	go test -v -cover ./...

setup-db:
	docker-compose up -d

cleanup:
	docker-compose down
	@ docker volume rm $$(docker volume ls -q)

server: 
	go run main.go

run: setup-db server

.PHONY: test run server setup-db