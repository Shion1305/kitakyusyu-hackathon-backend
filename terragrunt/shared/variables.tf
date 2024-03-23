variable "project_id" {
  description = "The GCP project ID."
  default     = "kyusyu-hackathon"
}

variable "region" {
  description = "The GCP region."
  default     = "asia-northeast1"
}

variable "zone" {
  description = "The GCP zone."
  default     = "asia-northeast1-a"
}

variable "access_token" {
  description = "The access token for the GCP service account."
}
