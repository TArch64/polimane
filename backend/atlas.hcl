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

variable "default_user" {
  type    = string
  default = getenv("BACKEND_DEFAULT_USER")
}

variable "default_password" {
  type    = string
  default = getenv("BACKEND_DEFAULT_PASSWORD")
}

data "template_dir" "migrations" {
  path = "migrations"
  vars = {
    default_user = var.default_user
    default_password = var.default_password
  }
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
