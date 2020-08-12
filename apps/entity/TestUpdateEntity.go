package entity

import (
	"strconv"
	"tests/helper"
	"time"
)

// TestUpdateEntity ..
func TestUpdateEntity() helper.TestResult {
	result := helper.TestResult{
		Name: "entity.UpdateEntity",
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

	// has token ,has input but missing id => 400
	testCases = append(testCases, helper.TestMeasure{
		Input: (&map[string]interface{}{
			"value": strconv.FormatInt(time.Now().UnixNano(), 10),
		}),
		Token:    helper.AccessToken,
		Expected: 400,
		Assert:   "StatusCode",
	})

	// has token ,has input with id not found => 400
	testCases = append(testCases, helper.TestMeasure{
		Input: (&map[string]interface{}{
			"id":    -1,
			"value": "ent_" + strconv.FormatInt(time.Now().UnixNano(), 10),
		}),
		Token:    helper.AccessToken,
		Expected: 400,
		Assert:   "StatusCode",
	})

	// has token ,has input with id found => 200
	testCases = append(testCases, helper.TestMeasure{
		Input: (&map[string]interface{}{
			"id":    5,
			"value": "ent_" + strconv.FormatInt(time.Now().UnixNano(), 10),
		}),
		Token:    helper.AccessToken,
		Expected: 200,
		Assert:   "StatusCode",
	})

	// has token ,has input with id found => 200
	testCases = append(testCases, helper.TestMeasure{
		Input: (&map[string]interface{}{
			"id":    5,
			"value": "V",
		}),
		Token:    helper.AccessToken,
		Expected: 200,
		Assert:   "StatusCode",
	})

	for _, testCase := range testCases {
		res, err := helper.Put("http://localhost:3000/api/v1/entity", testCase.Input, testCase.Token)

		helper.MeasureFunc(*res, err, &result, testCase)

	}
	return result
}
