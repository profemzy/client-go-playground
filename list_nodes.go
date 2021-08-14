package main

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func listNodes(clientset *kubernetes.Clientset) *corev1.NodeList {
	nodeList, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})

	if err != nil {
		return nil
	}

	return nodeList
}
