data "external_schema" "gorm" {
  program = ["go", "run", "-mod=mod", "./migrations"]
}

data "template_dir" "migrations" {
  path = "migrations"
}

env "dev" {
  src = data.external_schema.gorm.url
  dev = "postgresql://root@db:26257/atlas?sslmode=disable"
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
