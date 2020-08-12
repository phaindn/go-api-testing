package auth

import "tests/helper"

// TestGetAll ..
func TestGetAll() helper.TestResult {
	result := helper.TestResult{
		Name: "auth_app.GetAll",
	}

	var testCases []helper.TestMeasure

	testCases = append(testCases, helper.TestMeasure{
		Input:    (&map[string]interface{}{}),
		Expected: 401,
		Assert:   "StatusCode",
	})

	testCases = append(testCases, helper.TestMeasure{
		Input:    (&map[string]interface{}{}),
		Token:    helper.AccessToken,
		Expected: 200,
		Assert:   "StatusCode",
	})

	for _, testCase := range testCases {
		res, err := helper.Get("http://localhost:3000/api/v1/auth", testCase.Input, testCase.Token)

		helper.MeasureFunc(*res, err, &result, testCase)

	}
	return result
}
