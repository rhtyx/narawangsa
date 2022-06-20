commit:
	git add .
	git commit -m "${title}"

push:
	git push origin master

migrate-create:
	migrate create -ext sql -dir db/migration -seq ${name}

.PHONY: commit