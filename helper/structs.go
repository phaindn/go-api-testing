package helper

// TestResult ...
type TestResult struct {
	Name     string
	Success  int
	Fail     int
	Errors   []string
	FailedAt []int
}

// Total test cases
func (m TestResult) Total() int {
	return m.Success + m.Fail
}

// TestMeasure ...
type TestMeasure struct {
	Params   []Parameter
	Input    *map[string]interface{}
	Token    string
	Expected interface{}
	Assert   string
}

// Parameter ...
type Parameter struct {
	Key   string
	Value string
}
