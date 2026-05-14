package git

import (
	"os/exec"
	"strings"
)

func GetChangedFiles() ([]string, error) {

	cmd := exec.Command(
		"git",
		"diff",
		"--cached",
		"--name-only",
	)

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	files := strings.Split(
		strings.TrimSpace(string(output)),
		"\n",
	)

	return files, nil
}