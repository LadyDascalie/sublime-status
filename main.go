package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	getBranch()
	getTime()
}

func getBranch() {
	cmd := exec.Command("bash", "-c", `git branch | grep "*"`)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	output := string(stdout)
	fmt.Println(output, "  <|  ")
}

func getTime() {
	t := time.Now()
	h := t.Hour()
	m := t.Minute()

	fmt.Println(h, ":", m, "  |>")
}
