resource "cockroach_cluster" "app" {
  name              = local.app_name
  cloud_provider    = "AWS"
  plan              = "BASIC"
  delete_protection = true

  regions = [
    { name = "eu-central-1" }
  ]

  serverless = {
    usage_limits = {
      request_unit_limit = 30000000
      storage_mib_limit  = 10240
    }
  }
}

resource "cockroach_sql_user" "app" {
  cluster_id = cockroach_cluster.app.id
  name       = local.app_name
  password   = data.bitwarden_secret.backend_database_password.value
}

data "cockroach_cluster_cert" "ca_cert" {
  id = cockroach_cluster.app.id
}

data "cockroach_connection_string" "app" {
  id       = cockroach_cluster.app.id
  sql_user = cockroach_sql_user.app.name
  password = cockroach_sql_user.app.password
  database = "defaultdb"
  os       = "LINUX"
}
