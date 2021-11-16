package handler

import (
	"errors"
	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/entity"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type DeviceService interface {
	CreateDevice(device entity.DevicePayload) error
}

type DeviceController struct {
	DeviceService DeviceService
	AuthService   func(http.Handler) http.Handler
}

func NewDeviceController(deviceService DeviceService, authService func(http.Handler) http.Handler) *DeviceController {
	return &DeviceController{
		DeviceService: deviceService,
		AuthService:   authService,
	}
}

func (c *DeviceController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(c.AuthService)
	r.Post("/", c.Create)
	return r
}

// @tags create-device
// @summary Create device.
// @description Store device information.
// @security BearerJWT
// @id create-device-information
// @success 201
// @param payload body presenter.DevicePayload true "Device that wants to be stored."
// @produce json
// @router /devices [post]
func (c *DeviceController) Create(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(ContextKeyUserID).(string)
	if !ok {
		err := errors.New("user not in context")
		CheckError(err, w, r)
		return
	}

	var data presenter.DevicePayload
	if err := render.Bind(r, &data); err != nil {
		CheckError(err, w, r)
		return
	}

	device := entity.DevicePayload{
		UserID: userID,
		Device: entity.Device{
			DeviceName:                data.Device.DeviceName,
			IsDevice:                  data.Device.IsDevice,
			Brand:                     data.Device.Brand,
			Manufacturer:              data.Device.Manufacturer,
			ModelName:                 data.Device.ModelName,
			ModelID:                   data.Device.ModelID,
			DesignName:                data.Device.DesignName,
			ProductName:               data.Device.ProductName,
			DeviceYearClass:           data.Device.DeviceYearClass,
			TotalMemory:               data.Device.TotalMemory,
			OSName:                    data.Device.OSName,
			OSVersion:                 data.Device.OSVersion,
			OSBuildID:                 data.Device.OSBuildID,
			OSInternalBuildID:         data.Device.OSInternalBuildID,
			OSBuildFingerprint:        data.Device.OSBuildFingerprint,
			PlatformAPILevel:          data.Device.PlatformAPILevel,
			SupportedCPUArchitectures: data.Device.SupportedCPUArchitectures,
		},
		Battery: entity.Battery{
			BatteryLevel: data.Battery.BatteryLevel,
			LowPowerMode: data.Battery.LowPowerMode,
			BatteryState: data.Battery.BatteryState,
		},
		CreationDate: time.Now().In(time.Local),
	}

	err := c.DeviceService.CreateDevice(device)
	if err != nil {
		CheckError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
