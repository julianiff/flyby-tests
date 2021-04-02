package job

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julianiff/flyby-tests/testCase"
)

var jq chan<- Job

type jobHandler struct{}

func RegisterHandlers(jobQueue chan<- Job) {
	jq = jobQueue
	handler := new(jobHandler)
	http.Handle("/job", handler)
}

func (jh jobHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var testCases []testCase.TestCase
		err := json.NewDecoder(r.Body).Decode(&testCases)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		id, err := AddNewJob(testCases, jq)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprint(id)))

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
