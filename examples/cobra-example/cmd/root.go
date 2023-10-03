package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cobra-example",
		Short: "Cobra example",
		Long:  "Cobra example.",
		Run:   runRootCmd,
	}

	cmd.AddCommand(newVersionCmd())

	return cmd
}

func runRootCmd(cmd *cobra.Command, args []string) {
	fmt.Println("Hello, cobra-example!")
}

func Execute() {
	cobra.CheckErr(newRootCmd().Execute())
}
