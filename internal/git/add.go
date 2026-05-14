package git

import (
	"os"
	"os/exec"
)


func Add() error {
	cmd := exec.Command(
		"git",
		"add",
		".",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}