#context values for naming/label & tags
namespace   = "do"
environment = "playground"
name        = "devops-challenge"
attributes  = []
delimiter   = "-"
label_order = ["namespace", "environment", "name"]
tags = {
  Owner     = "Mariano"
  CreatedBy = "Terraform"
}

#compute values for K8s
node_count = 1