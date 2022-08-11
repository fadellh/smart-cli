/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fadellh/smart-cli/usecase"
	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "exec blockchain configuration",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("[Plese choose Fabric version]")
		fv := usecase.ChooseFabricVersion()

		fmt.Println("[Plese choose Monitoring Log]")
		log := usecase.ChooseMonitorLog()

		dataOrg := usecase.DataOrg()

		fmt.Println("[Plese Add Organization inside channel]")
		channels := usecase.AddMultipleOrgsInChannel(dataOrg)

		fmt.Println("[Plese Add Chaincode]")
		chainCodes := usecase.AddMultipleChainCode(channels)

		global := Global{
			FabricVersion: fv,
			Tls:           true,
			Monitoring: Monitoring{
				LogLevel: log,
			},
		}
		rawDataOrg := json.RawMessage(dataOrg)
		smr := SmartSample{
			Schema:     "https://github.com/hyperledger/releases/download/1.1.0/schema.json",
			Global:     global,
			Orgs:       rawDataOrg,
			Channels:   channels,
			ChainCodes: chainCodes,
		}

		smrJson, err := json.MarshalIndent(smr, "", " ")
		if err != nil {
			fmt.Printf("Failed %v\n", err)
			os.Exit(1)
		}

		fmt.Println("-->[Generate blockchain configuration]")
		err = ioutil.WriteFile("sample.json", smrJson, 0644)
		if err != nil {
			fmt.Printf("Failed %v\n", err)
			os.Exit(1)
		}
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
