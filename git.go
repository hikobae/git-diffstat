package main

import (
	"os"
	"os/exec"
	"strings"
)

type Git struct {}

func (g *Git) Diff(args ...string) ([]string, error) {
	return g.run("diff", args...)
}

func (g *Git) run(command string, args ...string) ([]string, error) {
	xargs := append([]string{command}, args...)
	cmd := exec.Command("git", xargs...)
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	lines := strings.TrimSpace(string(out))
	return strings.FieldsFunc(lines, func(r rune) bool {
		return r == '\n'
	}), err
}
