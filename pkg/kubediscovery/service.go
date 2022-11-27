package kubediscovery

import (
	"context"
	"flag"
	"path/filepath"
	"time"

	v1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

type IKubeDiscoveryService interface {
	FetchDeployments() *v1.DeploymentList
}

type KubeDiscoveryService struct {
	Client kubernetes.Interface
}

func NewServiceWithClientSet(clientset *kubernetes.Clientset) IKubeDiscoveryService {
	return &KubeDiscoveryService{Client: clientset}
}

func NewService(InCluster bool) IKubeDiscoveryService {
	var config *rest.Config
	var err error

	if InCluster {
		config, err = rest.InClusterConfig()
	}

	if !InCluster {
		var kubeconfig *string
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()
		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	time.Sleep(10 * time.Second)
	return &KubeDiscoveryService{clientset}
}

func (s *KubeDiscoveryService) FetchDeployments() *v1.DeploymentList {
	deployments, err := s.Client.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	return deployments
}
