// Contains code to configure nodes
//
// -- Start, Stop scion services
// -- Edit Path Policy
// -- IPTables (?)
// -- Chaos (?)

// KIM DA CHASCH DRA SCHAFFE

//Questions:
	//Upload default policy file which is then edited? Or only upload file when needed?
	//Upload default config file whit path to policy file? Probably yes 

//Command to use: sed -i 's/^[[:space:]]*AsBlackList: .*/ \ AsBlackList: [12, 32]/' test.yaml
package main

import (
	"net/http"
	"log"
	"os/exec"
	"encoding/json"
)

type UpdatePolicyRequest struct {
    ASList []string `json:"as_list"`
}

func updatePolicyHandler(w http.ResponseWriter, r *http.Request) {
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

    // Run sed command to update the config file
    configFile := "/etc/scion/policy.yaml" // Change as needed
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