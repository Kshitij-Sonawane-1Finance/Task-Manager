generate_migrations:
	export $$(cat .env | xargs) && atlas migrate diff --env gorm --dev-url $${ATLAS_DEV_URL}

apply_migrations:
	export $$(cat .env | xargs) && atlas migrate apply --env gorm --url $${}

migration_hash:
	atlas migrate hash --dir file://db/migrations