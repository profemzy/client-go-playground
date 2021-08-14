package main

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
)

func validateContainerResourceSpec(container corev1.Container) {
	if container.Resources.Requests.Cpu().IsZero() || container.Resources.Requests.Memory().IsZero() {
		fmt.Println(container.Name + " does not have resource limit specified")
	}
}
