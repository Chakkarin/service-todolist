.PHONY: setup
setup: 
	@ docker compose -f docker-compose.db.yml up -d
	@ go run database/migration/migration.go

.PHONY: reset
reset:
	@ docker compose -f docker-compose.db.yml down
	@ docker compose -f docker-compose.db.yml up -d

.PHONY: run
run:
	@ go run main.go

.PHONY: apidoc
apidoc:
	@ swag init