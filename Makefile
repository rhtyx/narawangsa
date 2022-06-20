commit:
	git add .
	git commit -m "${title}"

push:
	git push origin master

migrate-create:
	migrate create -ext sql -dir db/migration -seq ${name}

migrate-up:
	migrate -path db/migration -database "postgresql://narawangsa:narawangsa@localhost:5434/narawangsa_db?sslmode=disable" up

migrate-down:
	migrate -path db/migration -database "postgresql://narawangsa:narawangsa@localhost:5434/narawangsa_db?sslmode=disable" down

.PHONY: commit