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
func Hardware() (HardwareStruct, error) {
	var data []MainStruct
	var hardwareData HardwareStruct
	var err error
	if data, err = Exec(HardwareDT); err != nil {
		return HardwareStruct{}, err
	}

	for _, item := range data {
		if item.DataType == HardwareDT {
			if err := mapstructure.Decode(item.Items[0], &hardwareData); err != nil {
				return HardwareStruct{}, err
			}
			return hardwareData, nil
		}
	}

	return HardwareStruct{}, nil
}

//Memory get information about system memory
func Memory() (MemoryStruct, error) {
	var data []MainStruct
	var memoryData MemoryStruct
	var err error
	if data, err = Exec(MemoryDT); err != nil {
		return MemoryStruct{}, err
	}
	for _, item := range data {
		if item.DataType == MemoryDT {
			if err := mapstructure.Decode(item.Items[0], &memoryData); err != nil {
				return MemoryStruct{}, err
			}
			return memoryData, nil
		}
	}

	return MemoryStruct{}, nil
}

//Displays get information about system displays
func Displays() (DisplaysStruct, error) {
	var data []MainStruct
	var displayData DisplaysStruct
	var err error
	if data, err = Exec(DisplaysDT); err != nil {
		return DisplaysStruct{}, err
	}

	for _, item := range data {
		if item.DataType == DisplaysDT {
			if err := mapstructure.Decode(item.Items[0], &displayData); err != nil {
				return DisplaysStruct{}, err
			}
			return displayData, nil
		}
	}

	return DisplaysStruct{}, nil
}

//Audio get information about system audio devices
func Audio() (AudioStruct, error) {
	var data []MainStruct
	var audioData AudioStruct
	var err error
	if data, err = Exec(AudioDT); err != nil {
		return AudioStruct{}, err
	}

	for _, item := range data {
		if item.DataType == AudioDT {
			if err := mapstructure.Decode(item.Items[0], &audioData); err != nil {
				return AudioStruct{}, err
			}
			return audioData, nil
		}
	}

	return AudioStruct{}, nil
}
