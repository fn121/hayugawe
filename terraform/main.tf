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

resource "google_cloud_run_v2_service" "task" {
  name = "${var.service_name}-service"
  location = var.region
  ingress = "INGRESS_TRAFFIC_ALL"

  template {
    containers {
      image = "us-docker.pkg.dev/cloudrun/container/hello"

      resources {
        limits = {
          cpu    = "1"
          memory = "512Mi"
        }
      }
    }
  }

  # PENTING: Abaikan perubahan image agar tidak bentrok dengan deploy manual / CI/CD
  # important! ignore changes so hello run only once in first init
  lifecycle {
    ignore_changes = [
      template[0].containers[0].image,
      client,
      client_version,
    ]
  }

  depends_on = [ google_project_service.task_api_run ]
}

resource "google_cloud_run_v2_service_iam_binding" "task" {
  name = google_cloud_run_v2_service.task.name
  location = google_cloud_run_v2_service.task.location
  role = "roles/run.invoker"
  members = [ "allUsers" ]
}