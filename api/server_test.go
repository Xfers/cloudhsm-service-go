package api_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/Xfers/cloudhsm-service-go/api"
	"github.com/Xfers/cloudhsm-service-go/api/mocks"
	"github.com/stretchr/testify/assert"
)

func TestRunSignerServer(t *testing.T) {
	// Setup
	var testsData mocks.TestsData
	testsData.FillTestData()
	keys := testsData.Keys

	// Set up flags
	flags := make(map[string]interface{})
	flags["keys"] = keys

	// Run server
	go api.RunSignerServer(flags)

	// Wait for server to start
	for {
		resp, err := http.Get("http://localhost:8000/api/health")
		if err != nil {
			time.Sleep(100 * time.Millisecond)
			continue
		}
		if resp.StatusCode == 200 {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	// Create client
	client := &http.Client{}

	// Test cases
	for _, test := range testsData.Tests {
		for _, data := range test.Data {
			// Create request
			req, err := http.NewRequest(data.Method, "http://localhost:8000"+data.Endpoint, nil)
			if err != nil {
				t.Errorf("Error creating request: %v", err)
			}
			//Set request body
			reqBody, err := json.Marshal(data.Body)
			if err != nil {
				t.Errorf("Error marshalling request body: %v", err)
			}
			req.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", data.Headers.ContentType)

			// Send request
			resp, err := client.Do(req)
			if err != nil {
				t.Errorf("Error sending request: %v", err)
			}

			// Check response
			if resp.StatusCode != data.StatusCode {
				t.Errorf("Expected status code %v, got %v", data.StatusCode, resp.StatusCode)
			}

			// Read response body
			respBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}

			expected, err := json.Marshal(data.Expected)
			if err != nil {
				t.Fatal(err)
			}

			// Check response body
			// Compare two JSONs
			assert.JSONEq(t, string(expected), string(respBody), "In test case "+test.Name+": Response body does not match expected response body")
		}
	}
}
