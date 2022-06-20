commit:
	git add .
	git commit -m "${title}"

push:
	git push origin master

postgres:
	docker run --name postgres14_narawangsa -p 5434:5432 -e POSTGRES_USER=narawangsa -e POSTGRES_PASSWORD=narawangsa postgres:14-alpine

createdb:
	docker exec -it postgres14_narawangsa createdb --username=narawangsa --owner=narawangsa narawangsa_db

dropdb:
	docker exec -it postgres14_narawangsa dropdb narawangsa_db

migrate-create:
	migrate create -ext sql -dir db/migration -seq ${name}

migrate-up:
	migrate -path db/migration -database "postgresql://narawangsa:narawangsa@localhost:5434/narawangsa_db?sslmode=disable" up

migrate-down:
	migrate -path db/migration -database "postgresql://narawangsa:narawangsa@localhost:5434/narawangsa_db?sslmode=disable" down

.PHONY: commit