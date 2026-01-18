# kubex

kubex is a Go-based CLI tool for automating Kubernetes operations by
interacting directly with the Kubernetes API.

This project is built to understand how tools like kubectl work internally
and to practice real-world, production-inspired DevOps automation.

## Features

- List pods by namespace
- Create and delete namespaces
- Apply and delete deployments from YAML
- Apply and delete services from YAML
- Supports namespace flags
- Supports server-side dry-run validation
- Interacts directly with Kubernetes API using client-go

## Prerequisites

- Go 1.21 or later
- Access to a Kubernetes cluster
- A valid kubeconfig (kubectl must be configured)

## Installation

Clone the repository:

    git clone https://github.com/Vinaychinnu/kubex.git
    cd kubex

Build the binary:

    go build -o kubex

On Windows:

    go build -o kubex.exe

## Usage

After building the binary, run commands using:

    ./kubex <command>

On Windows:

    .\kubex.exe <command>

### Examples

    kubex pods list -n kube-system

    kubex namespace create dev
    kubex namespace delete dev

    kubex deploy apply -f manifests/nginx-deployment.yaml -n dev
    kubex deploy apply -f manifests/nginx-deployment.yaml -n dev --dry-run
    kubex deploy delete nginx-demo -n dev

    kubex service apply -f manifests/nginx-service.yaml -n dev
    kubex service delete nginx-svc -n dev

## Architecture

kubex follows a clean separation between CLI wiring and business logic.

A detailed architecture explanation is available here:
docs/architecture.md