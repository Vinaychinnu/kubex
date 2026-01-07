package pods

import (
	"context"
	"fmt"

	"github.com/Vinaychinnu/kubex/pkg/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ListPods lists pods in the given namespace.
func ListPods(namespace string) error {
	clientset, err := client.NewClient()
	if err != nil {
		return err
	}

	pods, err := clientset.CoreV1().
		Pods(namespace).
		List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	if len(pods.Items) == 0 {
		fmt.Printf("No pods found in namespace %q\n", namespace)
		return nil
	}

	fmt.Printf("Pods in namespace %q:\n", namespace)
	for _, pod := range pods.Items {
		fmt.Printf(
			"- %s\t%s\n",
			pod.Name,
			pod.Status.Phase,
		)
	}

	return nil
}
