package whipper

import (
	"context"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// KillPod terminates the specified pod in the given namespace.
func KillPod(podName, namespace string) error {
	// Set up the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		return err
	}

	// Create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	// Delete the pod
	return clientset.CoreV1().Pods(namespace).Delete(context.TODO(), podName, nil)
}
