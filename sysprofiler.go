package sysprofiler

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/groob/plist"
	"github.com/mitchellh/mapstructure"
)

//Exec get information about MacOs system using system_profiler command
func Exec(dataType ...string) (MainStruct, error) {
	var cmdOut, cmdErr bytes.Buffer
	var err error
	var rawData []RawStruct
	var systemProfilerData MainStruct

	cmd := exec.Command("system_profiler", append([]string{"-xml"}, dataType...)...)
	cmd.Stdout = &cmdOut
	cmd.Stderr = &cmdErr

	if err := cmd.Run(); err != nil {
		return MainStruct{}, fmt.Errorf("%s %s %s", cmdOut.String(), cmdErr.String(), err)
	}

	if err = plist.NewXMLDecoder(bytes.NewReader(cmdOut.Bytes())).Decode(&rawData); err != nil {
		return MainStruct{}, err
	}

	for _, data := range rawData {
		switch data.DataType {
		case HardwareDT:
			if systemProfilerData.Hardware, err = decodeHwData(data.Items); err != nil {
				return MainStruct{}, err
			}
			break
		case DisplaysDT:
			if systemProfilerData.Displays, err = decodeDisplayData(data.Items); err != nil {
				return MainStruct{}, err
			}
			break
		case AudioDT:
			if systemProfilerData.Audio, err = decodeAudioData(data.Items); err != nil {
				return MainStruct{}, err
			}
			break
		case MemoryDT:
			if systemProfilerData.Memory, err = decodeMemoryData(data.Items); err != nil {
				return MainStruct{}, err
			}
			break
		case OSDT:
			if systemProfilerData.OS, err = decodeOSData(data.Items); err != nil {
				return MainStruct{}, err
			}
			break
		case ApplicationsDT:
			if systemProfilerData.Applications, err = decodeAppData(data.Items); err != nil {
				return MainStruct{}, err
			}
			break
		default:
			break
		}
	}

	return systemProfilerData, nil

}

func decodeHwData(items []interface{}) ([]HardwareStruct, error) {
	var decodedData []HardwareStruct
	for _, item := range items {
		var decoded HardwareStruct
		if err := mapstructure.Decode(item, &decoded); err != nil {
			return nil, err
		}
		decodedData = append(decodedData, decoded)
	}

	return decodedData, nil
}

func decodeDisplayData(items []interface{}) ([]DisplaysStruct, error) {
	var decodedData []DisplaysStruct
	for _, item := range items {
		var decoded DisplaysStruct
		if err := mapstructure.Decode(item, &decoded); err != nil {
			return nil, err
		}
		decodedData = append(decodedData, decoded)
	}

	return decodedData, nil
}

func decodeAudioData(items []interface{}) ([]AudioStruct, error) {
	var decodedData []AudioStruct
	for _, item := range items {
		var decoded AudioStruct
		if err := mapstructure.Decode(item, &decoded); err != nil {
			return nil, err
		}
		decodedData = append(decodedData, decoded)
	}

	return decodedData, nil
}

func decodeMemoryData(items []interface{}) ([]MemoryStruct, error) {
	var decodedData []MemoryStruct
	for _, item := range items {
		var decoded MemoryStruct
		if err := mapstructure.Decode(item, &decoded); err != nil {
			return nil, err
		}
		decodedData = append(decodedData, decoded)
	}

	return decodedData, nil
}

func decodeOSData(items []interface{}) ([]OSStruct, error) {
	var decodedData []OSStruct
	for _, item := range items {
		var decoded OSStruct
		if err := mapstructure.Decode(item, &decoded); err != nil {
			return nil, err
		}
		decodedData = append(decodedData, decoded)
	}

	return decodedData, nil
}

func decodeAppData(items []interface{}) ([]ApplicationsStruct, error) {
	var decodedData []ApplicationsStruct
	for _, item := range items {
		var decoded ApplicationsStruct
		if err := mapstructure.Decode(item, &decoded); err != nil {
			return nil, err
		}
		decodedData = append(decodedData, decoded)
	}

	return decodedData, nil
}
