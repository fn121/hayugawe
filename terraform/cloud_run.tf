resource "google_cloud_run_v2_service" "task" {
  name = "${var.service_name}-service"
  location = var.region
  ingress = "INGRESS_TRAFFIC_ALL"

  scaling {
    min_instance_count = 0
  }

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