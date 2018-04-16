# Kubernetes Cluster

## Prerequisites

* [AWS CLI](https://aws.amazon.com/cli/)
* [kops](https://github.com/kubernetes/kops)
* [kubectl]()

## Install kops and kubectl

```bash
make dep-deploy
```

## Configure DNS

* Create hosted zone for subdomain
* Get NS servers of subdomain hosted zone
* Create subdomain NS record in parent domain with subdomain hosted zone NS servers

## Create S3 Bucket for State Store

```bash
aws s3api create-bucket --bucket clusters.k8s.jonathanfoster.io
aws s3api put-bucket-versioning --bucket clusters.k8s.jonathanfoster.io --versioning-configuration Status=Enabled
```

## Create Cluster

```bash
export KOPS_CLUSTER_NAME=k8s.jonathanfoster.io
export KOPS_STATE_STORE=s3://clusters.k8s.jonathanfoster-io

kops create cluster --cloud aws --zones us-east-1c --node-size t2.micro --master-size t2.micro --yes
```

## More Info

* [Lauching a Kubernetes cluster on AWS](https://github.com/kubernetes/kops/blob/master/docs/aws.md)
* [Manage Kubernetes Clusters on AWS Using Kops](https://aws.amazon.com/blogs/compute/kubernetes-clusters-aws-kops/)
