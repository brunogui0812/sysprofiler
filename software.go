package sysprofiler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/spakin/awk"
)

const (
	//OSDT Operating system overview data type
	OSDT = "SPSoftwareDataType"
	//ApplicationsDT Installed apps
	ApplicationsDT = "SPApplicationsDataType"
)

//OS get operating system overview data
func OS() ([]OSStruct, error) {
	var data MainStruct
	var err error
	if data, err = Exec(OSDT); err != nil {
		return nil, err
	}

	return data.OS, nil
}

//Applications get installed apps on the system
func Applications() ([]ApplicationsStruct, error) {
	var data MainStruct
	var err error
	if data, err = Exec(ApplicationsDT); err != nil {
		return nil, err
	}

	return data.Applications, nil
}

//Updates get software updates available
func Updates() ([]UpdateStruct, error) {
	var err error
	var data []UpdateStruct
	var updateDataJSON []byte
	var updateDataMapArray []map[string]interface{}
	var cmdOut, cmdErr bytes.Buffer
	cmd := exec.Command("softwareupdate", "-l", "--no-scan")
	cmd.Stdout = &cmdOut
	cmd.Stderr = &cmdErr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("%s %s %s", cmdOut.String(), cmdErr.String(), err)
	}

	updatesAwk := awk.NewScript()
	updatesAwk.SetFS(",")
	updatesAwk.AppendStmt(func(s *awk.Script) bool {
		return strings.Contains(strings.ToLower(s.F(0).String()), "title")
	},
		func(s *awk.Script) {
			updateDataMap := make(map[string]interface{})
			for _, lineData := range s.FStrings() {
				splittedLineData := strings.Split(lineData, ":")
				if len(splittedLineData) > 1 {
					key, value := strings.TrimSpace(strings.ToLower(splittedLineData[0])), strings.TrimSpace(splittedLineData[1])
					if key == "recommended" {
						switch strings.ToLower(value) {
						case "yes":
							updateDataMap[key] = true
						default:
							updateDataMap[key] = false
						}
						continue
					}
					updateDataMap[key] = value
				}
			}
			updateDataMapArray = append(updateDataMapArray, updateDataMap)
		})
	updatesAwk.Run(bytes.NewReader(cmdOut.Bytes()))

	if updateDataJSON, err = json.Marshal(updateDataMapArray); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(updateDataJSON, &data); err != nil {
		return nil, err
	}

	return data, nil
}
