package namespace

import (
	"context"
	"fmt"

	"github.com/Vinaychinnu/kubex/pkg/client"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// List lists all namespaces
func List() error {
	clientset, err := client.NewClient()
	if err != nil {
		return err
	}

	namespaces, err := clientset.CoreV1().
		Namespaces().
		List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, ns := range namespaces.Items {
		fmt.Println(ns.Name)
	}

	return nil
}

// Create creates a namespace
func Create(name string) error {
	clientset, err := client.NewClient()
	if err != nil {
		return err
	}

	ns := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}

	_, err = clientset.CoreV1().
		Namespaces().
		Create(context.TODO(), ns, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	fmt.Printf("Namespace %q created\n", name)
	return nil
}

// Delete deletes a namespace
func Delete(name string) error {
	clientset, err := client.NewClient()
	if err != nil {
		return err
	}

	err = clientset.CoreV1().
		Namespaces().
		Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	fmt.Printf("Namespace %q deleted\n", name)
	return nil
}
