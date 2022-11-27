package main

import (
	"Gadget/pkg/http/rest"
	"Gadget/pkg/kubediscovery"
	"fmt"
	"log"
	"net/http"
)

func main() {

	KubeDiscoveryService := kubediscovery.NewService()
	router := rest.Hanlder(KubeDiscoveryService)
	fmt.Println("The beer server is on tap now: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
