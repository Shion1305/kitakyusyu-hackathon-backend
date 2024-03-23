provider "google" {
  project = var.project_id
  region  = var.region
  zone    = var.zone

  access_token = var.access_token
}
