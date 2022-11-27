/* stress test for the server using ddosify */
package api

import (
	"fmt"
	"os/exec"
)

func stressTestDdosify() {
	// run ddosify
	// cmd := exec.Command("ddosify", "-c", "100", "-r", "100", "-m", "GET", "http://localhost:800")
	// err := cmd.Start()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // wait for ddosify to finish
	// err = cmd.Wait()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// Sign POST /api/sign/{keyName} , json {"data": "hello"} , keyame = "k1"
	// ddosify -m POST -t http://localhost:8000/api/sign/k1 -b "{\"data\": \"hello\"}" -l incremental

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
