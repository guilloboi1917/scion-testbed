// Contains code to configure nodes
//
// -- Start, Stop scion services
// -- Edit Path Policy
// -- IPTables (?)
// -- Chaos (?)

//Questions:
	//Upload default policy file which is then edited? Or only upload file when needed?
	//Upload default config file whit path to policy file? Probably yes 

//Command to use: sed -i 's/^[[:space:]]*AsBlackList: .*/ \ AsBlackList: [12, 32]/' test.yaml
//Command to use: sed -i 's/^[[:space:]]*IsdBlackList: .*/ \ IsdBlackList: [12, 32]/' test.yaml
package main

import (
	"net/http"
	"log"
	"os/exec"
	"encoding/json"
    "os"
)

//move to types.go?
type UpdatePolicyRequest struct {
    ASList []string `json:"as_list"`
    ISDList []string `json:"isd_list"`
}

var (
    policyFilePath = "/etc/scion/path-policy.yaml"
)

func updatePolicyASList(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        json.NewEncoder(w).Encode(APIResponse{
            Status:  "error",
            Message: "Method not allowed - use POST",
        })
        return
    }

    var req UpdatePolicyRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(APIResponse{
            Status:  "error",
            Message: "Invalid JSON body",
        })
        return
    }

    // Format AS list for sed, e.g. [AS12, AS15]
    asListStr := "["
    for i, as := range req.ASList {
        if i > 0 {
            asListStr += ", "
        }
        asListStr += as
    }
    asListStr += "]"

    // Run sed command to update the file
    configFile := policyFilePath
    sedCmd := "sed"
    sedArgs := []string{
        "-i",
        "s/^[[:space:]]*AsBlackList: .*/  AsBlackList: " + asListStr + "/",
        configFile,
    }
    cmd := exec.Command(sedCmd, sedArgs...)
    output, err := cmd.CombinedOutput()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(APIResponse{
            Status:  "error",
            Message: "Failed to update config: " + err.Error() + " Output: " + string(output),
        })
        return
    }
	log.Printf("Updated policy file %s with AS blacklist: %s", configFile, asListStr)

    json.NewEncoder(w).Encode(APIResponse{
        Status:  "success",
        Message: "Config updated to blacklist ASes: " + asListStr,
    })
}

func updatePolicyISDList(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        json.NewEncoder(w).Encode(APIResponse{
            Status:  "error",
            Message: "Method not allowed - use POST",
        })
        return
    }

    var req UpdatePolicyRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(APIResponse{
            Status:  "error",
            Message: "Invalid JSON body",
        })
        return
    }

    // Format ISD list for sed
    isdListStr := "["
    for i, isd := range req.ISDList {
        if i > 0 {
            isdListStr += ", "
        }
        isdListStr += isd
    }
    isdListStr += "]"

    // Run sed command to update the file
    configFile := policyFilePath
    sedCmd := "sed"
    sedArgs := []string{
        "-i",
        "s/^[[:space:]]*IsdBlackList: .*/  IsdBlackList: " + isdListStr + "/",
        configFile,
    }
    cmd := exec.Command(sedCmd, sedArgs...)
    output, err := cmd.CombinedOutput()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(APIResponse{
            Status:  "error",
            Message: "Failed to update config: " + err.Error() + " Output: " + string(output),
        })
        return
    }
	log.Printf("Updated policy file %s with ISD blacklist: %s", configFile, isdListStr)

    json.NewEncoder(w).Encode(APIResponse{
        Status:  "success",
        Message: "Config updated to blacklist ISDs: " + isdListStr,
    })
}

//return path-policy file
func getPolicyFiles(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    if r.Method != http.MethodGet {
        w.WriteHeader(http.StatusMethodNotAllowed)
        json.NewEncoder(w).Encode(APIResponse{
            Status:  "error",
            Message: "Method not allowed - use GET",
        })
        return
    }

    data, err := os.ReadFile(policyFilePath)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(APIResponse{
            Status:  "error",
            Message: "Failed to read policy file: " + err.Error(),
        })
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Write(data)   
}