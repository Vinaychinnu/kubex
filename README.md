# kubex

kubex is a Go-based CLI tool for automating Kubernetes operations by
interacting directly with the Kubernetes API.

This project is built to understand how tools like kubectl
work internally and to practice real-world DevOps automation.

## Features

- List pods by namespace
- Create and delete namespaces
- Apply and delete deployments from YAML
- Supports namespace flags
- Interacts directly with Kubernetes API using client-go

## Example Usage
```bash
kubex pods list -n kube-system
kubex namespace create dev
kubex deploy apply -f manifests/nginx-deployment.yaml -n dev
kubex deploy delete nginx-demo -n dev
