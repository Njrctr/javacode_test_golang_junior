build:
	docker-compose build wallet-backend

run:
	docker-compose up wallet-backend --force-recreate

test:
	go test -v ./...

migrate_up:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' up

migrate_down:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' down

swag:
	swag init -g cmd/main.go