/* stress test for the server using ddosify */
package api

import (
	"fmt"
	"os/exec"
)

func stressTestDdosify() {

	// Incremental load test
	cmd := exec.Command("ddosify", "-m", "POST", "-t", "http://localhost:8000/api/sign/k1", "-b", "{\"data\": \"hello\"}", "-l", "incremental")
	err := cmd.Start()
	if err != nil {
		fmt.Println("failed at start")
		fmt.Println(err)
	}
	// wait for ddosify to finish
	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
	}

	// Linear load test
	cmd = exec.Command("ddosify", "-m", "POST", "-t", "http://localhost:8000/api/sign/k1", "-b", "{\"data\": \"hello\"}")
	err = cmd.Start()
	if err != nil {
		fmt.Println("failed at start")
		fmt.Println(err)
	}
	// wait for ddosify to finish
	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
	}

	//TODO: write tests in config files https://github.com/ddosify/ddosify#config-file
}

func StressTest() {
	stressTestDdosify()
}
