/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"embed"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var sourceFile string

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build your agent",
	Long: `Build your agent

Example:
    hive-cli build -f source.py`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return buildAgent(sourceFile)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	buildCmd.PersistentFlags().StringVarP(&sourceFile, "file", "f", "", "agent source code file")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

//go:embed template.py
var f embed.FS

func buildAgent(source string) error {

	// convert source to string with proper indentation
	// write template to new file, replacing ## INSERT ## line with user code

	// Open the execution template
	template := "template.py"
	templateFile, err := f.Open(template)
	// templateFile, err := os.Open(template)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer templateFile.Close()

	// Open user source code
	sourceFile, err := os.Open(source)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer templateFile.Close()

	// Create a new file to write the modified lines
	agentFile, err := os.Create("agent.py")
	if err != nil {
		fmt.Println("Error creating modified file:", err)
		return err
	}
	defer agentFile.Close()

	// Create a scanner to read the existing file line by line
	scanner := bufio.NewScanner(templateFile)

	// Write the lines of the existing file to the modified file
	for scanner.Scan() {
		line := scanner.Text()

		// Check if the line is the target line
		if strings.Contains(line, "##### INSERT USER CODE HERE #####") {
			// Add the new file to the target line
			codeScanner := bufio.NewScanner(sourceFile)
			for codeScanner.Scan() {
				codeLine := codeScanner.Text()
				_, err := agentFile.WriteString("    " + codeLine + "\n")
				if err != nil {
					fmt.Println("Error writing to modified file:", err)
					return err
				}
			}
		} else {
			// Write the existing line to the modified file
			_, err := agentFile.WriteString(line + "\n")
			if err != nil {
				fmt.Println("Error writing to modified file:", err)
				return err
			}
		}
	}

	fmt.Println("Runnable AI Agent created: agent.py.")
	fmt.Println("To run locally, use this command:")
	fmt.Println("hive-cli run -f agent.py")

	return nil
}
