package api

import "net"

// Define request and response types

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type PingStartRequest struct {
	Dst   string `json:"dst"`
	Count *int   `json:"count,omitempty"`
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
