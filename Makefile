.PHONY: run
run:
	@export TELEGRAM_API_TOKEN=./internal/secret/.secret && go run cmd/*.go

.PHONY: build
build:
	@go build -o ./app cmd/*.go

.PHONY: test
test:
	go test -v ./...

.PHONY: docker-build
docker-build:
	@docker build -t tech-bot .

.PHONY: docker-run
docker-run:
	@docker run \
		--name tech-bot \
		-d \
		--rm \
		-p 80:8080 \
		-v `pwd`./internal/secret/.secret \
		-e TELEGRAM_API_TOKEN=./internal/secret/.secret \
		tech-bot

.PHONY: docker-stop
docker-stop:
	@docker stop tech-bot