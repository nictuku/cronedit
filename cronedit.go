// Package cronedit manipulates a user's crontab.
package cronedit

import (
	"fmt"
	"os/exec"
	"strings"
)

// edit inserts a command unless it's already present and returns the new string.
func edit(input string, insert string) (bool, string) {
	// Uses simple \n line splitting. If we expect crontabs to be very large, bufio.Scanner would be
	// better.
	for _, line := range strings.Split(input, "\n") {
		if line == insert {
			return false, input
		}
	}
	input = strings.TrimRight(input, "\n")
	return true, fmt.Sprintf("%v\n%v\n", input, insert)
}

// Insert adds the specified command to the current user's crontab, unless it's already present. It
// returns true if modifications were made.
func Insert(command string) (bool, error) {
	out, err := exec.Command("crontab", "-l").Output()
	if err != nil {
		return false, err
	}
	needChange, newContent := edit(string(out), command)
	if !needChange {
		return false, nil
	}
	// There is no other programmatic way of editing a user's crontab. That operation requires root
	// permission, so we need to use crontab(1) because of its suid privileges.
	cronEdit := exec.Command("crontab", "-")
	cronEdit.Stdin = strings.NewReader(newContent)
	return true, cronEdit.Run()
}
