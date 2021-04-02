package job

import (
	"github.com/julianiff/flyby-tests/testCase"
)

func AddNewJob(tCase []testCase.TestCase, jobQueue chan<- Job) (id int, err error) {
	job := Job{Status: false, ID: 1, TestCases: tCase}

	go func() {
		jobQueue <- job
	}()

	return 1, nil
}
