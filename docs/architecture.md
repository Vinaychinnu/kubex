# kubex Architecture

kubex is a Go-based CLI tool designed to automate common Kubernetes operations
by interacting directly with the Kubernetes API.

## High-Level Flow

User
→ CLI Commands (cobra)
→ Business Logic (pkg)
→ Kubernetes API (client-go)
→ Cluster

## Project Structure

- cmd/
  - CLI commands and flag parsing
  - No Kubernetes API logic

- pkg/
  - Core business logic
  - Interacts with Kubernetes resources

- pkg/client
  - Kubernetes client initialization
  - kubeconfig and authentication handling

## Design Principles

- Separation of concerns
- Reusable and testable logic
- Safe defaults for write operations
- Incremental extensibility
