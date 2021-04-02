package operator

import (
	"fmt"

	"github.com/julianiff/flyby-tests/job"
)

// the operator listens to a channel?

func StartOperator(jobQueue <-chan job.Job) {

	go func() {
		for {
			select {
			case msg := <-jobQueue:
				fmt.Printf("New Job received: ID: %s \n", fmt.Sprint(msg.ID))
				// do something with msg
			}
		}
	}()

}
