package handler

import (
	"net/http"
	"test/driver"
	"test/utils"

	"github.com/gorilla/mux"
)

var dbConnection *driver.DB

func DefaultMux(db *driver.DB) *mux.Router {
	dbConnection = db

	mux := mux.NewRouter().StrictSlash(true)

	mux.HandleFunc("/", defaultHandler)
	mux.HandleFunc("/sample/show", showSampleHandler)
	mux.NotFoundHandler = mux.NewRoute().HandlerFunc(defaultHandler).GetHandler()

	return mux
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	utils.ResponseJSON(w, r, true, 200, "Default response", make(map[string]interface{}))
}
