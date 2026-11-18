locals {
  service_account_id = "github-actions-sa"
  pool_id = "github_pool"
  provider_id         = "github-provider"
  github_repo         = var.github_repo
}

# 1. Aktifkan API yang dibutuhkan
resource "google_project_service" "iam" {
  project = var.project
  service = "iamcredentials.googleapis.com"
  disable_on_destroy = false
}

# Service account for github actions
resource "google_service_account" "github_sa" {
  account_id = local.service_account_id
  display_name = "GitHub Actions Cloud Run Deployer"
}

# set roles for gh actions sa
resource "google_project_iam_member" "github_actions" {
  for_each = toset([
      "roles/artifactregistry.writer",
    "roles/run.developer"
  ])

  project = var.project
  role = each.key
  member = "serviceAccount:${google_service_account.github_sa.email}"
}

# Allow GitHub repository to impersonate the Service Account
resource "google_service_account_iam_member" "wif_bind" {
  service_account_id = google_service_account.github_sa.name
  role               = "roles/iam.workloadIdentityUser"
  member             = "principalSet://iam.googleapis.com/${google_iam_workload_identity_pool.github_pool.name}/attribute.repository/${var.github_repo}"
}

# Workload Identity Pool
resource "google_iam_workload_identity_pool" "github_pool" {
  workload_identity_pool_id = "github-pool"
  display_name              = "GitHub Actions Pool"
  depends_on                = [google_project_service.iam]
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
  # dan hanya dari main
  attribute_condition = <<-EOT
  assertion.repository == '${var.github_repo}' 
EOT
  
  oidc {
    issuer_uri = "https://token.actions.githubusercontent.com"
  }
}


# 6. Output nilai yang dibutuhkan untuk file YAML GitHub Actions
output "workload_identity_provider" {
  value = google_iam_workload_identity_pool_provider.github_provider.name
}

output "service_account_email" {
  value = google_service_account.github_sa.email
}