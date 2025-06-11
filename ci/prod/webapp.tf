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
  depends_on = [null_resource.webapp_build]

  provisioner "local-exec" {
    command = "npx -y wrangler pages deploy $BUILD_DIST --project-name $PROJECT_NAME"

    environment = {
      CLOUDFLARE_ACCOUNT_ID = local.cloudflare_account_id
      CLOUDFLARE_API_TOKEN  = local.cloudflare_api_token
      BUILD_DIST   = local.webapp_build_dir
      PROJECT_NAME = cloudflare_pages_project.webapp.name
    }
  }
}
