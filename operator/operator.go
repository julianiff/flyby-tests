package operator

import (
	"fmt"

	"github.com/julianiff/flyby-tests/job"
	"github.com/julianiff/flyby-tests/worker"
)

// the operator listens to a channel?

func StartOperator(jobQueue <-chan job.Job) {

	go func() {
		for {
			select {
			case nextJob := <-jobQueue:
				fmt.Printf("New Job received: ID: %s \n", fmt.Sprint(nextJob.ID))

				for _, testCase := range nextJob.TestCases {
					switch testCase.Mode {
					case "compare":
						worker.Compare(testCase)

					case "ruleset":
						worker.ApplyRuleset(testCase)
					}
				}
			}
		}
	}()

}
