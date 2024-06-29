package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func mains(args []string) error {
	newArgs := make([]string, 0, len(args))
	for _, s := range args {
		newS := strings.ReplaceAll(s, `'`, `"`)
		newArgs = append(newArgs, newS)
	}
	cmd := exec.Command(newArgs[0], newArgs[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	if err := mains(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
