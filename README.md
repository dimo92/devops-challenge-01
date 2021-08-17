# devops-challenge-01

## Template Go / REST / Gorm
- https://github.com/qiangxue/go-rest-api


## Resources pre-requisites
- Digital Ocean Account

    - Space for tfstates persistence generated 
        - https://github.com/MattMorgis/digitalocean-spaces-terraform-backend
        - https://dev.to/jmarhee/digitalocean-spaces-as-a-terraform-backend-3lck

    - How to generate DO Token? 
        - https://docs.digitalocean.com/reference/api/api-reference/

    - How to generate Space & SpaceAccess Token / SpaceSecret Key? 
        - https://docs.digitalocean.com/reference/api/spaces-api/
            
## Workspace tools/config pre-requisites
- tfswitch: tool form manage Terraform versions
    - https://tfswitch.warrensbox.com/

- python3 / pip : required for aws-cli & pre-commit hooks
    - https://towardsdatascience.com/installing-multiple-alternative-versions-of-python-on-ubuntu-20-04-237be5177474

- aws-cli : Digital Ocean Space share config extends S3 Credentials
    - https://docs.aws.amazon.com/es_es/cli/latest/userguide/install-cliv2.html

- doctl : access to k8s cluster & download .kubeconfig file
    - https://github.com/digitalocean/doctl/releases/download/v1.63.1/doctl-1.63.1-linux-amd64.tar.gz

- kubectl : manage k8s resources
    - https://kubernetes.io/es/docs/tasks/tools/

- kubeselect : tool for switch between k8s clusters/context
    - https://github.com/fatliverfreddy/kubeselect


## Terraform

```
export SPACES_ACCESS_TOKEN="XXXXXXXXXX"
export SPACES_SECRET_KEY="XXXXXXXXXX"
export SPACE_BUCKET_NAME="XXXXXXXXXX"

#Workking on /devops/do/k8s

terraform init \
	-backend-config="access_key=$SPACES_ACCESS_TOKEN" \
	-backend-config="secret_key=$SPACES_SECRET_KEY" \
	-backend-config="bucket=$SPACE_BUCKET_NAME"
	
terraform plan    
```

## Technical debt / Nice to have
    - avoid autoincrement on insert failure, prevent jump between id's: innodb_autoinc_lock_mode=2