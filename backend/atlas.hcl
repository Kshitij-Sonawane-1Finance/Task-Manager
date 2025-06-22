# env "local" {
#   url = "postgres://postgres:kshitij@localhost:5432/golang_db?sslmode=disable"
# }

# data "gorm" "dev" {
#   path = "models"
#   dialect = "postgres"
# }

# migration {
#   dir = "file://db/migrations"
#   format = "sql"
# }


data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./loader",
    "--dialect", "postgres",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  # dev = env("ATLAS_DEV_URL")
  migration {
    dir = "file://db/migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}