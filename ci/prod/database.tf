data "aws_dynamodb_table" "database" {
  depends_on = [aws_lambda_function.lambda]
  name = local.app_name
}

resource "aws_backup_vault" "vault" {
  name = local.app_name
  tags = local.aws_common_tags
}

resource "aws_backup_plan" "vault" {
  name = aws_backup_vault.vault.name
  tags = local.aws_common_tags

  rule {
    rule_name         = "every-day"
    target_vault_name = aws_backup_vault.vault.name
    schedule          = "cron(0 4 * * ? *)"

    lifecycle {
      delete_after = 3
    }
  }
}

data "aws_iam_policy_document" "backup" {
  statement {
    effect = "Allow"

    principals {
      type = "Service"
      identifiers = ["backup.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "backup_role" {
  name               = "polimane-prod-backup-role"
  assume_role_policy = data.aws_iam_policy_document.backup.json
}

resource "aws_iam_role_policy_attachment" "database_backup_role" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSBackupServiceRolePolicyForBackup"
  role       = aws_iam_role.backup_role.name
}

resource "aws_backup_selection" "database" {
  name         = "database"
  iam_role_arn = aws_iam_role.backup_role.arn
  plan_id      = aws_backup_plan.vault.id
  resources = [data.aws_dynamodb_table.database.arn]

  lifecycle {
    ignore_changes = [resources]
  }
}
