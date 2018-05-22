package appsupport

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		retData := map[string]string{
			"status": "ok",
		}
		jsonData, _ := json.Marshal(retData)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	})
	return muxRouter
}
