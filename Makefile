.PHONY: migrate-schema run

migrate-schema:
	go run interface/cli/migration/main.go 

run:
	go run main.go
