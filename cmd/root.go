package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewCmdRoot() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "octo-pull",
		Short: "Octopus Energy API Tool",
		Long:  `Octo-Pull utilises the Octopus Energy API to pull meter readings to view and analyse your energy consumption`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("")
		},
	}

	rootCmd.AddCommand(NewCmdElectricity())

	return rootCmd
}

func Execute() {
	if err := NewCmdRoot().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
