locals {
  webapp_sources_dir = abspath("${path.root}/../../frontend")
  webapp_build_dir = abspath("${path.root}/tmp/webapp")

  webapp_sources_hash = sha1(join("", [
    for f in fileset(local.webapp_sources_dir, "**") : filesha1("${local.webapp_sources_dir}/${f}")
  ]))
}

# todo refactor to use null_resource with triggers
data "external" "webapp_build" {
  program = ["bash", "${path.module}/build/build.sh"]

  query = {
    build_id      = local.webapp_sources_hash
    build_image   = "polimane-prod-frontend"
    build_dockerfile = abspath("${path.root}/build/frontend.Dockerfile")
    build_context = local.webapp_sources_dir
    build_dist    = local.webapp_build_dir
    build_args = join("|", [
      "API_URL=https://${local.api_domain}/api"
    ])
  }
}
