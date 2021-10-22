package entity

import (
	"time"
)

type InstalledApp struct {
	ID               string
	User             BasicUser
	PackageName      string
	VersionName      string
	VersionCode      string
	FirstInstallTime string
	LastUpdateTime   string
	AppName          string
	Icon             string
	ApkDIR           string
	Size             string
	BatteryLevel     string
	CreationDate     time.Time
}

type InstalledAppPayload struct {
	ID               string
	UserID           string
	PackageName      string
	VersionName      string
	VersionCode      string
	FirstInstallTime string
	LastUpdateTime   string
	AppName          string
	Icon             string
	ApkDIR           string
	Size             string
	BatteryLevel     string
	CreationDate     time.Time
}
