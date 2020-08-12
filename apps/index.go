package apps

import (
	"backend/context/base/domain/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"tests/apps/auth"
	"tests/apps/entity"
	"tests/apps/intent"
	"tests/apps/utter"
	"tests/helper"
)

// Run ...
func Run() {

	setupToken("__TesterFull", "tester101", &helper.AccessToken)
	fmt.Println(helper.AccessToken)
	setupToken("__TesterAdmin", "tester101", &helper.AdminToken)
	fmt.Println(helper.AdminToken)
	setupToken("__TesterReviewer", "tester101", &helper.ReviewerToken)
	fmt.Println(helper.ReviewerToken)
	setupToken("__TesterEditor", "tester101", &helper.EditorToken)
	fmt.Println(helper.EditorToken)
	setupToken("__TesterDummy", "tester101", &helper.DummyToken)
	fmt.Println(helper.DummyToken)
	fmt.Printf(">>>>> TEST START\n______________________________________________________________________\n")

	var results []helper.TestResult
	results = append(results, auth.RunTests())
	results = append(results, entity.RunTests())
	results = append(results, intent.RunTests())
	results = append(results, utter.RunTests())

	final := helper.CountTests(results)
	fmt.Printf("______________________________________________________________________\n>>>>> TEST DONE: Total %d tests include %d successes and %d failed.\n", final.Total(), final.Success, final.Fail)
}

func setupToken(username string, password string, store *string) {
	res, err := helper.Post("http://localhost:3000/api/v1/login", &map[string]interface{}{
		"username": username,
		"password": password,
	}, "")
	if err == nil {
		body, errr := ioutil.ReadAll(res.Body)
		if errr == nil {
			var tokenDTO model.TokenDTO
			_ = json.Unmarshal([]byte(body), &tokenDTO)
			*store = tokenDTO.Token
			fmt.Println("Retrived token for " + username)
		} else {
			fmt.Printf("Can't login due to: %s\n", errr.Error())
		}
	} else {
		fmt.Printf("Can't login due to: %s\n", err.Error())
	}
}
