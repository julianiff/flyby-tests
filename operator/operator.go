package operator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"

	"github.com/julianiff/flyby-tests/job"
)

// the operator listens to a channel?

func StartOperator(jobQueue <-chan job.Job) {

	go func() {
		for {
			select {
			case nextJob := <-jobQueue:
				fmt.Printf("New Job received: ID: %s \n", fmt.Sprint(nextJob.ID))

				for _, testCase := range nextJob.TestCases {
					fmt.Println("Executing", testCase.Target)

					resp, err := http.Get(testCase.Target)
					if err != nil {
						log.Println(err)
						return
					}
					defer resp.Body.Close()

					body, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						log.Println(err)
						return
					}

					data, _ := json.Marshal(testCase.Equals)
					result, err := JSONBytesEqual(data, body)
					if err != nil {
						log.Println(err)
						return
					}
					fmt.Println(result)
				}
			}
		}
	}()

}

func JSONBytesEqual(a, b []byte) (bool, error) {
	var j, j2 interface{}
	if err := json.Unmarshal(a, &j); err != nil {
		return false, err
	}
	if err := json.Unmarshal(b, &j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}
