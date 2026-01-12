package deploy

import (
	"context"
	"fmt"
	"os"

	"github.com/Vinaychinnu/kubex/pkg/client"
	appsv1 "k8s.io/api/apps/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
)

// Apply reads a deployment YAML and creates it in the cluster
func Apply(filePath string, namespace string) error {
	clientset, err := client.NewClient()
	if err != nil {
		return err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var deployment appsv1.Deployment
	decoder := yaml.NewYAMLOrJSONDecoder(file, 1024)

	if err := decoder.Decode(&deployment); err != nil {
		return fmt.Errorf("failed to decode yaml: %w", err)
	}

	if namespace == "" {
		namespace = deployment.Namespace
	}
	if namespace == "" {
		namespace = "default"
	}

	deployment.Namespace = namespace

	deploymentsClient := clientset.AppsV1().Deployments(namespace)

	_, err = deploymentsClient.Create(
		context.TODO(),
		&deployment,
		metav1.CreateOptions{},
	)

	if err != nil {
		if apierrors.IsAlreadyExists(err) {
			fmt.Printf(
				"Deployment %q already exists in namespace %q\n",
				deployment.Name,
				namespace,
			)
			return nil
		}
		return err
	}

	fmt.Printf(
		"Deployment %q created in namespace %q\n",
		deployment.Name,
		namespace,
	)

	return nil
}

// Delete deletes a deployment by name from the given namespace
func Delete(name string, namespace string) error {
	clientset, err := client.NewClient()
	if err != nil {
		return err
	}

	err = clientset.AppsV1().
		Deployments(namespace).
		Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	fmt.Printf(
		"Deployment %q deleted from namespace %q\n",
		name,
		namespace,
	)

	return nil
}
