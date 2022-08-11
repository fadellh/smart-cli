/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/fadellh/smart-cli/usecase"
	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "exec blockchain configuration",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		fv := usecase.ChooseFabricVersion()
		log := usecase.ChooseMonitorLog()
		channels := usecase.AddMultipleOrgsInChannel()

		global := Global{
			FabricVersion: fv,
			Tls:           true,
			Monitoring: Monitoring{
				LogLevel: log,
			},
			Channels: channels,
		}
		fmt.Println(global)
	},
}

func init() {
	rootCmd.AddCommand(execCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// execCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// execCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
