package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var taskType string
var taskId string
var showConfig bool

type Config struct {
	TaskType string `json:"taskType,omitempty"`
	TaskID  string `json:"taskId,omitempty"`
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure aicli settings",
	Run: func(cmd *cobra.Command, args []string) {

		config := Config{}

		fileData, err := os.ReadFile("config.json")
		if err == nil {
			_ = json.Unmarshal(fileData, &config)
		}

		if cmd.Flags().Changed("type") {
			config.TaskType = taskType
			fmt.Println("Task type set to:", taskType)
		}

		if cmd.Flags().Changed("task") {
			config.TaskID = taskId
			fmt.Println("Task ID set to:", taskId)
		}

		if showConfig {
			fmt.Println("Current configuration:")
			fmt.Printf("Task Type: %s\n", config.TaskType)
			fmt.Printf("Task ID: %s\n", config.TaskID)
			return
		}

		file, err := os.Create("config.json")
		if err != nil {
			fmt.Println("Error creating config file:", err)
			return
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")

		err = encoder.Encode(config)
		if err != nil {
			fmt.Println("Error writing config:", err)
			return
		}

		fmt.Println("Configuration complete.")
	},
}

func init() {

	configCmd.Flags().StringVarP(
		&taskType,
		"type",
		"t",
		"feat",
		"include task type in commit message",
	)

	configCmd.Flags().StringVarP(
		&taskId,
		"task",
		"i",
		"",
		"include task ID in commit message",
	)

	configCmd.Flags().BoolVarP(
		&showConfig,
		"show",
		"s",
		false,
		"show current configuration",
	)

	rootCmd.AddCommand(configCmd)
}


func ReadConfig() (Config, error) {

	config := Config{}

	fileData, err := os.ReadFile("config.json")
	if err != nil {

		// dosya yoksa boş config dön
		if os.IsNotExist(err) {
			return config, nil
		}

		return config, err
	}

	err = json.Unmarshal(fileData, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}