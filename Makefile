test:
	go test -v -cover ./...

run:
	docker-compose up -d

cleanup:
	docker-compose down
	@ docker volume rm $$(docker volume ls -q)

server: 
	go run main.go

.PHONY: test run server cleanup