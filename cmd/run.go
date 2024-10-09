package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/syx0310/Apple-Monitor-Go/pkg/apple"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start monitoring Apple products",
	Run: func(cmd *cobra.Command, args []string) {
		err := apple.StartMonitoring()
		if err != nil {
			fmt.Println("Error starting monitoring:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringP("region", "r", "", "Set the region")
	viper.BindPFlag("region", runCmd.Flags().Lookup("region"))

	runCmd.Flags().StringP("prefix", "p", "", "Set the URL prefix")
	viper.BindPFlag("prefix", runCmd.Flags().Lookup("prefix"))
}
