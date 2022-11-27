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

// func AddDefinitions(c *gin.Context) {
// 	var newDefinition deployments.Deployment
// 	if err := c.BindJSON(&newDefinition); err != nil {
// 		return
// 	}

// 	definitionList = append(definitionList, newDefinition)
// 	c.IndentedJSON(http.StatusCreated, newDefinition)
// }

// func GetDefinitionByName(c *gin.Context) {
// 	name := c.Param("name")
// 	for _, a := range definitionList {
// 		if a.Name == name {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "definition not found"})
// }
