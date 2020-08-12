package auth

import "tests/helper"

// TestLogin ...
func TestLogin() helper.TestResult {
	result := helper.TestResult{
		Name: "auth_app.Login",
	}

	var testCases []helper.TestMeasure

	testCases = append(testCases, helper.TestMeasure{
		Input: (&map[string]interface{}{
			"username": "PhaiNDN",
			"password": "phaindn123",
		}),
		Expected: 200,
		Assert:   "StatusCode",
	})

	testCases = append(testCases, helper.TestMeasure{
		Input:    (&map[string]interface{}{}),
		Expected: 400,
		Assert:   "StatusCode",
	})

	testCases = append(testCases, helper.TestMeasure{
		Input: (&map[string]interface{}{
			"username": "PhaiNDN",
		}),
		Expected: 400,
		Assert:   "StatusCode",
	})

	testCases = append(testCases, helper.TestMeasure{
		Input: (&map[string]interface{}{
			"password": "PhaiNDN",
		}),
		Expected: 400,
		Assert:   "StatusCode",
	})

	testCases = append(testCases, helper.TestMeasure{
		Input: (&map[string]interface{}{
			"username": "wrongusername",
			"password": "wrongpassword",
		}),
		Expected: 401,
		Assert:   "StatusCode",
	})

	for index, testCase := range testCases {
		res, err := helper.Post("http://localhost:3000/api/v1/login", testCase.Input, testCase.Token)

		if err != nil {
			result.Fail++
			result.Errors = append(result.Errors, err.Error())
			result.FailedAt = append(result.FailedAt, index)
			continue
		}

		helper.MeasureFunc(*res, err, &result, testCase)

	}
	return result
}
