/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"scionctl/cmd/handler"

	"github.com/spf13/cobra"
)

// PingStopCmd represents the start ping command
var PingStopCmd = &cobra.Command{
	Use:   "stop <node>",
	Args:  cobra.ExactArgs(1),
	Short: "Stop a ping operation at a node",
	Long:  `Initiate a ping operation between two SCION nodes`,

	Run: func(cmd *cobra.Command, args []string) {
		handler.HandlePingStop(args)
	},
}

func init() {
	PingCmd.AddCommand(PingStopCmd)
}
