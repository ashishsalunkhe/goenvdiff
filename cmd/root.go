package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ashishsalunkhe/goenvdiff/internal"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var fromRef, toRef, path string
var jsonOutput bool

var rootCmd = &cobra.Command{
	Use:   "goenvdiff",
	Short: "Compare .env files between Git branches or commits",
	Run: func(cmd *cobra.Command, args []string) {
		fromBytes, err := internal.ReadEnvFromGit(fromRef, path)
		if err != nil {
			fmt.Printf("Error reading from %s: %v\n", fromRef, err)
			os.Exit(1)
		}

		toBytes, err := internal.ReadEnvFromGit(toRef, path)
		if err != nil {
			fmt.Printf("Error reading from %s: %v\n", toRef, err)
			os.Exit(1)
		}

		fromEnv, err := godotenv.Unmarshal(string(fromBytes))
		if err != nil {
			fmt.Printf("Error parsing env from %s: %v\n", fromRef, err)
			os.Exit(1)
		}

		toEnv, err := godotenv.Unmarshal(string(toBytes))
		if err != nil {
			fmt.Printf("Error parsing env from %s: %v\n", toRef, err)
			os.Exit(1)
		}

		diffs := internal.DiffEnvs(fromEnv, toEnv)

		if jsonOutput {
			jsonBytes, _ := json.MarshalIndent(diffs, "", "  ")
			fmt.Println(string(jsonBytes))
		} else {
			internal.PrintDiff(diffs)
		}
	},
}

func Execute() {
	rootCmd.PersistentFlags().StringVar(&fromRef, "from", "main", "Git ref to compare from")
	rootCmd.PersistentFlags().StringVar(&toRef, "to", "feature", "Git ref to compare to")
	rootCmd.PersistentFlags().StringVar(&path, "path", ".env", "Path to env file")
	rootCmd.PersistentFlags().BoolVar(&jsonOutput, "json", false, "Output in JSON format")
	rootCmd.Execute()
}
