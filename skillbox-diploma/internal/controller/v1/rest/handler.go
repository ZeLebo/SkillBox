package rest

import (
	d "diploma/domain"
	s "diploma/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func StartServer() {
	router := mux.NewRouter()

	router.HandleFunc("/mms", handleMMS)
	router.HandleFunc("/support", handleSupport)
	router.HandleFunc("/accendent", handleAccendent)
	router.HandleFunc("/test", handleTest).Methods("GET", "OPTIONS")

	http.ListenAndServe("localhost:8383", router)
}

// todo all the above need to be in some struct
func handleMMS(w http.ResponseWriter, r *http.Request) {
	response(w, r, s.MMSCollection)
}

func handleSupport(w http.ResponseWriter, r *http.Request) {
	response(w, r, s.SupportCollection)
}

func handleAccendent(w http.ResponseWriter, r *http.Request) {
	response(w, r, s.AccendentCollection)
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Write([]byte(d.TestByteString))
}

func response(w http.ResponseWriter, _ *http.Request, responseStruct interface{}) {
	response, err := json.Marshal(responseStruct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(response)
}
