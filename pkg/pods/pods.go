package pods

import (
	"fmt"

	"github.com/Vinaychinnu/kubex/pkg/client"
)

// ListPods will later list pods.
// It verifies Kubernetes connectivity.
func ListPods() error {
	clientset, err := client.NewClient()
	if err != nil {
		return err
	}

	version, err := clientset.Discovery().ServerVersion()
	if err != nil {
		return err
	}

	fmt.Printf("Connected to Kubernetes cluster. Server version: %s\n", version.String())
	return nil
}
