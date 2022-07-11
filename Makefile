include .env
ifeq ($(Environment),development)
DOCKER_COMMAND=docker-compose
else
DOCKER_COMMAND=docker-compose -f docker-compose.prod.yml
endif
MIGRATE=${DOCKER_COMMAND} exec web migrate -path=core/migrations -database "postgres://${DBUsername}:${DBPassword}@${DBHost}:${DBPort}/${DBName}?sslmode=disable" -verbose

migrate-up:
		$(MIGRATE) up
migrate-down:
		$(MIGRATE) down
force:
		@read -p  "Which version do you want to force?" VERSION; \
		$(MIGRATE) force $$VERSION

goto:
		@read -p  "Which version do you want to migrate?" VERSION; \
		$(MIGRATE) goto $$VERSION

drop:
		$(MIGRATE) drop

create-migration:
		@read -p  "What is the name of migration?" NAME; \
		${MIGRATE} create -ext sql -seq -dir migration  $$NAME

.PHONY: migrate-up migrate-down force goto drop create

.PHONY: migrate-up migrate-down force goto drop create auto-create
