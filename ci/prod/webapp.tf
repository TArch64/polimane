locals {
  webapp_domain         = "app.${local.domain}"
  webapp_release_cookie = "webapp-release"
}

resource "cloudflare_worker" "webapp" {
  account_id = local.cloudflare_account_id
  name       = "${local.app_name}-webapp"
}

resource "cloudflare_worker_version" "webapp" {
  depends_on         = [null_resource.webapp_build]
  account_id         = local.cloudflare_account_id
  worker_id          = cloudflare_worker.webapp.id
  compatibility_date = "2025-01-01"
  main_module        = "worker.js"

  assets = {
    directory = "${local.webapp_build_dir}/public"

    config = {
      not_found_handling = "single-page-application",
      run_worker_first   = ["/*"]
    }
  }

  bindings = [
    {
      name = "ASSETS"
      type = "assets"
    }
  ]

  modules = [
    {
      name         = "worker.js"
      content_file = "${local.webapp_build_dir}/worker.js"
      content_type = "application/javascript+module"
    }
  ]

  placement = {
    mode = "smart"
  }
}

resource "cloudflare_workers_deployment" "webapp" {
  account_id  = local.cloudflare_account_id
  script_name = cloudflare_worker.webapp.name
  strategy    = "percentage"

  versions = [
    {
      percentage = 100,
      version_id = cloudflare_worker_version.webapp.id
    }
  ]
}

resource "cloudflare_workers_custom_domain" "webapp" {
  depends_on  = [null_resource.webapp_build]
  account_id  = local.cloudflare_account_id
  zone_id     = local.cloudflare_zone_id
  hostname    = local.webapp_domain
  service     = cloudflare_worker.webapp.name
  environment = "production"
}

resource "cloudflare_ruleset" "set_webapp_worker_version" {
  zone_id = local.cloudflare_zone_id
  kind    = "zone"
  name    = "Set Webapp Worker Version"
  phase   = "http_request_late_transform"

  rules = [
    {
      action     = "rewrite"
      expression = "http.cookie contains \"${local.webapp_release_cookie}\""

      action_parameters = {
        headers = {
          "Cloudflare-Workers-Version-Key" = {
            operation = "set"
            value     = "http.request.cookies[\"${local.webapp_release_cookie}\"][0]"
          }
        }
      }
    }
  ]
}

resource "cloudflare_ruleset" "webapp_cache" {
  zone_id = local.cloudflare_zone_id
  kind    = "zone"
  name    = "Webapp Cache"
  phase   = "http_response_headers_transform"

  rules = [
    {
      action     = "rewrite"
      expression = "starts_with(http.request.uri.path, \"/assets/\")"

      action_parameters = {
        headers = {
          "Cache-Control" = {
            operation = "set"
            value     = "public, max-age=31536000, immutable"
          }
        }
      }
    }
  ]
}
