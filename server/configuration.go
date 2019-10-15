package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/redhatinsighs/insights-operator-controller/storage"
	"io"
	"io/ioutil"
	"net/http"
)

func getClusterConfiguration(writer http.ResponseWriter, request *http.Request, storage storage.Storage) {
	cluster, found := mux.Vars(request)["cluster"]
	if !found {
		writer.WriteHeader(http.StatusBadRequest)
		io.WriteString(writer, "Cluster ID needs to be specified")
		return
	}

	configuration, err := storage.ListClusterConfiguration(cluster)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		io.WriteString(writer, err.Error())
		return
	}
	json.NewEncoder(writer).Encode(configuration)
}

func newClusterConfiguration(writer http.ResponseWriter, request *http.Request, storage storage.Storage) {
	cluster, found := mux.Vars(request)["cluster"]
	if !found {
		writer.WriteHeader(http.StatusBadRequest)
		io.WriteString(writer, "Cluster ID needs to be specified")
		return
	}

	username, foundUsername := request.URL.Query()["username"]
	reason, foundReason := request.URL.Query()["reason"]
	description, foundDescription := request.URL.Query()["description"]

	if !foundUsername {
		writer.WriteHeader(http.StatusBadRequest)
		io.WriteString(writer, "User name needs to be specified\n")
		return
	}

	if !foundReason {
		writer.WriteHeader(http.StatusBadRequest)
		io.WriteString(writer, "Reason needs to be specified\n")
		return
	}

	if !foundDescription {
		writer.WriteHeader(http.StatusBadRequest)
		io.WriteString(writer, "Description needs to be specified\n")
		return
	}

	configuration, err := ioutil.ReadAll(request.Body)
	if err != nil || len(configuration) == 0 {
		writer.WriteHeader(http.StatusBadRequest)
		io.WriteString(writer, "Configuration needs to be provided in the request body")
		return
	}

	configurations, err := storage.CreateClusterConfiguration(cluster, username[0], reason[0], description[0], string(configuration))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		io.WriteString(writer, err.Error())
		return
	}
	json.NewEncoder(writer).Encode(configurations)
}

func enableClusterConfiguration(writer http.ResponseWriter, request *http.Request, storage storage.Storage) {
	cluster, found := mux.Vars(request)["cluster"]
	if !found {
		writer.WriteHeader(http.StatusBadRequest)
		io.WriteString(writer, "Cluster ID needs to be specified")
		return
	}

	username, foundUsername := request.URL.Query()["username"]
	reason, foundReason := request.URL.Query()["reason"]

	if !foundUsername {
		writer.WriteHeader(http.StatusBadRequest)
		io.WriteString(writer, "User name needs to be specified\n")
		return
	}

	if !foundReason {
		writer.WriteHeader(http.StatusBadRequest)
		io.WriteString(writer, "Reason needs to be specified\n")
		return
	}

	configurations, err := storage.EnableClusterConfiguration(cluster, username[0], reason[0])
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		io.WriteString(writer, err.Error())
		return
	}
	json.NewEncoder(writer).Encode(configurations)
}

func disableClusterConfiguration(writer http.ResponseWriter, request *http.Request, storage storage.Storage) {
	cluster, found := mux.Vars(request)["cluster"]
	if !found {
		writer.WriteHeader(http.StatusBadRequest)
		io.WriteString(writer, "Cluster ID needs to be specified")
		return
	}

	username, foundUsername := request.URL.Query()["username"]
	reason, foundReason := request.URL.Query()["reason"]

	if !foundUsername {
		writer.WriteHeader(http.StatusBadRequest)
		io.WriteString(writer, "User name needs to be specified\n")
		return
	}

	if !foundReason {
		writer.WriteHeader(http.StatusBadRequest)
		io.WriteString(writer, "Reason needs to be specified\n")
		return
	}

	configurations, err := storage.DisableClusterConfiguration(cluster, username[0], reason[0])
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		io.WriteString(writer, err.Error())
		return
	}
	json.NewEncoder(writer).Encode(configurations)
}