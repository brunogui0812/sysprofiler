package sysprofiler

import (
	"github.com/mitchellh/mapstructure"
)

const (
	//HardwareDT hardware data type
	HardwareDT = "SPHardwareDataType"
	//AudioDT audio data type
	AudioDT = "SPAudioDataType"
	//MemoryDT memory data type
	MemoryDT = "SPMemoryDataType"
	//DisplaysDT displays data type
	DisplaysDT = "SPDisplaysDataType"
)

//Hardware get hardware overview data
func Hardware() ([]HardwareStruct, error) {
	var data []MainStruct
	var hardwareData []HardwareStruct
	var err error
	if data, err = Exec(HardwareDT); err != nil {
		return nil, err
	}

	for _, dataItem := range data {
		if dataItem.DataType == HardwareDT {
			for _, item := range dataItem.Items {
				var decoded HardwareStruct
				if err := mapstructure.Decode(item, &decoded); err != nil {
					return nil, err
				}
				hardwareData = append(hardwareData, decoded)
			}
		}
	}

	return hardwareData, nil
}

//Memory get information about system memory
func Memory() ([]MemoryStruct, error) {
	var data []MainStruct
	var memoryData []MemoryStruct
	var err error
	if data, err = Exec(MemoryDT); err != nil {
		return nil, err
	}
	for _, dataItem := range data {
		if dataItem.DataType == MemoryDT {
			for _, item := range dataItem.Items {
				var decoded MemoryStruct
				if err := mapstructure.Decode(item, &decoded); err != nil {
					return nil, err
				}
				memoryData = append(memoryData, decoded)
			}
		}
	}

	return memoryData, nil
}

//Displays get information about system displays
func Displays() ([]DisplaysStruct, error) {
	var data []MainStruct
	var displayData []DisplaysStruct
	var err error
	if data, err = Exec(DisplaysDT); err != nil {
		return nil, err
	}

	for _, dataItem := range data {
		if dataItem.DataType == DisplaysDT {
			for _, item := range dataItem.Items {
				var decoded DisplaysStruct
				if err := mapstructure.Decode(item, &decoded); err != nil {
					return nil, err
				}
				displayData = append(displayData, decoded)
			}
		}
	}

	return displayData, nil
}

//Audio get information about system audio devices
func Audio() ([]AudioStruct, error) {
	var data []MainStruct
	var audioData []AudioStruct
	var err error
	if data, err = Exec(AudioDT); err != nil {
		return nil, err
	}

	for _, dataItem := range data {
		if dataItem.DataType == AudioDT {
			for _, item := range dataItem.Items {
				var decoded AudioStruct
				if err := mapstructure.Decode(item, &decoded); err != nil {
					return nil, err
				}
				audioData = append(audioData, decoded)
			}
		}
	}

	return audioData, nil
}
