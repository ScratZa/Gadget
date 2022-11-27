package kubediscovery

type IKubeDiscoveryService interface {
	FetchDeployments() *Deployment
}

type KubeDiscoveryService struct {
}

func NewService() IKubeDiscoveryService {
	return &KubeDiscoveryService{}
}

func (s *KubeDiscoveryService) FetchDeployments() *Deployment {
	return NewDeployment("test", "test")
}
