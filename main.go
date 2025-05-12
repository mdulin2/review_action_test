package main

import (
	"fmt"
	"log"
	"os/exec"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())

	cmd := exec.Command("bash", "-c", "ls -l && echo 'Hello from bash'")

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(output))
}
