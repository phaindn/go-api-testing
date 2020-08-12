package entity

import (
	"tests/helper"
)

// TestDeleteEntity ..
func TestDeleteEntity() helper.TestResult {
	result := helper.TestResult{
		Name: "entity.DeleteEntity",
	}

	var testCases []helper.TestMeasure

	// no token => 401
	testCases = append(testCases, helper.TestMeasure{
		Params: []helper.Parameter{
			helper.Parameter{
				Key:   "id",
				Value: "50",
			},
		},
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

	// has token ,has input with id not found => 400
	testCases = append(testCases, helper.TestMeasure{
		Params: []helper.Parameter{
			helper.Parameter{
				Key:   "id",
				Value: "-1",
			},
		},
		Input:    (&map[string]interface{}{}),
		Token:    helper.AccessToken,
		Expected: 400,
		Assert:   "StatusCode",
	})

	// // has token ,has input with id found => 200 !!!! ONLY USE WHEN DONE
	// testCases = append(testCases, helper.TestMeasure{
	// 	Params: []helper.Parameter{
	// 		helper.Parameter{
	// 			Key:   "id",
	// 			Value: "50",
	// 		},
	// 	},
	// 	Input:    (&map[string]interface{}{}),
	// 	Token:    helper.AccessToken,
	// 	Expected: 200,
	// 	Assert:   "StatusCode",
	// })

	for _, testCase := range testCases {
		res, err := helper.ParamDelete("http://localhost:3000/api/v1/entity?id=:id", testCase.Input, testCase.Token, testCase.Params)

		helper.MeasureFunc(*res, err, &result, testCase)

	}
	return result
}
