resource "google_cloud_run_service" "main1" {
  name     = "main1"
  location = "asia-northeast1"

  metadata {
    labels = {
      "cloud.googleapis.com/location" = "asia-northeast1"
    }
    annotations = {
      "run.googleapis.com/client-name"    = "cloud-console"
      "serving.knative.dev/creator"       = "shion1305@a.shion.pro"
      "serving.knative.dev/lastModifier"  = "shion1305@a.shion.pro"
      "run.googleapis.com/operation-id"   = "c4484fff-7201-42fa-b943-c93af66e669c"
      "run.googleapis.com/ingress"        = "all"
      "run.googleapis.com/ingress-status" = "all"
      "run.googleapis.com/minScale"       = "0"
    }
  }

  template {
    metadata {
      labels = {
        "client.knative.dev/nonce"            = "5754c3e0-ad0a-4364-9d17-45ff94238904"
        "run.googleapis.com/startupProbeType" = "Default"
      }
      annotations = {
        "run.googleapis.com/client-name"       = "cloud-console"
        "autoscaling.knative.dev/maxScale"     = "4"
        "run.googleapis.com/startup-cpu-boost" = "true"
      }
    }

    spec {
      container_concurrency = 80
      timeout_seconds       = 300
      service_account_name  = "435876017528-compute@developer.gserviceaccount.com"

      containers {
        name  = "nginx-1"
        image = "asia-northeast1-docker.pkg.dev/kyusyu-hackathon/main/nginx:latest"
        ports {
          name           = "http1"
          container_port = 80
        }
        resources {
          limits = {
            cpu    = "1000m"
            memory = "512Mi"
          }
        }
        startup_probe {
          timeout_seconds   = 240
          period_seconds    = 240
          failure_threshold = 1
          tcp_socket {
            port = 80
          }
        }
      }

      containers {
        name  = "backend-1"
        image = "asia-northeast1-docker.pkg.dev/kyusyu-hackathon/main/backend:latest"
        resources {
          limits = {
            cpu    = "1000m"
            memory = "512Mi"
          }
        }
      }

      containers {
        name  = "frontend-stg-1"
        image = "asia-northeast1-docker.pkg.dev/kyusyu-hackathon/main/frontend-stg:latest"
        env {
          name  = "PORT"
          value = "3000"
        }
        resources {
          limits = {
            cpu    = "1000m"
            memory = "512Mi"
          }
        }
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}

data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "noauth" {
  location = google_cloud_run_service.main1.location
  project  = google_cloud_run_service.main1.project
  service  = google_cloud_run_service.main1.name

  policy_data = data.google_iam_policy.noauth.policy_data
}
