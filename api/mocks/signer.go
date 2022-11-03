package mocks

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"runtime"
)

type TestsData struct {
	Keys  map[string]string `json:"keys"`
	Tests []Test            `json:"tests"`
}

type Test struct {
	Name string     `json:"name"`
	Data []TestData `json:"data"`
}

type TestData struct {
	Endpoint   string            `json:"endpoint"`
	Body       map[string]string `json:"body"`
	Expected   map[string]string `json:"expected"`
	Method     string            `json:"method"`
	StatusCode int               `json:"statusCode"`
	Headers    Headers           `json:"headers"`
}

type Headers struct {
	ContentType string `json:"content-type"`
}

func (t *TestsData) FillTestData() {
	// Read in test data
	// get root path of the command ran
	_, filename, _, _ := runtime.Caller(0)
	rootPath := path.Dir(filename)
	// read in test data
	data, err := ioutil.ReadFile(rootPath + "/test_data.json")

	if err != nil {
		panic(err)
	}

	// Unmarshal json
	err = json.Unmarshal(data, &t)
	if err != nil {
		panic(err)
	}
}
