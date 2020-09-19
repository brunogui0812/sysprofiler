package sysprofiler

const (
	//OSDT Operating system overview data type
	OSDT = "SPSoftwareDataType"
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
