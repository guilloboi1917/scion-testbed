package handler

import (
	"fmt"
	"scionctl/internal/api"
	"scionctl/internal/config"
	"scionctl/internal/pprinter"
	"time"
)

// HandlePingStop handles the logic for stopping a ping operation
//
// args: [node]
func HandlePingStop(args []string) {
	node, exist := config.CmdNodeManager.GetNode(args[0])

	if !exist {
		pprinter.PrintError(fmt.Errorf("Node [%s] does not exist", args[0]))
		return
	}

	c := api.NewClient(api.ClientConfig{
		BaseURL: "http://" + node.Addr + ":" + fmt.Sprint(node.Port), Timeout: time.Second * 5,
	})

	resp, err := c.StopPing()
	if err != nil {
		pprinter.PrintError(err)
		return
	}

	pprinter.HTTPResponseToStdout(resp, err)
}
