locals {
  lambda_sources_dir = abspath("${path.root}/../../backend")
  lambda_build_dir = abspath("${path.root}/tmp/lambda")
  lambda_build_zip = abspath("${local.lambda_build_dir}/bootstrap.zip")

  lambda_sources_hash = sha1(join("", [
    for f in fileset(local.lambda_sources_dir, "**") : filesha1("${local.lambda_sources_dir}/${f}")
  ]))
}

resource "null_resource" "lambda_build" {
  triggers = { sources_hash = local.lambda_sources_hash }

  provisioner "local-exec" {
    command = "bash ${path.module}/build/build.sh"

    environment = {
      BUILD_ID      = local.lambda_sources_hash
      BUILD_IMAGE   = "polimane-prod-backend"
      BUILD_DOCKERFILE = abspath("${path.root}/build/backend.Dockerfile")
      BUILD_CONTEXT = local.lambda_sources_dir
      BUILD_DIST    = local.lambda_build_dir
    }
  }
}
