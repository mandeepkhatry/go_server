package response

import (
	"encoding/json"
	"net/http"
)

//Define Data creation response
func ResponseCreated(w http.ResponseWriter, code int, id interface{}) {
	response, _ := json.Marshal(map[string]interface{}{"id": id})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
