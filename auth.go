package main

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func authenticate() *kubernetes.Clientset {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()

	// get the cluster kubeconfig and use it in the clientset
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})
	config, err := kubeconfig.ClientConfig()
	if err != nil {
		panic(err)
	}
	return kubernetes.NewForConfigOrDie(config)
}
