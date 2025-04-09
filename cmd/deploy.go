package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/psschwei/agent-cli/pkg/utils"
	"github.com/spf13/cobra"
)

var deployImage string
var agentName string

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy your agent",
	Long: `Deploy your agent to Kubernetes.

Examples:
    hive deploy -t my-agent:latest`,

	RunE: func(cmd *cobra.Command, args []string) error {
		return deployAgent(agentName, deployImage)
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	deployCmd.PersistentFlags().StringVarP(&deployImage, "tag", "t", "", "Agent image to deploy")
	deployCmd.PersistentFlags().StringVarP(&agentName, "name", "n", "", "Name to use for agent")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deployCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func deployAgent(name, image string) error {

	if name == "" {
		return fmt.Errorf("deploy agent: no name specified")
	}

	if image == "" {
		return fmt.Errorf("deploy agent: no image specified")
	}

	kubeCheck := exec.Command("kubectl", "config", "current-context")
	if err := kubeCheck.Run(); err != nil {
		return fmt.Errorf("unable to connect to kubernetes cluster")
	}

	deployment := fmt.Sprintf(`
apiVersion: apps/v1
kind: Deployment
metadata:
  name: %s
  labels:
    app: %s
spec:
  replicas: 1
  selector:
    matchLabels:
      app: %s
  template:
    metadata:
      labels:
        app: %s
    spec:
      containers:
      - name: %s
        image: %s
        imagePullPolicy: IfNotPresent
        envFrom:
        - configMapRef:
            name: agent-configmap
        - secretRef:
            name: agent-secrets
        ports:
        - containerPort: 8080`, name, name, name, name, name, image)

	service := fmt.Sprintf(`
apiVersion: v1
kind: Service
metadata:
  name: %s-svc
spec:
  selector:
    app: %s
  ports:
  - protocol: TCP
    port: 8080
    targetPort: 8080`, name, name)

	// create deployment
	kubeDeploy := exec.Command("kubectl", "create", "-f", "-")
	kubeDeploy.Stdin = strings.NewReader(deployment)
	if err := utils.RunCommandWithOutput(kubeDeploy); err != nil {
		return fmt.Errorf("unable to create deployment: %w", err)
	}

	// create service
	kubeSvc := exec.Command("kubectl", "create", "-f", "-")
	kubeSvc.Stdin = strings.NewReader(service)
	if err := utils.RunCommandWithOutput(kubeSvc); err != nil {
		return fmt.Errorf("unable to create service: %w", err)
	}

	return nil
}
