package mongo

import (
	"time"

	"freelancer/college-app/go/entity"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Device struct {
	DeviceName                string   `bson:"device_name"`
	IsDevice                  bool     `bson:"is_device"`
	Brand                     string   `bson:"brand"`
	Manufacturer              string   `bson:"manufacturer"`
	ModelName                 string   `bson:"model_name"`
	ModelID                   string   `bson:"model_id"`
	DesignName                string   `bson:"design_name"`
	ProductName               string   `bson:"product_name"`
	DeviceYearClass           string   `bson:"device_year_class"`
	TotalMemory               string   `bson:"total_memory"`
	OSName                    string   `bson:"os_name"`
	OSVersion                 string   `bson:"os_version"`
	OSBuildID                 string   `bson:"os_build_id"`
	OSInternalBuildID         string   `bson:"os_internal_build_id"`
	OSBuildFingerprint        string   `bson:"os_build_fingerprint"`
	PlatformAPILevel          string   `bson:"platform_api_level"`
	SupportedCPUArchitectures []string `bson:"supported_cpu_architectures"`
}

type DevicePayloadModel struct {
	ID           bson.ObjectId `bson:"_id"`
	UserID       bson.ObjectId `bson:"user_id"`
	Device       Device        `bson:"device"`
	Battery      Battery       `bson:"battery"`
	CreationDate time.Time     `bson:"creation_date"`
}

func toDevicePayloadModel(device entity.DevicePayload) DevicePayloadModel {

	var deviceID bson.ObjectId
	if device.ID != "" {
		deviceID = bson.ObjectIdHex(device.ID)
	} else {
		deviceID = bson.NewObjectId()
	}

	var userID bson.ObjectId
	if device.UserID != "" {
		userID = bson.ObjectIdHex(device.UserID)
	} else {
		userID = bson.NewObjectId()
	}

	return DevicePayloadModel{
		ID:     deviceID,
		UserID: userID,
		Device: Device{
			DeviceName:                device.Device.DeviceName,
			IsDevice:                  device.Device.IsDevice,
			Brand:                     device.Device.Brand,
			Manufacturer:              device.Device.Manufacturer,
			ModelName:                 device.Device.ModelName,
			ModelID:                   device.Device.ModelID,
			DesignName:                device.Device.DesignName,
			ProductName:               device.Device.ProductName,
			DeviceYearClass:           device.Device.DeviceYearClass,
			TotalMemory:               device.Device.TotalMemory,
			OSName:                    device.Device.OSName,
			OSVersion:                 device.Device.OSVersion,
			OSBuildID:                 device.Device.OSBuildID,
			OSInternalBuildID:         device.Device.OSInternalBuildID,
			OSBuildFingerprint:        device.Device.OSBuildFingerprint,
			PlatformAPILevel:          device.Device.PlatformAPILevel,
			SupportedCPUArchitectures: device.Device.SupportedCPUArchitectures,
		},
		Battery: Battery{
			BatteryLevel: device.Battery.BatteryLevel,
			LowPowerMode: device.Battery.LowPowerMode,
			BatteryState: device.Battery.BatteryState,
		},
		CreationDate: device.CreationDate,
	}
}

type DeviceRepository struct {
	Session      *mgo.Session
	DatabaseName string
}

func NewDeviceRepository(repository *Repository) *DeviceRepository {
	return &DeviceRepository{
		Session:      repository.Session,
		DatabaseName: repository.DatabaseName,
	}
}

func (r *DeviceRepository) CreateDevice(device entity.DevicePayload) error {

	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("devices")

	installedAppM := toDevicePayloadModel(device)
	err := com.Insert(&installedAppM)
	if err != nil {
		return err
	}

	return nil
}
