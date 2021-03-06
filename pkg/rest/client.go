package rest

import (
	"github.com/google/logger"
	"github.com/gorilla/mux"
	"myproj.com/clmgr-coordinator/config"
	"net/http"
	"strings"
)

type (
	client struct{}

	Client interface {
		Start() error
	}
)

func NewClient() Client {
	return &client{}
}

func (*client) Start() error {
	router := mux.NewRouter()

	router.HandleFunc("/hostname", HostnameHandler).Methods("GET")
	router.HandleFunc("/node", ListNodes).Methods("GET")
	router.HandleFunc("/node/{hostname}/label", AddLabelHandler).Methods("POST")
	router.HandleFunc("/resource", AddResource).Methods("POST")
	router.HandleFunc("/resource/list", ShowResources).Methods("GET")
	router.HandleFunc("/resource/{name}", ShowResource).Methods("GET")
	router.HandleFunc("/resource/{name}/conf", ConfigureResource).Methods("POST")
	router.HandleFunc("/resource/{name}", RemoveResource).Methods("DELETE")
	router.HandleFunc("/cluster/conf", ShowCluster).Methods("GET")
	router.HandleFunc("/cluster/conf", ConfigureCluster).Methods("POST")

	defaultPort := strings.Split(config.Config.CoordinatorAddress, ":")[1]
	logger.Infof("Start listener on port %s", defaultPort)
	http.ListenAndServe(":"+defaultPort, router)
	return nil
}
