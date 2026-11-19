terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "6.8.0"
    }
  }
}

provider "google" {
  project = var.project
  region  = var.region
  zone = var.zone
}

resource "google_project_service" "task_api" {
  project = var.project
  service = "artifactregistry.googleapis.com"

  depends_on = [ google_billing_project_info.project_billing ]
}

resource "google_billing_project_info" "project_billing" {
  project = var.project
  billing_account = var.billing_account
}

# @todo update artifact for whole project instead per service
resource "google_artifact_registry_repository" "task_repo" {
  format = "DOCKER"
  repository_id = "${var.service_name}-repo"
  description = "Docker repository untuk ${var.service_name} image"
  location = var.region

  depends_on = [ google_project_service.task_api ]
}


resource "google_project_service" "task_api_run" {
  project = var.project
  service = "run.googleapis.com"

  depends_on = [ google_billing_project_info.project_billing ]
}
