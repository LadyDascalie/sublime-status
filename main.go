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
		fmt.Print(output, "◀ ", clock)
		return output, err
	}
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
		seed := emoji.Sprint(":cactus:")
		fmt.Print(folder, output, seed)
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
	if output == "0" {
    clean := emoji.Sprint(":small_blue_diamond:")
		fmt.Print(clean, "Clean dir. ")
	} else {
		sakura := emoji.Sprint(":small_red_triangle:")
		fmt.Print(sakura, output, "files changed ")
	}
	return output
}
