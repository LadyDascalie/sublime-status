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
	output := string(stdout)
	if err != nil {
		fmt.Print("Not a Git repo ◀")
		return // Let it run despite the error
	} else {
		fmt.Print(output, "◀ ")
	}

}

func getTime() {
	t := time.Now()
	h := t.Hour()
	m := t.Minute()

	fmt.Print(h, ":", m, " ▶")
}
