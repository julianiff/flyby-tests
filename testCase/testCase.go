package testCase

type TestCase struct {
	Method string
	Target string
	Mode   string
	Equals interface{}
}

type TestCases []TestCase

func (tcs *TestCases) Validate() error {
	// Add validation step to TestCases to validate input

	return nil
}
