package jobs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TestCase struct {
	Method string
}

type jobsHandler struct{}

func RegisterHandlers() {
	handler := new(jobsHandler)
	http.Handle("/jobs", handler)
}

func (jh jobsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("im printing")
	switch r.Method {
	case http.MethodPost:
		var testCases TestCase
		err := json.NewDecoder(r.Body).Decode(&testCases)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		j, err := json.Marshal(testCases)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
