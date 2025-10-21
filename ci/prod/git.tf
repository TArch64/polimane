data "external" "git_commit" {
  program = ["git", "log", "--pretty=format:{ \"sha\": \"%H\" }", "-1", "HEAD"]
}

locals {
  git_commit_sha = data.external.git_commit.result.sha
}
