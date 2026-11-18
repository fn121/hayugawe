locals {
  service_account_id = "github-actions-sa"
  pool_id = "github_pool"
  provider_id         = "github-provider"
  github_repo         = var.github_repo
}

# 1. Aktifkan API yang dibutuhkan
resource "google_project_service" "iam_api" {
  project = var.project
  service = "iamcredentials.googleapis.com"
  disable_on_destroy = false
}

resource "google_service_account" "github_sa" {
  account_id = local.service_account_id
  display_name = "GitHub Actions Cloud Run Deployer"
}

# Workload Identity Pool
resource "google_iam_workload_identity_pool" "github_pool" {
  workload_identity_pool_id = "github-pool"
  display_name              = "GitHub Actions Pool"
  depends_on                = [google_project_service.iam_api]
}

# Workload Identity Provider
resource "google_iam_workload_identity_pool_provider" "github_provider" {
  project                            = var.project
  workload_identity_pool_id          = google_iam_workload_identity_pool.github_pool.workload_identity_pool_id
  workload_identity_pool_provider_id = local.provider_id
  display_name                       = "GitHub OIDC Provider"

  # Mapping atribut dari token GitHub ke GCP
  attribute_mapping = {
    "google.subject"       = "assertion.sub"
    "attribute.actor"      = "assertion.actor"
    "attribute.repository" = "assertion.repository"
  }

  # Keamanan Ekstra: Hanya izinkan request dari repositori spesifik Anda
  attribute_condition = "assertion.repository == '${local.github_repo}'"

  oidc {
    issuer_uri = "https://githubusercontent.com"
  }
}

# 6. Output nilai yang dibutuhkan untuk file YAML GitHub Actions
output "workload_identity_provider" {
  value = google_iam_workload_identity_pool_provider.github_provider.name
}

output "service_account_email" {
  value = google_service_account.github_sa.email
}