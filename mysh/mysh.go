package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		prompt()

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		// Remove the newline character.
		input = strings.TrimSuffix(input, "\n")

		args := parseInput(input)

		if err := execInput(args); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func prompt() {
	fmt.Print("mysh-0.1$ ")
}

func parseInput(line string) (args []string) {
	line = strings.TrimSuffix(line, "\n")
	return strings.Split(line, " ")
}

func execBuiltInCommand(args []string) error {
	cmd := exec.Command("bash", "-c", args...)
	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command return the error.
	return cmd.Run()
}

func execExternalCommand(args []string) error {
	cmd := exec.Command(args[0], args[1:]...)
	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command return the error.
	return cmd.Run()
}

func execInput(args []string) error {
	cmd := exec.Command(args[0], args[1:]...)
	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command return the error.
	return cmd.Run()
}
