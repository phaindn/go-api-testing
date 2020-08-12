package entity

import (
	"strconv"
	"tests/helper"
	"time"
)

// TestCreateEntity ..
func TestCreateEntity() helper.TestResult {
	result := helper.TestResult{
		Name: "entity.CreateEntity",
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

	// has token ,has input and value existed => 200 & return existed item
	testCases = append(testCases, helper.TestMeasure{
		Input: (&map[string]interface{}{
			"value": "V",
		}),
		Token:    helper.AccessToken,
		Expected: 200,
		Assert:   "StatusCode",
	})

	// has token ,has input => 200
	testCases = append(testCases, helper.TestMeasure{
		Input: (&map[string]interface{}{
			"value": strconv.FormatInt(time.Now().UnixNano(), 10),
		}),
		Token:    helper.AccessToken,
		Expected: 200,
		Assert:   "StatusCode",
	})

	for _, testCase := range testCases {
		res, err := helper.Post("http://localhost:3000/api/v1/entity", testCase.Input, testCase.Token)

		helper.MeasureFunc(*res, err, &result, testCase)

	}
	return result
}
