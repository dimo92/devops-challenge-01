terraform {
  backend "s3" {
    endpoint                    = "sfo3.digitaloceanspaces.com"
    region                      = "us-west-1" // needed
    bucket                      = "do-devops-challenge-playground-space-tfstates"
    key                         = "devops/do/k8s/terraform.tfstate"
    skip_credentials_validation = true
    skip_metadata_api_check     = true
  }
}