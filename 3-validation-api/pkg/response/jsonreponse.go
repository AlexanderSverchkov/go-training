package response

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
