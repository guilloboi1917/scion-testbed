package api

import "net"

// Define request and response types

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type FileInfo struct {
	Index int32  `json:"index"`
	Name  string `json:"name"`
	Size  int64  `json:"size"`
}

type PingStartRequest struct {
	Dst   string `json:"dst"`
	Count *int   `json:"count,omitempty"`
}

type PingListAPIResponse struct {
	APIResponse
	Data []FileInfo `json:"data,omitempty"` // Override the Data field
}

type ScionNode struct {
	Addr string
	Name string
	Port int16
	ISD  int16
	AS   int16
}

func NodeToIP(node ScionNode) net.IPAddr {
	// Cast it to net.IPAddr
	return net.IPAddr{IP: net.IP(node.Addr)}
}
