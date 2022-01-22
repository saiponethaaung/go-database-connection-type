package utils

import (
	"encoding/json"
	"net/http"
)

// ResponseJSON .
func ResponseJSON(w http.ResponseWriter, r *http.Request, responseStatus bool, responseCode int, responseMsg string, responseData map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")

	data := make(map[string]interface{})

	data["status"] = responseStatus
	data["code"] = responseCode
	data["msg"] = responseMsg
	data["data"] = responseData

	bytes, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong!"))
		return
	}

	w.WriteHeader(responseCode)
	w.Write(bytes)
}

// Response500JSON .
func Response500JSON(w http.ResponseWriter, r *http.Request, responseMsg string, responseData map[string]interface{}) {
	ResponseJSON(w, r, false, 500, responseMsg, responseData)
}

// Response422JSON .
func Response422JSON(w http.ResponseWriter, r *http.Request, responseMsg string, responseData map[string]interface{}) {
	ResponseJSON(w, r, false, 422, responseMsg, responseData)
}

// Response404JSON .
func Response404JSON(w http.ResponseWriter, r *http.Request, responseMsg string, responseData map[string]interface{}) {
	ResponseJSON(w, r, false, 404, responseMsg, responseData)
}
