package api

// Define API routes
var (
	PingStartRoute         string = "/api/dispatch/ping/start"
	PingStopRoute          string = "/api/dispatch/ping/stop"
	PingListAvailableRoute string = "/api/dispatch/ping/files"
	PingStatusRoute        string = "/api/dispatch/ping/status"

	ScionStartPingRoute     string = "/api/dispatch/scionping/start"
	ScionStopPingRoute      string = "/api/dispatch/scionping/stop"
	ScionListAvailableRoute string = "/api/dispatch/scionping/files"
	ScionPingStatusRoute    string = "/api/dispatch/scionping/status"

	CaptureStartRoute         string = "/api/capture/start"
	CaptureStopRoute          string = "/api/capture/stop"
	CaptureListAvailableRoute string = "/api/capture/files"
	CaptureStatusRoute        string = "/api/capture/status"
)
