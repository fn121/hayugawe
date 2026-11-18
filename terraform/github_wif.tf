# 1. Aktifkan API yang dibutuhkan
resource "google_project_service" "iam_credentials" {
  project = var.project
  service = "iamcredentials.googleapis.com"
  disable_on_destroy = false
}