package auth

import (
	"tests/helper"
)

// RunTests ...
func RunTests() helper.TestResult {
	var testResults []helper.TestResult
	testResults = append(testResults, TestLogin())
	testResults = append(testResults, TestGetAll())
	testResults = append(testResults, TestGetAllRoles())
	testResults = append(testResults, TestGetUserInfo())
	testResults = append(testResults, TestToggleUser())
	testResults = append(testResults, TestUpdatePermission())

	helper.PrintResults(testResults)
	return helper.CountTests(testResults)
}
