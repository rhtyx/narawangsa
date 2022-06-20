commit:
	git add .
	git commit -m "${title}"

push:
	git push origin master

.PHONY: commit