package worker

import (
	"fmt"
	"strings"

	"github.com/julianiff/flyby-tests/testCase"
)

func ApplyRuleset(testCase testCase.TestCase) {

	str := fmt.Sprint(testCase.Equals)

	s := strings.Split(str, " ")

	fmt.Println(s)

	// Delegate the work in here to a worker
	// Operator should not contain too much logic just
	// delegate the work
	fmt.Println("rules:", testCase.Equals)
}
