MIGRATION_DIR = db/migrations
DB_URL=admin:DevPass_2026!@tcp(127.0.0.1:3306)/auth_dev

# Create a new migration -> gmake migrate-create name="create_entity_table
migrate-create:
	goose -dir $(MIGRATION_DIR) create $(name) sql

migrate-up: # -> gmake migrate-up
	goose -dir $(MIGRATION_DIR) mysql '$(DB_URL)' up

migrate-down: # -> gmake migrate-down
	goose -dir $(MIGRATION_DIR) mysql '$(DB_URL)' down

# Rollback all migrations and reset database # gmake migrate-reset
migrate-reset:
	goose -dir $(MIGRATION_DIR) mysql '$(DB_URL)' reset

# Show current migration status # gmake migrate-status
migrate-status:
	goose -dir $(MIGRATION_DIR) mysql '$(DB_URL)' status

# Redo last migration (Down then Up) # gmake migrate-redo
migrate-redo:
	goose -dir $(MIGRATION_DIR) mysql '$(DB_URL)' redo

# Run specific migration version # gmake migrate-version version=20200101120000
migrate-to:
	goose -dir $(MIGRATION_DIR) mysql '$(DB_URL)' up-to $(version)

# Rollback to a specific migration version # gmake migrate-down-to version=20200101120000
migrate-down-to:
	goose -dir $(MIGRATION_DIR) mysql '$(DB_URL)' down-to $(version)

# Force a specific migration version # gmake migrate-force version=20200101120000
migrate-force:
	goose -dir $(MIGRATION_DIR) mysql '$(DB_URL)' force $(version)

# Print Goose help # gmake migrate-help
migrate-help:
	goose -h