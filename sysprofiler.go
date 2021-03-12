package sysprofiler

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/groob/plist"
)

//Exec get information about MacOs system using system_profiler command
func Exec(dataType ...string) ([]MainStruct, error) {
	var cmdOut, cmdErr bytes.Buffer
	var err error
	var systemProfilerData []MainStruct

	cmd := exec.Command("system_profiler", append([]string{"-xml"}, dataType...)...)
	cmd.Stdout = &cmdOut
	cmd.Stderr = &cmdErr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("%s %s %s", cmdOut.String(), cmdErr.String(), err)
	}

	if err = plist.NewXMLDecoder(bytes.NewReader(cmdOut.Bytes())).Decode(&systemProfilerData); err != nil {
		return nil, err
	}

	return systemProfilerData, nil

}
