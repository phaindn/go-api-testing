package auth

import "tests/helper"

// TestGetUserInfo ..
func TestGetUserInfo() helper.TestResult {
	result := helper.TestResult{
		Name: "auth_app.GetUserInfo",
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
		res, err := helper.Get("http://localhost:3000/api/v1/auth/info", testCase.Input, testCase.Token)

		helper.MeasureFunc(*res, err, &result, testCase)

	}
	return result
}
