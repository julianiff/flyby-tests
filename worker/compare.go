package worker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"

	"github.com/julianiff/flyby-tests/testCase"
)

func Compare(testCase testCase.TestCase) error {

	resp, err := http.Get(testCase.Target)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("Could not parse Target: %i", http.StatusBadRequest)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("Could not parse body: %i", http.StatusBadRequest)
	}

	data, _ := json.Marshal(testCase.Equals)
	result, err := JSONBytesEqual(data, body)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("Could not parse Equals: %i", http.StatusBadRequest)
	}
	fmt.Println(result)
	return nil
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
