resource "cloudflare_pages_project" "webapp" {
  account_id        = local.cloudflare_account_id
  name              = "polimane-prod-admin"
  production_branch = "main"

  lifecycle {
    # fix endless update of the project
    ignore_changes = [build_config, deployment_configs]
  }
}

resource "cloudflare_pages_domain" "webapp" {
  name         = local.domain
  account_id   = local.cloudflare_account_id
  project_name = cloudflare_pages_project.webapp.name
}

resource "null_resource" "webapp_deploy" {
  triggers = { sources_hash = local.webapp_sources_hash }

  depends_on = [
    null_resource.webapp_build,
    cloudflare_pages_project.webapp
  ]

  provisioner "local-exec" {
    command = "bash ${path.module}/job/run.sh"

    environment = {
      JOB_ID        = local.webapp_sources_hash
      BUILD_IMAGE   = "polimane-prod-frontend-deploy"
      BUILD_DOCKERFILE = abspath("${path.root}/job/frontend.Dockerfile")
      BUILD_CONTEXT = local.webapp_sources_dir

      BUILD_SECRET = jsonencode(["CLOUDFLARE_API_TOKEN"])
      CLOUDFLARE_API_TOKEN = data.bitwarden_secret.cloudflare_api_token.value

      BUILD_ARGS = jsonencode(["CLOUDFLARE_ACCOUNT_ID", "PROJECT_NAME"])
      CLOUDFLARE_ACCOUNT_ID = local.cloudflare_account_id
      PROJECT_NAME          = cloudflare_pages_project.webapp.name
    }
  }
}
