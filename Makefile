#!bin/bash
export DOMAIN_ADMIN=localhost:5010

run-admin:
	go run cmd/admin/main.go

swagger-admin:
	swag init -d ./ -g cmd/admin/main.go \
		-o ./docs/admin