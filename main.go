package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	// getDir()
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

func getDir() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Print("Couldn't get path")
	} else {
		fmt.Print("|| ", dir, " ||")
	}
}

func minifyPath(path string) {
	strings.Split(path, "/")
}
