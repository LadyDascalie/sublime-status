package main

import (
	"fmt"
	"os/exec"
	"strings"
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
		return output, err // Let it run despite the error
	}
	clock := emoji.Sprint(":clock1:")
	fmt.Print(output, "◀ ", clock)
	return output, err
}

func getTime() (printed bool, err error) {
	t := time.Now()
	h := t.Hour()
	m := t.Minute()
	_, err = fmt.Print(h, ":", m, " ▶")
	if err != nil {
		fmt.Print("Error printing time to stdout")

	} else {
		printed = true
	}
	return printed, err
}

func getDir() (output string) {
	cmd := exec.Command("bash", "-c", `echo "${PWD##*/}"`)
	stdout, err := cmd.Output()
	output = string(stdout)
	if err != nil {
		fmt.Print("Err. getting folder name")
	} else {
		folder := emoji.Sprint(":open_file_folder:")
		branch := emoji.Sprint(":octopus:")
		fmt.Print(folder, output, branch)
	}
	return output
}

func countDirtyFiles() (output string) {
	cmd := exec.Command("bash", "-c", `git status | grep modified: | wc -l`)
	stdout, err := cmd.Output()
	output = string(stdout)
	if err != nil {
		fmt.Print("Couldnt read dir.")
	}
	split := strings.SplitAfter(strings.TrimSpace(output), "0")[0]
	if split == "0" {
		clean := emoji.Sprint(":small_blue_diamond:")
		fmt.Print(clean, "Clean dir. ")
	} else {
		dirty := emoji.Sprint(":small_red_triangle:")
		fmt.Print(dirty, output, "files changed ")
	}
	return output
}
