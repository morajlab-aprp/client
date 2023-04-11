package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
)

func execute() {
	cmd := exec.Command("ls")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))
}

func ask() error {
	const port = 3000
	requestURL := fmt.Sprintf("http://localhost:%d", port)

	res, err := http.Get(requestURL)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	fmt.Printf("Body: %s", body)

	return nil
}

func main() {
	// ask()
	execute()
}
