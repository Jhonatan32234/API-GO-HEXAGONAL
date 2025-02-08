package useCase_trabajador

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CheckPositionChange struct {
	apiURL string
}

func NewCheckpositionChange(apiURL string) *CheckPositionChange {
	return &CheckPositionChange{apiURL: apiURL}
}

func (csc *CheckPositionChange) Execute(trabajadorID int32) (string, error) {
	url := fmt.Sprintf("%s/trabajador/posicion/%d",csc.apiURL,trabajadorID)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    var result map[string]string
    if err := json.Unmarshal(body, &result); err != nil {
        return "", err
    }

	return result["posicion"], nil
}

