package entity

import (
	"time"
)

type Device struct {
	DeviceName                string
	IsDevice                  bool
	Brand                     string
	Manufacturer              string
	ModelName                 string
	ModelID                   string
	DesignName                string
	ProductName               string
	DeviceYearClass           string
	TotalMemory               string
	OSName                    string
	OSVersion                 string
	OSBuildID                 string
	OSInternalBuildID         string
	OSBuildFingerprint        string
	PlatformAPILevel          string
	SupportedCPUArchitectures []string
}

type Battery struct {
	BatteryLevel float64
	LowPowerMode bool
	BatteryState string
}

type DevicePayload struct {
	ID           string
	UserID       string
	Device       Device
	Battery      Battery
	CreationDate time.Time
}
