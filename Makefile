run:
	npm run build --prefix app/
	go build
	go run todoApp

build:
	npm run build --prefix app/
	go build

dockerimg:
	make build
	docker build -t mgkeiji/todoapp .