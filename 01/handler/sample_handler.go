package handler

import (
	"net/http"
	"test/repository"
	"test/utils"
)

func showSampleHandler(w http.ResponseWriter, r *http.Request) {
	samepleRepo := repository.NewSQLSampleRepo(dbConnection)

	list, err := samepleRepo.Fetch()

	if err != nil {
		panic(err)
	}

	results := make(map[string]interface{})
	results["list"] = list

	utils.ResponseJSON(w, r, true, 200, "Sample response list", results)
}
