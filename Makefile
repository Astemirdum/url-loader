
.PHONY: run
run:
	go run cmd/main.go file test/test.txt -g 3

.PHONY: build
build:
	go build -o bin/url-loader cmd/main.go

.PHONY: build-d
build-d:
	docker compose up -d --build

.PHONY: run-d
run-d:
	docker exec -it url-loader url-loader file test/test.txt -g 10

.PHONY: lint
lint:
	go vet ./...
	golangci-lint run --fix

.PHONY: test
test:
	go test -v -race -timeout 90s -count=1 -shuffle=on  -coverprofile cover.out ./...
	@go tool cover -func cover.out | grep total | awk '{print $3}'
	go tool cover -html="cover.out" -o coverage.html

.PHONY: docker-login
docker-login:
	docker login -u ${REGISTRY_USER} -p ${REGISTRY_PASS}

.PHONY: push
push:
	docker compose push
