data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./model",
    "--dialect", "postgres"
  ]
}

data "template_dir" "migrations" {
  path = "migrations"
}

env "dev" {
  src = data.external_schema.gorm.url
  dev = "postgres://postgres:postgres@db:5432/atlas?sslmode=disable"
  url = getenv("BACKEND_DATABASE_URL")

  migration {
    dir = data.template_dir.migrations.url
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

env "prod" {
  src = data.external_schema.gorm.url
  url = getenv("BACKEND_DATABASE_URL")

  migration {
    dir = data.template_dir.migrations.url
    lock_timeout = "5m"
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
