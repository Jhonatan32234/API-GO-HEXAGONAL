package useCase_trabajador

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CheckSalaryChange struct {
	apiURL string
}

func NewCheckSalaryChange(apiURL string) *CheckSalaryChange {
	return &CheckSalaryChange{apiURL: apiURL}
}

func (csc *CheckSalaryChange) Execute(trabajadorID int32) (int32, error) {
	url := fmt.Sprintf("%s/trabajador/salario/%d",csc.apiURL,trabajadorID)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("error: status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return 0, err
    }

    var result map[string]int32
    if err := json.Unmarshal(body, &result); err != nil {
        return 0, err
    }

    return result["salario"], nil
}

