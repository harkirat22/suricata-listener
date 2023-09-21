package whipper

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Whipper struct {
	clientset *kubernetes.Clientset
}

func NewWhipper() (*Whipper, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &Whipper{clientset: clientset}, nil
}

// FindPodByIP finds the pod with the given IP.
func (w *Whipper) FindPodByIP(ip string) (string, string, error) {
	pods, err := w.clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return "", "", err
	}

	for _, pod := range pods.Items {
		if pod.Status.PodIP == ip {
			return pod.Name, pod.Namespace, nil
		}
	}

	return "", "", fmt.Errorf("No pod found with IP: %s", ip)
}

// KillPod terminates the specified pod in the given namespace.
func (w *Whipper) KillPod(podName, namespace string) error {
	return w.clientset.CoreV1().Pods(namespace).Delete(context.TODO(), podName, metav1.DeleteOptions{})
}
