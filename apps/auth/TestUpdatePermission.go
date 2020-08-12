package auth

import "tests/helper"

// TestUpdatePermission ..
func TestUpdatePermission() helper.TestResult {
	result := helper.TestResult{
		Name: "auth_app.UpdatePermission",
	}

	var testCases []helper.TestMeasure

	// no token, user not found => 401
	testCases = append(testCases, helper.TestMeasure{
		Input: (&map[string]interface{}{}),
		Params: []helper.Parameter{
			helper.Parameter{
				Key:   "id",
				Value: "not-found",
			},
		},
		Expected: 401,
		Assert:   "StatusCode",
	})

	// no token, user found => 401
	testCases = append(testCases, helper.TestMeasure{
		Input: (&map[string]interface{}{}),
		Params: []helper.Parameter{
			helper.Parameter{
				Key:   "id",
				Value: "-MERqcgyetBoVo-iDSuC",
			},
		},
		Expected: 401,
		Assert:   "StatusCode",
	})

	// has token without permission => 401
	testCases = append(testCases, helper.TestMeasure{
		Input: (&map[string]interface{}{}),
		Token: helper.DummyToken,
		Params: []helper.Parameter{
			helper.Parameter{
				Key:   "id",
				Value: "-MERqcgyetBoVo-iDSuC",
			},
		},
		Expected: 401,
		Assert:   "StatusCode",
	})

	// has token with permission, user not found => 400
	testCases = append(testCases, helper.TestMeasure{
		Input: (&map[string]interface{}{}),
		Token: helper.AdminToken,
		Params: []helper.Parameter{
			helper.Parameter{
				Key:   "id",
				Value: "not-found",
			},
		},
		Expected: 400,
		Assert:   "StatusCode",
	})

	// has token with permission, user found, no input => 400
	testCases = append(testCases, helper.TestMeasure{
		Input: (&map[string]interface{}{}),
		Token: helper.AdminToken,
		Params: []helper.Parameter{
			helper.Parameter{
				Key:   "id",
				Value: "-METScEYm_81QcbyGuJt",
			},
		},
		Expected: 400,
		Assert:   "StatusCode",
	})

	// has token with permission, user found, has input => 200
	testCases = append(testCases, helper.TestMeasure{
		Input: (&map[string]interface{}{
			"roleIDs": []string{"editor"},
		}),
		Token: helper.AdminToken,
		Params: []helper.Parameter{
			helper.Parameter{
				Key:   "id",
				Value: "-MERqcgyetBoVo-iDSuC",
			},
		},
		Expected: 200,
		Assert:   "StatusCode",
	})

	for _, testCase := range testCases {
		res, err := helper.ParamPut("http://localhost:3000/api/v1/auth/:id/permission", testCase.Input, testCase.Token, testCase.Params)

		helper.MeasureFunc(*res, err, &result, testCase)

	}
	return result
}
