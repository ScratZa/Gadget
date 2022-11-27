package main

import (
	"Gadget/pkg/kubediscovery"
	"testing"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestHello(t *testing.T) {
	want := "test"

	if got := kubediscovery.NewDeployment("test", "test"); got.Name != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestGetDeployment(t *testing.T) {
	obj := &appsv1.Deployment{ObjectMeta: metav1.ObjectMeta{
		Name: "demo-deployment",
	},
	}
	clientSet := fake.NewSimpleClientset(obj)
	var service *kubediscovery.KubeDiscoveryService = &kubediscovery.KubeDiscoveryService{Client: clientSet}

	want := v1.DeploymentList{Items: []v1.Deployment{*obj}}
	got := *service.FetchDeployments()
	if got.Items[0].Namespace != want.Items[0].Namespace {
		t.Errorf("Miss matched Deployments")
	}

}
