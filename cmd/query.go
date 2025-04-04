package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var queryPrompt string
var queryHost string

type acp struct {
	Prompt   string  `json:"prompt"`
	Response *string `json:"response"`
}

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Send a query to a remote agent",
	Long: `Send a query to a remote agent.

Example:
   hive query -p "What is the weather in San Francisco?" -h https://my-agent.example.com`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return queryAgent(queryPrompt, queryHost)
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
	queryCmd.PersistentFlags().StringVarP(&queryPrompt, "prompt", "p", "", "Prompt to send to the agent")
	queryCmd.PersistentFlags().StringVarP(&queryHost, "host", "u", "http://localhost:8080", "Host to query agent (default: http://localhost:8080)")
}

func queryAgent(query, url string) error {
	fmt.Printf("Prompt: %s\n", query)
	payload := map[string]string{"prompt": query}
	payloadBytes, _ := json.Marshal(payload)
	response := acp{}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payloadBytes))
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{
		Timeout: 300 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("unable to query agent: %w", err)
	}

	resBody, _ := io.ReadAll(res.Body)

	if err = json.Unmarshal(resBody, &response); err != nil {
		return fmt.Errorf("unable to process agent response: %w", err)
	}

	fmt.Printf("Response: %s\n", *response.Response)

	return nil
}
