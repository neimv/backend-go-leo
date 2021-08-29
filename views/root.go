package root

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandleRootRequest(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["hello"] = "World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
