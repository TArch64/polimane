locals {
  lambda_sources_dir = abspath("${path.root}/../../backend")
  lambda_build_dir = abspath("${path.root}/tmp/lambda")
  lambda_build_zip = abspath("${local.lambda_build_dir}/boostrap.zip")

  lambda_sources_hash = sha1(join("", [
    for f in fileset(local.lambda_sources_dir, "**") : filesha1("${local.lambda_sources_dir}/${f}")
  ]))
}

data "external" "lambda_build" {
  program = ["bash", "${path.module}/build/build.sh"]

  query = {
    build_id      = local.lambda_sources_hash
    build_image   = "polimane-prod-backend"
    build_dockerfile = abspath("${path.root}/build/backend.Dockerfile")
    build_context = local.lambda_sources_dir
    build_dist    = local.lambda_build_dir
  }
}
