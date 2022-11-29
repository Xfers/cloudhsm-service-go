package test

import (
	"encoding/json"
	"fmt"
	"os"

	"os/exec"
)

type configs []config

type config struct {
	Name  string `json:"name"`
	Steps []step `json:"steps"`
}

type step struct {
	Name                     string      `json:"name"`
	ConcurrentUsersMax       int         `json:"concurrent_users_max"`
	ConcurrentUsersMin       int         `json:"concurrent_users_min"`
	ConcurrentUsersDelimiter int         `json:"concurrent_users_delimiter"`
	Endpoint                 string      `json:"endpoint"`
	RequestBody              interface{} `json:"request_body"`
}

type allResults struct {
	Name    string    `json:"name"`
	Results []Results `json:"results"`
}

type Results struct {
	Name    string   `json:"name"`
	Results []result `json:"results"`
}

type result struct {
	ConcurrentUsers int `json:"concurrent_users"`
	SuccessPerc     int `json:"success_perc"`
}

func RunLoadTest(configJson string) ([]allResults, error) {

	// check if valid JSON
	if !json.Valid([]byte(configJson)) {
		return nil, fmt.Errorf("Invalid JSON")
	}

	// unmarshal JSON
	var configs configs
	err := json.Unmarshal([]byte(configJson), &configs)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling JSON")
	}

	allResultsSlice := make([]allResults, 0)
	// loop trough configs
	for _, config := range configs {
		var allResults allResults
		allResults.Name = config.Name
		// loop through steps
		for _, step := range config.Steps {
			var results Results
			results.Name = step.Name

			increment := (step.ConcurrentUsersMax - step.ConcurrentUsersMin) / step.ConcurrentUsersDelimiter

			for i := step.ConcurrentUsersMin; i <= step.ConcurrentUsersMax; i += increment {
				body := step.RequestBody
				// body to string
				bodyString, err := json.Marshal(body)
				if err != nil {
					return nil, fmt.Errorf("Error marshalling JSON")
				}
				cmd := exec.Command(
					"ddosify", "-m", "POST", "-t", step.Endpoint, "-b", string(bodyString),
					"-n", fmt.Sprintf("%d", i), "-o", "stdout-json",
				)

				// get output of the command
				output, err := cmd.Output()
				if err != nil {
					return nil, err
				}

				// unmarshal JSON
				var r result
				err = json.Unmarshal(output, &r)
				if err != nil {
					return nil, fmt.Errorf("Error unmarshalling JSON")
				}

				r.ConcurrentUsers = i
				results.Results = append(results.Results, r)
			}
			allResults.Results = append(allResults.Results, results)
		}
		allResultsSlice = append(allResultsSlice, allResults)
	}

	// save to file as well
	_ = saveToFile(allResultsSlice)

	return allResultsSlice, nil
}

func saveToFile(allResultsSlice []allResults) error {
	f, err := os.Create("results.json")
	if err != nil {
		return fmt.Errorf("Error creating file")
	}
	defer f.Close()

	// marshal JSON
	resJson, err := json.Marshal(allResultsSlice)
	if err != nil {
		return fmt.Errorf("Error marshalling JSON")
	}

	// write to file
	_, err = f.Write(resJson)
	if err != nil {
		return fmt.Errorf("Error writing to file")
	}

	return nil
}
