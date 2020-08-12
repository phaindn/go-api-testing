package helper

import (
	"fmt"
	"net/http"
)

// AccessToken full roles
var AccessToken string

// DummyToken no role
var DummyToken string

// EditorToken only editor role
var EditorToken string

// ReviewerToken only reviewer role
var ReviewerToken string

// AdminToken only superuser role
var AdminToken string

// PrintResults ...
func PrintResults(results []TestResult) {
	for _, result := range results {
		PrintResult(result)
	}
}

// PrintResult ...
func PrintResult(r TestResult) {
	fmt.Printf("[%s]: %d successes, %d fails:\n", r.Name, r.Success, r.Fail)
	for _, err := range r.Errors {
		fmt.Printf("---- %s\n", err)
	}
}

// CountTests ...
func CountTests(testCases []TestResult) TestResult {
	output := TestResult{
		Fail:    0,
		Success: 0,
	}
	for _, testCase := range testCases {
		output.Fail += testCase.Fail
		output.Success += testCase.Success
	}
	return output
}

// MeasureFunc ...
func MeasureFunc(res http.Response, err error, result *TestResult, testCase TestMeasure) {

	if err != nil {
		(*result).Fail++
		(*result).Errors = append((*result).Errors, err.Error())
		return
	}

	if testCase.Assert == "StatusCode" {
		if res.StatusCode == testCase.Expected {
			(*result).Success++
		} else {
			(*result).Fail++
			(*result).Errors = append((*result).Errors, fmt.Sprintf("Expected %w got %w", testCase.Expected, res.StatusCode))
		}
	}
}

// // RunTests ...
// func RunTests(url string, testCases []TestMeasure, result *TestResult, testFunc string) {
// 	funcs := map[string]interface{}{
// 		"Get":         Get,
// 		"Post":        Post,
// 		"Put":         Put,
// 		"Delete":      Delete,
// 		"ParamGet":    ParamGet,
// 		"ParamPost":   ParamPost,
// 		"ParamPut":    ParamPut,
// 		"ParamDelete": ParamDelete,
// 	}

// 	for _, testCase := range testCases {
// 		var res *http.Response
// 		var err error
// 		if strings.Index(testFunc, "Param") > -1 {
// 			res, err = funcs[testFunc].(func(string, *map[string]interface{}, string))(url, testCase.Input, testCase.Token)
// 		} else {
// 			res, err = funcs[testFunc].(func(string, *map[string]interface{}, string, []Parameter))(url, testCase.Input, testCase.Token, testCase.Params)
// 		}
// 		if err != nil {
// 			(*result).Fail++
// 			(*result).Errors = append((*result).Errors, err.Error())
// 			continue
// 		}

// 		MeasureFunc(*res, result, testCase)

// 	}
// }
