package job

import (
	"github.com/julianiff/flyby-tests/testCase"
)

type Job struct {
	Status    bool
	ID        int
	TestCases []testCase.TestCase
}
