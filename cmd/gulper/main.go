package main

import (
	"Gadget/pkg/http/rest"
	"Gadget/pkg/kubediscovery"
	"log"
	"net/http"
)

func main() {

	KubeDiscoveryService := kubediscovery.NewService(true)
	router := rest.Hanlder(KubeDiscoveryService)
	log.Fatal(http.ListenAndServe(":8080", router))
}
