package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize states
	initCaptureState()
	initPingState()
	initScionPingState()

	// Register Capture API endpoints
	http.HandleFunc("/api/capture/start", startCapture)
	http.HandleFunc("/api/capture/stop", stopCapture)
	http.HandleFunc("/api/capture/status", getCaptureStatus)
	http.HandleFunc("/api/capture/files", getAvailableCaptures)

	// Register Ping API endpoints
	http.HandleFunc("/api/dispatch/ping/start", startPing)
	http.HandleFunc("/api/dispatch/ping/stop", stopPing)
	http.HandleFunc("/api/dispatch/ping/files", getAvailablePingResults)

	// Register ScionPing API endpoints
	http.HandleFunc("/api/dispatch/scionping/start", startScionPing)
	http.HandleFunc("/api/dispatch/scionping/stop", stopScionPing)
	http.HandleFunc("/api/dispatch/scionping/files", getAvailableScionPingResults)
	http.HandleFunc("/api/dispatch/scionping/status", getScionPingStatus)

	// Misc endpoints
	http.HandleFunc("/api/file", fileHandler)

	// Config API endpoints

	log.Println("SCION AS Container API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
