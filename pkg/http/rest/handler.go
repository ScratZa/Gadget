package rest

import (
	"Gadget/pkg/kubediscovery"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Hanlder(kubediscovery kubediscovery.IKubeDiscoveryService) http.Handler {
	router := httprouter.New()
	router.GET("/deployments", GetDeployments(kubediscovery))
	return router
}

func GetDeployments(kubediscovery kubediscovery.IKubeDiscoveryService) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		deployments := kubediscovery.FetchDeployments()
		json.NewEncoder(w).Encode(deployments)
	}
}
