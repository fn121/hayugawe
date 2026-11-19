variable "project" {
  type = string
  description = "unique project name that you created in GCP"
}

variable "region" {
  type = string
  description = "GCP region name"
  default = "us-east1"
}

variable "zone" {
  type = string
  description = "GCP zone. should match region"
  default = "us-east1-a"
}

variable "service_name" {
  type=string
  default = "task"
}

variable "billing_account" {
  type = string
  description = "GCP billing account"
}

variable "github_repo" {
  type = string
  description = "github repo. format username/reponame"
}
