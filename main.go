package main

import (
	"fmt"

	_ "k8s.io/client-go/plugin/pkg/client/auth/azure"
)

const deploymentName = "test-deployment"
const defaultNamespace = "default"

func main() {

	// Authenticate to the cluster
	clientset := authenticate()

	// Display the names of the nodes in the cluster
	nodeList := listNodes(clientset)
	for _, n := range nodeList.Items {
		fmt.Println(n.Name)
	}

	// Get all pods in a namespace
	podList := getPods(clientset, "backend")

	for _, p := range podList.Items {
		// check if pod has resource definition
		for _, c := range p.Spec.Containers {
			validateContainerResourceSpec(c)
		}
	}
	// Create test deployment if it does not already exist
	createDeployment(clientset)

	// Delete test deployment
	deleteDeployment(clientset, deploymentName)
}

func int32Ptr(i int32) *int32 { return &i }
