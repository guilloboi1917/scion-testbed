package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

func getFilesFromDirectory(directory string) ([]FileInfo, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	fileInfos := make([]FileInfo, 0)

	for idx, file := range files {
		if !file.IsDir() {
			info, err := file.Info()
			if err != nil {
				continue
			}
			fileInfos = append(fileInfos, FileInfo{
				Index: int32(idx),
				Name:  file.Name(),
				Size:  info.Size(),
			})
		}
	}

	return fileInfos, nil
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fileName := r.URL.Query().Get("name")
	if fileName == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(APIResponse{
			Status:  "error",
			Message: "Invalid request - No filename found",
		})
		return
	}

	src := r.URL.Query().Get("src")
	if src == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(APIResponse{
			Status:  "error",
			Message: "Invalid request - No src found (ping | scionping | capture)",
		})
		return
	}

	var fileDir = ""
	var fileSuffix = ".log"

	switch src {
	case "ping":
		fileDir = pingResultLocation
	case "scionping":
		fileDir = scionPingResultLocation
	default:
		fileDir = packetCapturesResultLocation
		fileSuffix = ".pcap"
	}

	var fileLocation = fileDir + fileName + fileSuffix

	if _, err := os.Stat(fileLocation); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(APIResponse{
				Status:  "error",
				Message: "File doesn't exist",
			})
			return
		}
	}

	// For file download, set appropriate headers and serve file
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, fileLocation)
}
