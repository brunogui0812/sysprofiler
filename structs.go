package sysprofiler

import "time"

//HardwareStruct hardware overview
type HardwareStruct struct {
	BootRomVersion string `mapstructure:"boot_rom_version"`
	CPUType        string `mapstructure:"cpu_type"`
	CPUSpeed       string `mapstructure:"current_processor_speed"`
	CPUCores       int    `mapstructure:"number_processors"`
	CPUNumber      int    `mapstructure:"packages"`
	L2CacheCore    string `mapstructure:"l2_cache_core"`
	L3Cache        string `mapstructure:"l3_cache"`
	MachineModel   string `mapstructure:"machine_model"`
	MachineName    string `mapstructure:"machine_name"`
	PhysicalMemory string `mapstructure:"physical_memory"`
	PlatformUUID   string `mapstructure:"platform_UUID"`
	Serial         string `mapstructure:"serial_number"`
}

//MonitorsStruct data about system monitors
type MonitorsStruct struct {
	Name         string `mapstructure:"_name"`
	Resolution   string `mapstructure:"spdisplays_resolution"`
	SerialNumber string `mapstructure:"spdisplays_display-serial-number"`
}

//DisplaysStruct data about system displays
type DisplaysStruct struct {
	Model    string           `mapstructure:"sppci_model"`
	Vendor   string           `mapstructure:"spdisplays_vendor"`
	VRAM     string           `mapstructure:"_spdisplays_vram"`
	Monitors []MonitorsStruct `mapstructure:"spdisplays_ndrvs"`
}

//AudioStruct data about system audio devices
type AudioStruct struct {
	Items []struct {
		Name         string `mapstructure:"_name"`
		OutputDevice int    `mapstructure:"coreaudio_device_output"`
	} `mapstructure:"_items"`
}

//MemoryStruct data about system memory
type MemoryStruct struct {
	Items []struct {
		Name         string `mapstructure:"_name"`
		Manufacturer string `mapstructure:"dimm_manufacturer"`
		PartNumber   string `mapstructure:"dimm_part_number"`
		SerialNumber string `mapstructure:"dimm_serial_number"`
		Size         string `mapstructure:"dimm_size"`
		Speed        string `mapstructure:"dimm_speed"`
		Status       string `mapstructure:"dimm_status"`
		Type         string `mapstructure:"dimm_type"`
	} `mapstructure:"_items"`
}

//OSStruct Operating system overview data
type OSStruct struct {
	BootMode        string `mapstructure:"boot_mode"`
	BootVolume      string `mapstructure:"boot_volume"`
	KernelVersion   string `mapstructure:"kernel_version"`
	HostName        string `mapstructure:"local_host_name"`
	Version         string `mapstructure:"os_version"`
	SecureVM        string `mapstructure:"secure_vm"`
	SystemIntegrity string `mapstructure:"system_integrity"`
	Uptime          string `mapstructure:"uptime"`
	Username        string `mapstructure:"user_name"`
}

//ApplicationsStruct infor about installed Apps
type ApplicationsStruct struct {
	Name         string    `mapstructure:"_name"`
	ArchKind     string    `mapstructure:"arch_kind"`
	LastModified time.Time `mapstructure:"lastModified"`
	ObtainedFrom string    `mapstructure:"obtained_from"`
	Path         string    `mapstructure:"path"`
	Version      string    `mapstructure:"version"`
	SignedBy     []string  `mapstructure:"signed_by"`
}

//UpdateStruct Software Update list
type UpdateStruct struct {
	Title       string `json:"title"`
	Version     string `json:"version"`
	Size        string `json:"size"`
	Recommended bool   `json:"recommended"`
	Action      string `json:"action"`
}

//MainStruct data about the whole system
type MainStruct struct {
	DataType string        `plist:"_dataType"`
	Items    []interface{} `plist:"_items"`
}
