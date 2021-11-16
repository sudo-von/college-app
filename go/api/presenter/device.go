package presenter

import (
	"net/http"
)

type DevicePayload struct {
	Device struct {
		DeviceName                string   `json:"device_name" example:"Fake device name"`
		IsDevice                  bool     `json:"is_device"`
		Brand                     string   `json:"brand" example:"Fake brand"`
		Manufacturer              string   `json:"manufacturer" example:"Fake manufacturer"`
		ModelName                 string   `json:"model_name" example:"Fake model name"`
		ModelID                   string   `json:"model_id" example:"Fake model id"`
		DesignName                string   `json:"design_name" example:"Fake design name"`
		ProductName               string   `json:"product_name" example:"Fake product name"`
		DeviceYearClass           string   `json:"device_year_class" example:"Fake device year class"`
		TotalMemory               string   `json:"total_memory" example:"Fake total memory"`
		OSName                    string   `json:"os_name" example:"Fake os name"`
		OSVersion                 string   `json:"os_version" example:"Fake os version"`
		OSBuildID                 string   `json:"os_build_id" example:"Fake os build id"`
		OSInternalBuildID         string   `json:"os_internal_build_id" example:"Fake os internal build id"`
		OSBuildFingerprint        string   `json:"os_build_fingerprint" example:"Fake os build fingerprint"`
		PlatformAPILevel          string   `json:"platform_api_level" example:"Fake platform api level"`
		SupportedCPUArchitectures []string `json:"supported_cpu_architectures"`
	} `json:"device"`
	Battery struct {
		BatteryLevel float64 `json:"battery_level" example:"70.5"`
		LowPowerMode bool    `json:"low_power_mode"`
		BatteryState string  `json:"battery_state" example:"Fake battery state"`
	} `json:"battery"`
}

func (dp *DevicePayload) validate() (err error) {
	return
}

func (dp *DevicePayload) Bind(r *http.Request) error {
	if err := dp.validate(); err != nil {
		return err
	}
	return nil
}
