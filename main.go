package main

import (
	"net/http"

	"github.com/julianiff/flyby-tests/jobs"
)

func main() {
	jobs.RegisterHandlers()
	http.ListenAndServe(":4500", nil)
}
