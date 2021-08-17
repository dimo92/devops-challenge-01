module "label" {
  source      = "cloudposse/label/null"
  namespace   = var.namespace
  environment = var.environment
  stage       = var.stage
  name        = var.name
  attributes  = var.attributes
  delimiter   = var.delimiter
  label_order = var.label_order
  tags        = var.tags
}

resource "digitalocean_kubernetes_cluster" "k8s_main_cluster" {
  name    = "master-${module.label.id}"
  region  = "nyc1"
  version = "1.21.2-do.2"

  node_pool {
    name       = "pool-${module.label.id}-${var.node_count}"
    size       = "s-2vcpu-4gb"
    node_count = var.node_count
  }
}

provider "kubernetes" {
  load_config_file = false
  host             = digitalocean_kubernetes_cluster.k8s_main_cluster.endpoint
  token            = digitalocean_kubernetes_cluster.k8s_main_cluster.kube_config[0].token
  cluster_ca_certificate = base64decode(
    digitalocean_kubernetes_cluster.k8s_main_cluster.kube_config[0].cluster_ca_certificate
  )
}