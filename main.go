package main

import (
	"net/http"

	"github.com/julianiff/flyby-tests/job"
	"github.com/julianiff/flyby-tests/operator"
)

func main() {
	jobQueue := make(chan job.Job)

	job.RegisterHandlers(jobQueue)
	operator.StartOperator(jobQueue)

	http.ListenAndServe(":4600", nil)
}
