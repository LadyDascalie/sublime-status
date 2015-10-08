package main

import (
	"fmt"
	"os/exec"
	"time"

	"gopkg.in/kyokomi/emoji.v1"
)

func main() {
	countDirtyFiles()
	getDir()
	getBranch()
	getTime()
}

func getBranch() (output string, err error) {
	cmd := exec.Command("bash", "-c", `git branch | grep "*"`)
	stdout, err := cmd.Output()
	output = string(stdout)
	if err != nil {
		fmt.Print("Not a Git repo ◀")
		return output, err // Let it run despite the error
	} else {
		clock := emoji.Sprint(":clock1:")
		fmt.Print(output, clock, "◀ ")
		return output, err
	}
}

func getTime() {
	t := time.Now()
	h := t.Hour()
	m := t.Minute()
	fmt.Print(h, ":", m, " ▶")
}

func getDir() {
	cmd := exec.Command("bash", "-c", `echo "${PWD##*/}"`)
	stdout, err := cmd.Output()
	output := string(stdout)
	if err != nil {
		fmt.Print("Err. getting folder name")
	} else {
		fmt.Print("|| ", output, "|| ")
	}
}

func countDirtyFiles() {
	cmd := exec.Command("bash", "-c", `git status | grep modified: | wc -l`)
	stdout, err := cmd.Output()
	output := string(stdout)
	if err != nil {
		fmt.Print("Couldnt read dir.")
	}
	if output == "0" {
		fmt.Print("Clean dir. ")
	} else {
		rose := emoji.Sprint(":rose:")
		fmt.Print(rose, output, "files changed ")
	}

}
