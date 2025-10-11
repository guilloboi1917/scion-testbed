package handler

import (
	"fmt"
	"scionctl/internal/api"
	"scionctl/internal/config"
	"scionctl/internal/printer"
	"time"
)

// Parse input to IPs
//
// args[0] is the source node as ScionNode type
//
// args[1] is the destination node as ScionNode type
//
// Example usage: scionctl ping start scion11 scion31 -c 5
func HandlePingStart(args []string, pingCount int) {
	senderNode, receiverNode, err := validateArgs(args)
	if err != nil {
		printer.PrintError(err)
		return
	}

	// Timeout currently after 5s, arbitrary value
	c := api.NewClient(api.ClientConfig{
		BaseURL: receiverNode.Addr + ":" + fmt.Sprint(receiverNode.Port),
		Timeout: time.Second * 5,
	})

	// Need to make sure pingCount is set
	resp, err := c.StartPing(senderNode.Addr, pingCount)
	printer.HTTPResponseToStdout(resp, err)
}

func validateArgs(args []string) (api.ScionNode, api.ScionNode, error) {
	senderNode, exists := config.CmdNodeManager.GetNode(args[0])
	if !exists {
		return api.ScionNode{}, api.ScionNode{}, fmt.Errorf("Sender node [%s] does not exist", args[0])
	}
	receiverNode, exists := config.CmdNodeManager.GetNode(args[1])
	if !exists {
		return api.ScionNode{}, api.ScionNode{}, fmt.Errorf("Receiver node [%s] does not exist", args[1])
	}

	return *senderNode, *receiverNode, nil
}
