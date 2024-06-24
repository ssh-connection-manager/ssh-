package git

import (
	"os/exec"
	"ssh+/app/output"
	"strings"
)

func GetRelease() string {
	cmd := exec.Command("git", "describe", "--abbrev=0", "--tags")

	test, err := cmd.Output()
	if err != nil {
		output.GetOutError("Error getting repository version")
	}

	return strings.TrimSpace(string(test))
}
