build:
	docker compose build --no-cache

up:
	docker compose up

postgres-container:
	set -a && source .env && set +a && docker exec -it $${POSTGRES_CONTAINER} bash

migrate:
	psqldef -U postgres -p 5432 english_vocabulary </var/lib/postgresql/schema/schema.sql