resource "cloudflare_ruleset" "webapp_to_subdomain_migration_redirect" {
  zone_id = local.cloudflare_zone_id
  kind    = "zone"
  name    = "Webapp to Subdomain Migration Redirect"
  phase   = "http_request_dynamic_redirect"

  rules = [
    {
      action      = "redirect"
      expression  = "(http.host eq \"${local.domain}\")"
      description = "Redirect root domain to webapp subdomain"

      action_parameters = {
        from_value = {
          status_code           = 301
          preserve_query_string = true
          target_url = {
            value = "https://${local.webapp_domain}"
          }
        }
      }
    }
  ]
}
