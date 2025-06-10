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

data "external" "webapp_deploy" {
  depends_on = [data.external.webapp_build]
  program = ["bash", "${path.module}/deploy/webapp.sh"]

  query = {
    # deploy_id is used to track changes in the webapp source code
    deploy_id = local.webapp_sources_hash

    cloudflare_account_id = local.cloudflare_account_id
    cloudflare_api_token  = local.cloudflare_api_token
    build_dist            = local.webapp_build_dir
    project_name          = cloudflare_pages_project.webapp.name
  }
}
