locals {
  worker_sources_dir  = local.lambda_sources_dir
  worker_sources_hash = local.lambda_sources_hash
  worker_build_dir = abspath("${path.root}/tmp/worker")
  worker_build_zip = abspath("${local.worker_build_dir}/bootstrap.zip")
}

resource "null_resource" "worker_build" {
  triggers = { sources_hash = local.worker_sources_hash }

  provisioner "local-exec" {
    command = "bash ${path.module}/build/build.sh"

    environment = {
      BUILD_ID      = local.worker_sources_hash
      BUILD_IMAGE   = "polimane-prod-worker"
      BUILD_DOCKERFILE = abspath("${path.root}/build/backend.Dockerfile")
      BUILD_CONTEXT = local.worker_sources_dir
      BUILD_DIST    = local.worker_build_dir
      BUILD_TARGET  = "worker"
    }
  }
}
