package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/syx0310/Apple-Monitor-Go/cmd"
	"github.com/syx0310/Apple-Monitor-Go/pkg/apple"
	"github.com/syx0310/Apple-Monitor-Go/pkg/logger"
)

func main() {
	cobra.OnInitialize(func() {
		if err := apple.InitConfig(); err != nil {
			fmt.Println("Failed to initialize config:", err)
			os.Exit(1)
		}
	})
	logger.InitLogger()

	cmd.Execute()
}
