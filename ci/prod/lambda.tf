locals {
  lambda_sources_dir = abspath("${path.root}/../../backend")
  lambda_bin = abspath("${path.root}/tmp/lambda")
}

resource "null_resource" "lambda_build" {
  triggers = {
    sources = sha1(join("", [
      for f in fileset(local.lambda_sources_dir, "*") : filesha1("${local.lambda_sources_dir}/${f}")
    ]))
  }

  provisioner "local-exec" {
    command     = "make out_file=\"${local.lambda_bin}\" prod"
    working_dir = local.lambda_sources_dir
  }
}