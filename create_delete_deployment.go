package main

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	appv1 "k8s.io/api/apps/v1"
)

func createDeployment(clientset *kubernetes.Clientset)  {

	labels := map[string]string{"app": "test-app"}

	newDeployment := &appv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:   deploymentName,
			Labels: labels,
		},
		Spec: appv1.DeploymentSpec{
			Replicas: int32Ptr(1),

			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "web",
							Image: "nginx",
							Resources: corev1.ResourceRequirements{
								Limits: corev1.ResourceList{
									"cpu":    resource.MustParse("500m"),
									"memory": resource.MustParse("128Mi"),
								},
								Requests: corev1.ResourceList{
									"cpu":    resource.MustParse("250m"),
									"memory": resource.MustParse("64Mi"),
								},
							},
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	deployment, err := clientset.AppsV1().Deployments(defaultNamespace).Get(context.Background(), deploymentName, metav1.GetOptions{})

	if errors.IsNotFound(err) {
		// Create deployment only if it does not already exist
		deployment, err = clientset.AppsV1().Deployments(defaultNamespace).Create(context.Background(), newDeployment, metav1.CreateOptions{})
		if err != nil {
			panic(err)
		}
		fmt.Println(deployment.Name)
	}
}

func deleteDeployment(clientset *kubernetes.Clientset, deploymentName string)  {
	err := clientset.AppsV1().Deployments(defaultNamespace).Delete(context.Background(), deploymentName, metav1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
}
