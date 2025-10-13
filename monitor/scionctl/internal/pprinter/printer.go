package pprinter

// Using go pretty print
// https://github.com/jedib0t/go-pretty

// Functionalities for printing to stdout in a nice way

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"scionctl/internal/api"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func PrintTable(input TableInput) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(input.header)
	t.AppendRows(input.content)
	t.Render()
}

func FileInfosTable(fileInfos []api.FileInfo) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"idx", "fn", "size"})

	for _, fi := range fileInfos {
		t.AppendRow(table.Row{fi.Index, fi.Name, fi.Size})
	}

	t.Render()
}

// Simple HTTP Response to Stdout
func HTTPResponseToStdout(resp *http.Response, err error) {
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Check if response is nil
	if resp == nil {
		fmt.Println("Error: nil response received")
		return
	}

	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:")

	buf := new(strings.Builder)
	_, copyErr := io.Copy(buf, resp.Body)
	if copyErr != nil {
		PrintError(copyErr)
		return
	}

	jsonTransformer := text.NewJSONTransformer("", "    ")
	prettyJSON := jsonTransformer(buf.String())
	fmt.Println(prettyJSON)

	// Add a newline for better formatting
	fmt.Println()
}

func PrintError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
