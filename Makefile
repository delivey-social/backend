.PHONY: docker-build docker-run docker-build-dev docker-run-dev test coverage

DEV_CONTAINER=comida-app-backend-container-dev
DEV_NAME=comida-app-backend-dev

PROD_CONTAINER=comida-app-backend-container
PROD_NAME=comida-app-backend
PORT=3001

docker-build:
	docker build -t $(PROD_NAME) . -f Dockerfile

docker-run:
	docker run -it --rm -p $(PORT):$(PORT)  --name $(PROD_CONTAINER) $(PROD_NAME)

docker-build-dev:
	docker build -t $(DEV_NAME) . -f Dockerfile.dev

docker-run-dev:
	docker run -it --rm -p $(PORT):$(PORT)  --name $(DEV_CONTAINER) $(DEV_NAME)

test:
	go test ./...

coverage:
	go test -coverprofile cover.prof ./...
	covreport -o cover.html
	cmd.exe /c start cover.html
