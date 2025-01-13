package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var AgentFile string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run an agent locally",
	Long: `Run your agent on your laptop. The agent will be available at http://localhost:8000



Examples:
    hive-cli run -f agent.py`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAgent(AgentFile)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	rootCmd.PersistentFlags().StringVarP(&AgentFile, "file", "f", "", "file of agent source code")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runAgent(file string) error {
	if file == "" {
		return fmt.Errorf("no source code file provided")
	}

	fmt.Println("Agent is available at http://localhost:8000")
	fmt.Println(`Sample command:
curl -X POST --json '{"prompt" : What is the capital of the United States?"}' http://localhost:8000`)
	fmt.Println("Press Ctrl+C to exit")

	executeAgent := exec.Command("python", file)
	if err := runCommandWithOutput(executeAgent); err != nil {
		return fmt.Errorf("error executing agent: %w", err)
	}

	return nil
}

func runCommand(c *exec.Cmd) error {
	if out, err := c.CombinedOutput(); err != nil {
		fmt.Println(string(out))
		return err
	}
	return nil
}

func runCommandWithOutput(c *exec.Cmd) error {
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		return fmt.Errorf("piping output: %w", err)
	}
	fmt.Print("\n")
	return nil
}
