package pods

import "fmt"

// ListPods contains the business logic for listing pods.
// For now, its only printing a message.
// Later, will make this talk to the Kubernetes API.
func ListPods() error {
	fmt.Println("Listing pods (from pkg/pods)")
	return nil
}
