package useCase_jefe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ExperienceResponse struct {
	AnosExperiencia int `json:"anosexperiencia"`
}

type CheckExperienceChange struct {
	apiURL string
}

func NewCheckExperienceChange(apiURL string) *CheckExperienceChange {
	return &CheckExperienceChange{apiURL: apiURL}
}

func (cec *CheckExperienceChange) Execute(jefeID int32) (int, error) {
	url := fmt.Sprintf("%s/jefeproyecto/experiencia/%d", cec.apiURL, jefeID)
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

	var result ExperienceResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}

	return result.AnosExperiencia, nil
}
