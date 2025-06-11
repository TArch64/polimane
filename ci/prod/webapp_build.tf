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

      BUILD_ARGS = jsonencode({
        API_URL = "https://${local.api_domain}/api"
      })
    }
  }
}
