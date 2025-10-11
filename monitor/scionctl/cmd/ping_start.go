/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"scionctl/cmd/handler"

	"github.com/spf13/cobra"
)

var (
	pingCount int
)

// PingStartCmd represents the start ping command
var PingStartCmd = &cobra.Command{
	Use:   "start <source> <target>",
	Args:  cobra.ExactArgs(2),
	Short: "Start a ping operation between two nodes",
	Long:  `Initiate a ping operation between two SCION nodes`,

	Run: func(cmd *cobra.Command, args []string) {
		handler.HandlePingStart(args, pingCount)
	},
}

func init() {
	PingCmd.AddCommand(PingStartCmd)
	PingStartCmd.Flags().IntVarP(&pingCount, "count", "c", -1, "Number of ping packets to send")
}
