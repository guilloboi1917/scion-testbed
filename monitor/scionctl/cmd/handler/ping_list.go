package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"scionctl/internal/api"
	"scionctl/internal/config"
	"scionctl/internal/pprinter"
	"time"
)

// Needs documentation
func HandlePingGetResults(args []string) {
	node, exist := config.CmdNodeManager.GetNode(args[0])

	if !exist {
		pprinter.PrintError(fmt.Errorf("Node [%s] does not exist", args[0]))
		return
	}

	c := api.NewClient(api.ClientConfig{
		BaseURL: "http://" + node.Addr + ":" + fmt.Sprint(node.Port), Timeout: time.Second * 5,
	})

	resp, err := c.GetResultsPing()
	if err != nil {
		pprinter.PrintError(err)
		return
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		pprinter.PrintError(err)
		return
	}

	var apiResp = api.PingListAPIResponse{}
	err = json.Unmarshal(bodyBytes, &apiResp)
	if err != nil {
		pprinter.PrintError(err)
		return
	}
	// We only care about the Data part
	if resp.StatusCode != http.StatusOK {
		pprinter.PrintError(fmt.Errorf("Error: %s", apiResp.Message))
		return
	}

	fmt.Printf("Available ping logfiles on %s:\n\n", node.Name)

	fileInfos := apiResp.Data

	pprinter.FileInfosTable(fileInfos)
}
