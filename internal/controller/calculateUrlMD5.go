package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"urlsMd5/internal/dataModel"
	"urlsMd5/internal/service"
)

// CalculateUrlMd5 API returns md5 of respose to the client
func CalculateUrlMd5(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		return
	}

	Body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	defer r.Body.Close()
	var vData dataModel.UrlsModel

	err = json.Unmarshal(Body, &vData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	clientsData := make(chan string, len(vData.Urls))
	for _, data := range vData.Urls {
		clientsData <- data
	}
	go service.Worker(clientsData)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("OK")
}
