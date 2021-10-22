package presenter

import (
	"freelancer/college-app/go/entity"
	"net/http"
)

var (
	ErrInstalledAppAlreadyRegistered = "INSTALLED_APP_ALREADY_REGISTERED"
)

type InstalledAppsList struct {
	Total         int                    `json:"total" example:"1"`
	InstalledApps []InstalledAppResponse `json:"results"`
}

func (ial *InstalledAppsList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type InstalledAppResponse struct {
	ID               string    `json:"id" example:"613aab578a6ef50007e622be"`
	PackageName      string    `json:"package_name" example:"Fake package name"`
	VersionName      string    `json:"version_name" example:"Fake version name"`
	VersionCode      string    `json:"version_code" example:"Fake version code"`
	FirstInstallTime string    `json:"first_install_time" example:"Fake first install time"`
	LastUpdateTime   string    `json:"last_update_ime" example:"Fake last update time"`
	AppName          string    `json:"app_name" example:"Fake app name"`
	Icon             string    `json:"icon" example:"Fake icon"`
	ApkDIR           string    `json:"apk_dir" example:"Fake apk dir"`
	Size             string    `json:"size" example:"Fake size"`
	BatteryLevel     string    `json:"battery_level" example:"Fake battery level"`
	User             BasicUser `json:"user"`
}

func (iar *InstalledAppResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ToInstalledAppPresenter(installedApp entity.InstalledApp) InstalledAppResponse {

	user := BasicUser{
		ID:    installedApp.User.ID,
		Name:  installedApp.User.Name,
		Email: installedApp.User.Email,
	}

	return InstalledAppResponse{
		ID:               installedApp.ID,
		User:             user,
		PackageName:      installedApp.PackageName,
		VersionName:      installedApp.VersionName,
		VersionCode:      installedApp.VersionCode,
		FirstInstallTime: installedApp.FirstInstallTime,
		LastUpdateTime:   installedApp.LastUpdateTime,
		AppName:          installedApp.AppName,
		Icon:             installedApp.Icon,
		ApkDIR:           installedApp.ApkDIR,
		Size:             installedApp.Size,
		BatteryLevel:     installedApp.BatteryLevel,
	}
}

type InstalledAppPayload struct {
	PackageName      string `json:"package_name" example:"Fake package name"`
	VersionName      string `json:"version_name" example:"Fake version name"`
	VersionCode      string `json:"version_code" example:"Fake version code"`
	FirstInstallTime string `json:"first_install_time" example:"Fake first install time"`
	LastUpdateTime   string `json:"last_update_ime" example:"Fake last update time"`
	AppName          string `json:"app_name" example:"Fake app name"`
	Icon             string `json:"icon" example:"Fake icon"`
	ApkDIR           string `json:"apk_dir" example:"Fake apk dir"`
	Size             string `json:"size" example:"Fake size"`
	BatteryLevel     string `json:"battery_level" example:"Fake battery level"`
}

func (iap *InstalledAppPayload) validate() (err error) {
	return
}

func (iap *InstalledAppPayload) Bind(r *http.Request) error {
	if err := iap.validate(); err != nil {
		return err
	}
	return nil
}
