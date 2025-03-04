package api

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

type API struct {
}

func (A API) PostAuth(w http.ResponseWriter, r *http.Request) {
	log.Traceln("API::PostAuth")
}

func (A API) GetStatus(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (A API) GetStore(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (A API) PutStore(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (A API) GetStoreStoreId(w http.ResponseWriter, r *http.Request, storeId string, params GetStoreStoreIdParams) {
	//TODO implement me
	panic("implement me")
}
