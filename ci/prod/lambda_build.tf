locals {
  lambda_sources_dir = abspath("${path.root}/../../backend")
  lambda_dist_dir = abspath("${path.root}/tmp/lambda")
  lambda_zip = "${local.lambda_dist_dir}/lambda.zip"

  lambda_sources_hash = sha1(join("", [
    for f in fileset(local.lambda_sources_dir, "**") : filesha1("${local.lambda_sources_dir}/${f}")
  ]))
}

resource "null_resource" "lambda_build" {
  triggers = { sources = local.lambda_sources_hash }

  provisioner "local-exec" {
    command     = "make out_dir=\"${local.lambda_dist_dir}\" prod"
    working_dir = local.lambda_sources_dir
  }
}