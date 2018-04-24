package main

import (
	"fmt"
	"os"
	"os/exec"
)

var newsSources = []string{
	"https://mashable.com",
	"https://techcrunch.com",
	"https://wired.com",
	"https://arstechnica.com",
}

func main() {
	for _, v := range newsSources {
		cmd := exec.Command("open", "-a", "/Applications/Google Chrome.app", fmt.Sprintf("%s", v))
		cmd.Stderr = os.Stdout
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error encountered while trying to run cmd: %v\nERR:%s", cmd.Args, err.Error())
		}
	}
}
