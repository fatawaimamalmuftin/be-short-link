include .env

SEEDER_PATH=db/seeding
MIGRATE_PATH=db/migrations
DB_URL=postgres://$(DF_DB_USER):$(DF_DB_PASS)@$(DF_DB_HOST):$(DF_DB_PORT_H)/$(DF_DB_NAME)?sslmode=disable

migrate-create:
	@read -p "Enter migration name: " NAME; \
	migrate create -ext sql -dir $(MIGRATE_PATH) -seq $$NAME

migrate-up:
	migrate -path $(MIGRATE_PATH) -database "$(DB_URL)" -verbose up
	
migrate-down-all:
	migrate -path $(MIGRATE_PATH) -database "$(DB_URL)" -verbose drop -f

migrate-down:
	migrate -path $(MIGRATE_PATH) -database "$(DB_URL)" -verbose down 1

migrate-force:
	@read -p "Enter version: " VERSION; \
	migrate -path $(MIGRATE_PATH) -database "$(DB_URL)" force $$VERSION

seed-create:
	mkdir -p $(SEEDER_PATH)
	touch $(SEEDER_PATH)/$(shell printf "%06d" $$(find $(SEEDER_PATH) -name "*.sql" | wc -l | awk '{print $$1+1}'))_$(n).sql

seed-up:
	go run db/seeding/main.go

seed-all:
	@for file in $(SEEDER_PATH)/*.sql; do \
		echo "Seeding $$file..."; \
		psql "$(DB_URL)" < "$$file" || exit 1; \
	done

up:
	docker compose up -d --build
	make logs

down:
	docker compose down --rmi local

up:
	docker compose up -d --build
	make logs

down:
	docker compose down --rmi local

reset:
	make down
	make up

ps:
	docker compose ps -a

logs:
	docker logs -f sl-be

swag:
	swag init -g cmd/main.go -o docs

reset:
	make down
	make up
	make logs