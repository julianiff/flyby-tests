package job

import (
	"github.com/julianiff/flyby-tests/testCase"
)

func AddNewJob(tCase []testCase.TestCase, jobQueue chan<- Job) (id int, err error) {
	// Add db connection here to get lattest ID
	// Status is initialized with pending, we want a non-blocking behaviour
	job := Job{Status: "pending", ID: 1, TestCases: tCase}

	go func() {
		jobQueue <- job
	}()

	return job.ID, nil
}
