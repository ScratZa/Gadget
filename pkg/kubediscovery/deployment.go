package kubediscovery

type Deployment struct {
	Name string
	Yaml string
}

func NewDeployment(name, yaml string) *Deployment {
	return &Deployment{
		Name: name,
		Yaml: yaml,
	}
}
