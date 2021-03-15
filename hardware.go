package sysprofiler

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
	var data MainStruct
	var err error
	if data, err = Exec(HardwareDT); err != nil {
		return nil, err
	}

	return data.Hardware, nil
}

//Memory get information about system memory
func Memory() ([]MemoryStruct, error) {
	var data MainStruct
	var err error
	if data, err = Exec(MemoryDT); err != nil {
		return nil, err
	}

	return data.Memory, nil
}

//Displays get information about system displays
func Displays() ([]DisplaysStruct, error) {
	var data MainStruct
	var err error
	if data, err = Exec(DisplaysDT); err != nil {
		return nil, err
	}

	return data.Displays, nil
}

//Audio get information about system audio devices
func Audio() ([]AudioStruct, error) {
	var data MainStruct
	var err error
	if data, err = Exec(AudioDT); err != nil {
		return nil, err
	}

	return data.Audio, nil
}
