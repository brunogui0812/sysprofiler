package sysprofiler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

//Exec get information about MacOs system using system_profiler command
func Exec(dataType ...string) (MainStruct, error) {
	var cmdOut, cmdErr bytes.Buffer
	var err error
	var systemProfilerData MainStruct

	cmd := exec.Command("system_profiler", append([]string{"-json"}, dataType...)...)
	cmd.Stdout = &cmdOut
	cmd.Stderr = &cmdErr

	if err := cmd.Run(); err != nil {
		return MainStruct{}, fmt.Errorf("%s %s %s", cmdOut.String(), cmdErr.String(), err)
	}

	if err = json.Unmarshal(cmdOut.Bytes(), &systemProfilerData); err != nil {
		return MainStruct{}, err
	}

	return systemProfilerData, nil

}
