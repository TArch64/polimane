data "bitwarden_project" "app" {
  id              = "84e18518-5d88-4777-b079-b30601193522"
  organization_id = "d2229b1b-9b38-4b45-a845-b30601180fe5"
}

data "bitwarden_secret" "cloudflare_api_token" {
  key             = "cloudflare_api_token"
  organization_id = data.bitwarden_project.app.organization_id
}

data "bitwarden_secret" "cockroach_api_key" {
  key             = "cockroach_api_key"
  organization_id = data.bitwarden_project.app.organization_id
}

data "bitwarden_secret" "backend_default_user" {
  key             = "backend_default_user"
  organization_id = data.bitwarden_project.app.organization_id
}

data "bitwarden_secret" "backend_default_password" {
  key             = "backend_default_password"
  organization_id = data.bitwarden_project.app.organization_id
}

data "bitwarden_secret" "backend_secret_key" {
  key             = "backend_secret_key"
  organization_id = data.bitwarden_project.app.organization_id
}

data "bitwarden_secret" "backend_sentry_dsn" {
  key             = "backend_sentry_dsn"
  organization_id = data.bitwarden_project.app.organization_id
}

data "bitwarden_secret" "frontend_sentry_dsn" {
  key             = "frontend_sentry_dsn"
  organization_id = data.bitwarden_project.app.organization_id
}

data "bitwarden_secret" "frontend_sentry_auth_token" {
  key             = "frontend_sentry_auth_token"
  organization_id = data.bitwarden_project.app.organization_id
}

data "bitwarden_secret" "backend_database_password" {
  key             = "backend_database_password"
  organization_id = data.bitwarden_project.app.organization_id
}

resource "bitwarden_secret" "backend_database_cert" {
  key             = "backend_database_cert"
  note            = "backend_database_cert"
  project_id      = data.bitwarden_project.app.id
  organization_id = data.bitwarden_project.app.organization_id
  value = sensitive(data.cockroach_cluster_cert.ca_cert.cert)
}

resource "bitwarden_secret" "backend_database_url" {
  key             = "backend_database_url"
  note            = "backend_database_url"
  project_id      = data.bitwarden_project.app.id
  organization_id = data.bitwarden_project.app.organization_id
  value = sensitive("${data.cockroach_connection_string.app.connection_string}&sslrootcert=/tmp/postgres/ca-cert.pem")
}

data "bitwarden_secret" "backend_workos_client_id" {
  key             = "backend_workos_client_id"
  organization_id = data.bitwarden_project.app.organization_id
}

data "bitwarden_secret" "backend_workos_api_key" {
  key             = "backend_workos_api_key"
  organization_id = data.bitwarden_project.app.organization_id
}
