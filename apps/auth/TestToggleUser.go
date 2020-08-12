package auth

import "tests/helper"

// TestToggleUser ..
func TestToggleUser() helper.TestResult {
	result := helper.TestResult{
		Name: "auth_app.ToggleUser",
	}

	var testCases []helper.TestMeasure

	// no body, no token
	testCases = append(testCases, helper.TestMeasure{
		Input:    (&map[string]interface{}{}),
		Expected: 401,
		Assert:   "StatusCode",
	})

	// has body, no token
	testCases = append(testCases, helper.TestMeasure{
		Input: (&map[string]interface{}{
			"userID":   "",
			"isActive": false,
		}),
		Expected: 401,
		Assert:   "StatusCode",
	})

	// no body, has token
	testCases = append(testCases, helper.TestMeasure{
		Input:    (&map[string]interface{}{}),
		Token:    helper.AccessToken,
		Expected: 400,
		Assert:   "StatusCode",
	})

	// has body, has token without required id
	testCases = append(testCases, helper.TestMeasure{
		Input: (&map[string]interface{}{
			"isActive": false,
		}),
		Token:    helper.AccessToken,
		Expected: 400,
		Assert:   "StatusCode",
	})

	// has body, has token with id not found
	testCases = append(testCases, helper.TestMeasure{
		Input: (&map[string]interface{}{
			"userID":   "not-found",
			"isActive": false,
		}),
		Token:    helper.AccessToken,
		Expected: 400,
		Assert:   "StatusCode",
	})

	// has body, has token with valid id: toggle themself
	testCases = append(testCases, helper.TestMeasure{
		Input: (&map[string]interface{}{
			"userID":   "-METUM-1GAscE9pCDYwe",
			"isActive": false,
		}),
		Token:    helper.AccessToken,
		Expected: 400,
		Assert:   "StatusCode",
	})

	// has body, has token with valid id: endable other
	testCases = append(testCases, helper.TestMeasure{
		Input: (&map[string]interface{}{
			"userID":   "-MERqcgyetBoVo-iDSuC",
			"isActive": false,
		}),
		Token:    helper.AccessToken,
		Expected: 200,
		Assert:   "StatusCode",
	})

	// has body, has token with valid id: disable other
	testCases = append(testCases, helper.TestMeasure{
		Input: (&map[string]interface{}{
			"userID":   "-MERqcgyetBoVo-iDSuC",
			"isActive": false,
		}),
		Token:    helper.AccessToken,
		Expected: 200,
		Assert:   "StatusCode",
	})

	for _, testCase := range testCases {
		res, err := helper.Post("http://localhost:3000/api/v1/auth/toggle-user", testCase.Input, testCase.Token)

		helper.MeasureFunc(*res, err, &result, testCase)

	}
	return result
}
