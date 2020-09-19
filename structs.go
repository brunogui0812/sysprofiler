package sysprofiler

//HardwareStruct hardware overview
type HardwareStruct struct {
	BootRomVersion string `json:"boot_rom_version"`
	CPUType        string `json:"cpu_type"`
	CPUSpeed       string `json:"current_processor_speed"`
	CPUCores       int    `json:"number_processors"`
	CPUNumber      int    `json:"packages"`
	L2CacheCore    string `json:"l2_cache_core"`
	L3Cache        string `json:"l3_cache"`
	MachineModel   string `json:"machine_model"`
	MachineName    string `json:"machine_name"`
	PhysicalMemory string `json:"physical_memory"`
	PlatformUUID   string `json:"platform_UUID"`
	Serial         string `json:"serial_number"`
}

//MonitorsStruct data about system monitors
type MonitorsStruct struct {
	Name         string `json:"_name"`
	Resolution   string `json:"spdisplays_resolution"`
	SerialNumber string `json:"spdisplays_display-serial-number"`
}

//DisplaysStruct data about system displays
type DisplaysStruct struct {
	Model    string           `json:"sppci_model"`
	Vendor   string           `json:"spdisplays_vendor"`
	VRAM     string           `json:"_spdisplays_vram"`
	Monitors []MonitorsStruct `json:"spdisplays_ndrvs"`
}

//AudioStruct data about system audio devices
type AudioStruct struct {
	Items []struct {
		Name         string `json:"_name"`
		OutputDevice int    `json:"coreaudio_device_output"`
	} `json:"_items"`
}

//MemoryStruct data about system memory
type MemoryStruct struct {
	Items []struct {
		Name         string `json:"_name"`
		Manufacturer string `json:"dimm_manufacturer"`
		PartNumber   string `json:"dimm_part_number"`
		SerialNumber string `json:"dimm_serial_number"`
		Size         string `json:"dimm_size"`
		Speed        string `json:"dimm_speed"`
		Status       string `json:"dimm_status"`
		Type         string `json:"dimm_type"`
	} `json:"_items"`
}

//OSStruct Operating system overview data
type OSStruct struct {
	BootMode        string `json:"boot_mode"`
	BootVolume      string `json:"boot_volume"`
	KernelVersion   string `json:"kernel_version"`
	HostName        string `json:"local_host_name"`
	Version         string `json:"os_version"`
	SecureVM        string `json:"secure_vm"`
	SystemIntegrity string `json:"system_integrity"`
	Uptime          string `json:"uptime"`
	Username        string `json:"user_name"`
}

//MainStruct data about the whole system
type MainStruct struct {
	Hardware []HardwareStruct `json:"SPHardwareDataType"`
	Displays []DisplaysStruct `json:"SPDisplaysDataType"`
	Audio    []AudioStruct    `json:"SPAudioDataType"`
	Memory   []MemoryStruct   `json:"SPMemoryDataType"`
	OS       []OSStruct       `json:"SPSoftwareDataType"`
}
