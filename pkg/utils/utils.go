package utils

import (
	"fmt"
	"os"
	"os/exec"
)

// CheckDocker checks that Docker is running on the users local system.
func CheckDocker() error {
	dockerCheck := exec.Command("docker", "stats", "--no-stream")
	if err := dockerCheck.Run(); err != nil {
		return fmt.Errorf("docker not running")
	}
	return nil
}

func RunCommandWithOutput(c *exec.Cmd) error {
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		return fmt.Errorf("piping output: %w", err)
	}
	fmt.Print("\n")
	return nil
}
