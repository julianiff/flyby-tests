package testCase

type TestCase struct {
	Method string
	Target string
	Mode   string
	Equals map[string]interface{}
}

type TestCases []TestCase

func (tcs TestCases) Validate() error {
	// Add validation step to TestCases if err not nil

	return nil
}
