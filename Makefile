DB_HOST ?= $(POSTGRES_HOST)
DB_PORT ?= $(POSTGRES_PORT)
DB_NAME ?= $(POSTGRES_DB)
DB_USER ?= $(POSTGRES_USER)
DB_PASS ?= $(POSTGRES_PASSWORD)

DATABASE_URL := "postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"

create:
	createdb banking

drop:
	dropdb banking

mup:
	migrate -path db/migration -database $(DATABASE_URL) -verbose up

mdown:
	migrate -path db/migration -database $(DATABASE_URL) -verbose down

sqlc:
	sqlc generate

.PHONY: create drop mup mdown sqlc