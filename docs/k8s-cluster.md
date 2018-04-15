# Kubernetes Cluster

## Prerequisites

* [AWS CLI](https://aws.amazon.com/cli/)
* [kops](https://github.com/kubernetes/kops)
* [kubectl]()

## Install Deployment Tools

```bash
make dep-deploy
```

## Create kops IAM User

```bash
aws iam create-group --group-name kops

aws iam attach-group-policy --policy-arn arn:aws:iam::aws:policy/AmazonEC2FullAccess --group-name kops
aws iam attach-group-policy --policy-arn arn:aws:iam::aws:policy/AmazonRoute53FullAccess --group-name kops
aws iam attach-group-policy --policy-arn arn:aws:iam::aws:policy/AmazonS3FullAccess --group-name kops
aws iam attach-group-policy --policy-arn arn:aws:iam::aws:policy/IAMFullAccess --group-name kops
aws iam attach-group-policy --policy-arn arn:aws:iam::aws:policy/AmazonVPCFullAccess --group-name kops

aws iam create-user --user-name kops
aws iam add-user-to-group --user-name kops --group-name kops
aws iam create-access-key --user-name kops
```

## Configure DNS

* Create hosted zone for subdomain
* Get NS servers of subdomain hosted zone
* Create subdomain NS record in parent domain with subdomain hosted zone NS servers

```bash
# TODO: Test scripts
export DOMAIN=jonathanfoster.io
export SUBDOMAIN=k8s.jonathanfoster.io

# Create hosted zone for subdomain and get name servers
export NAME_SERVERS=$(aws route53 create-hosted-zone --name $(SUBDOMAIN) --caller-reference $(uuidgen) | jq .DelegationSet.NameServers)

# Get parent hosted zone ID
export PARENT_ZONE_ID=$(aws route53 list-hosted-zones | jq '.HostedZones[] | select(.Name=="$(DOMAIN).") | .Id)

# TODO: Escape quotes, same for $NAME_SERVERS
# Create subdomain NS record in parent domain
echo "{
  "Comment": "Create a subdomain NS record in the parent domain",
  "Changes": [
    {
      "Action": "CREATE",
      "ResourceRecordSet": {
        "Name": "$SUBDOMAIN",
        "Type": "NS",
        "TTL": 300,
        "ResourceRecords": $NAME_SERVERS
      }
    }
  ]
}" > subdomain.json
aws route53 change-resource-record-sets --hosted-zone-id $PARENT_ZONE_ID --change-batch file://subdomain.json

# Test DNS configuration
dig ns $SUBDOMAIN
```

## Create S3 Bucket for State Store

```bash
aws s3api create-bucket --bucket k8s-jonathanfoster-io
aws s3api put-bucket-versioning --bucket k8s-jonathanfoster-io --versioning-configuration Status=Enabled
export KOPS_STATE_STORE=s3://k8s-jonathanfoster-io
```

## Create Cluster

```bash
export NODE_SIZE=${NODE_SIZE:-t2.nano}
export MASTER_SIZE=${MASTER_SIZE:-t2.nano}
export ZONES=${ZONES:-"us-east-1a"}
export MASTER_ZONES=${MASTER_ZONES:-"us-east-1a"}
export KOPS_STATE_STORE="s3://k8s-jonathanfoster-io"

kops create cluster k8s.jonathanfoster.io \
  --node-count 1 \
  --zones $ZONES \
  --node-size $NODE_SIZE \
  --master-size $MASTER_SIZE \
  --master-zones $MASTER_ZONES \
  --networking weave \
  --topology private \
  --bastion="true" \
  --yes
```

## More Info

* [Lauching a Kubernetes cluster on AWS](https://github.com/kubernetes/kops/blob/master/docs/aws.md)
* [Manage Kubernetes Clusters on AWS Using Kops](https://aws.amazon.com/blogs/compute/kubernetes-clusters-aws-kops/)
