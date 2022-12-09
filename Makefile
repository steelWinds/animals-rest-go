run-tests:
	cd app; \
    	go test -v ./...

run-prod-app:
	docker compose -f docker/docker-compose.yml -f docker/docker-compose.prod.yml up -d

run-dev-app:
	docker compose -f docker/docker-compose.yml -f docker/docker-compose.dev.yml up -d