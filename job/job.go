package job

import (
	"github.com/julianiff/flyby-tests/testCase"
)

type Job struct {
	Status    string
	ID        int
	TestCases []testCase.TestCase
}
