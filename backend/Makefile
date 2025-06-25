generate_migrations:
	@. ./.env.sh && \
	export CGO_ENABLED=0 && \
	atlas migrate diff --env gorm --dev-url $$ATLAS_DEV_URL

apply_migrations:
	@. ./.env.sh && \
	atlas migrate apply --env gorm --url $$DB_URL

migration_hash:
	atlas migrate hash --dir file://db/migrations