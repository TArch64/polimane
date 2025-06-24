data "bitwarden_secret" "frontend_sentry_dsn" {
  key = "frontend_sentry_dsn"
}

data "bitwarden_secret" "frontend_sentry_auth_token" {
  key = "frontend_sentry_auth_token"
}

locals {
  webapp_sources_dir = abspath("${path.root}/../../frontend")
  webapp_build_dir = abspath("${path.root}/tmp/webapp")

  webapp_sources_hash = sha1(join("", [
    for f in fileset(local.webapp_sources_dir, "**") : filesha1("${local.webapp_sources_dir}/${f}")
  ]))
}

resource "null_resource" "webapp_build" {
  triggers = { sources_hash = local.webapp_sources_hash }

  provisioner "local-exec" {
    command = "bash ${path.module}/build/build.sh"

    environment = {
      BUILD_ID      = local.webapp_sources_hash
      BUILD_IMAGE   = "polimane-prod-frontend"
      BUILD_DOCKERFILE = abspath("${path.root}/build/frontend.Dockerfile")
      BUILD_CONTEXT = local.webapp_sources_dir
      BUILD_DIST    = local.webapp_build_dir

      BUILD_SECRET = jsonencode(["FRONTEND_PUBLIC_SENTRY_DSN", "SENTRY_AUTH_TOKEN"])
      FRONTEND_PUBLIC_SENTRY_DSN = data.bitwarden_secret.frontend_sentry_dsn.value
      SENTRY_AUTH_TOKEN          = data.bitwarden_secret.frontend_sentry_auth_token.value


      BUILD_ARGS = jsonencode(["FRONTEND_PUBLIC_API_URL", "FRONTEND_PUBLIC_SENTRY_RELEASE"])
      FRONTEND_PUBLIC_API_URL        = "https://${local.api_domain}/api",
      FRONTEND_PUBLIC_SENTRY_RELEASE = local.webapp_sources_hash
    }
  }
}
