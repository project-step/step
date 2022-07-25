package main

import (
	"fmt"
	"os"

	"github.com/stepneko/neko-session/cli"

	"github.com/spf13/cobra"
)

func main() {
	var stepCmd = &cobra.Command{
		Use:          "step [command] (flags)",
		Short:        "Step command-line interface and server",
		Long:         `Step command-line interface and server.`,
		SilenceUsage: true,
	}

	stepCmd.AddCommand(
		cli.CreateSessionCmd(),
	)

	if err := stepCmd.Execute(); err != nil {
		fmt.Println("Execute command failed", err)
		os.Exit(2)
	}
}
