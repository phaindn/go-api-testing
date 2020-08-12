package intent

import (
	"tests/helper"
)

// TestCreateIntent ..
func TestCreateIntent() helper.TestResult {
	result := helper.TestResult{
		Name: "intent.CreateIntent",
	}

	var testCases []helper.TestMeasure

	// no token => 401
	testCases = append(testCases, helper.TestMeasure{
		Input:    (&map[string]interface{}{}),
		Expected: 401,
		Assert:   "StatusCode",
	})

	// has token, no input => 400
	testCases = append(testCases, helper.TestMeasure{
		Input:    (&map[string]interface{}{}),
		Token:    helper.AccessToken,
		Expected: 400,
		Assert:   "StatusCode",
	})

	// has token, missing value => 400
	testCases = append(testCases, helper.TestMeasure{
		Input: (&map[string]interface{}{
			"level": 3,
		}),
		Token:    helper.AccessToken,
		Expected: 400,
		Assert:   "StatusCode",
	})

	// has token, missing level => 400
	testCases = append(testCases, helper.TestMeasure{
		Input: (&map[string]interface{}{
			"value": "_Test_ThamNganKepNgan",
		}),
		Token:    helper.AccessToken,
		Expected: 400,
		Assert:   "StatusCode",
	})

	// has token ,has input => 200
	testCases = append(testCases, helper.TestMeasure{
		Input: (&map[string]interface{}{
			"value": "_Test_ThamNganKepNgan",
			"level": 3,
		}),
		Token:    helper.AccessToken,
		Expected: 200,
		Assert:   "StatusCode",
	})

	for _, testCase := range testCases {
		res, err := helper.Post("http://localhost:3000/api/v1/intent", testCase.Input, testCase.Token)

		helper.MeasureFunc(*res, err, &result, testCase)

	}
	return result
}
