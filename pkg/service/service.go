package service

import (
	"context"
	"fmt"
	"os"

	"github.com/Vinaychinnu/kubex/pkg/client"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
)

// Apply creates a Service from a YAML manifest
func Apply(filePath string, namespace string) error {
	clientset, err := client.NewClient()
	if err != nil {
		return err
	}

	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	var svc corev1.Service
	decoder := yaml.NewYAMLOrJSONDecoder(f, 1024)
	if err := decoder.Decode(&svc); err != nil {
		return err
	}

	// Namespace precedence: flag > YAML > default
	if namespace == "" {
		namespace = svc.Namespace
	}
	if namespace == "" {
		namespace = "default"
	}

	svc.Namespace = namespace

	svcClient := clientset.CoreV1().Services(namespace)
	_, err = svcClient.Create(context.TODO(), &svc, metav1.CreateOptions{})
	if err != nil {
		if apierrors.IsAlreadyExists(err) {
			fmt.Printf("Service %q already exists in namespace %q\n", svc.Name, namespace)
			return nil
		}
		return err
	}

	fmt.Printf("Service %q created in namespace %q\n", svc.Name, namespace)
	return nil
}

// Delete deletes a Service by name
func Delete(name string, namespace string) error {
	if namespace == "" {
		namespace = "default"
	}

	clientset, err := client.NewClient()
	if err != nil {
		return err
	}

	err = clientset.CoreV1().
		Services(namespace).
		Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	fmt.Printf("Service %q deleted from namespace %q\n", name, namespace)
	return nil
}
