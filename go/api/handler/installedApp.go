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

type InstalledAppService interface {
	GetInstalledAppByName(userID, appName string) (*entity.InstalledApp, error)
	GetInstalledAppsByUserID(userID string) ([]entity.InstalledApp, *int, error)
	CreateInstalledApp(newInstalledApp entity.InstalledAppPayload) error
}

type InstalledAppController struct {
	InstalledAppService InstalledAppService
	AuthService         func(http.Handler) http.Handler
}

func NewInstalledAppController(installedAppService InstalledAppService, authService func(http.Handler) http.Handler) *InstalledAppController {
	return &InstalledAppController{
		InstalledAppService: installedAppService,
		AuthService:         authService,
	}
}

func (c *InstalledAppController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(c.AuthService)
	r.Get("/", c.List)
	r.Post("/", c.Create)
	return r
}

// @tags installed-apps
// @summary List installed apps.
// @description List installed apps.
// @security BearerJWT
// @id list-installed-apps
// @produce json
// @success 200 {object} presenter.InstalledAppsList
// @router /installed-apps [get]
func (c *InstalledAppController) List(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(ContextKeyUserID).(string)
	if !ok {
		err := errors.New("user not in context")
		CheckError(err, w, r)
		return
	}

	list, total, err := c.InstalledAppService.GetInstalledAppsByUserID(userID)
	if err != nil {
		CheckError(err, w, r)
		return
	}

	res := presenter.InstalledAppsList{
		Total:         *total,
		InstalledApps: make([]presenter.InstalledAppResponse, 0, len(list)),
	}

	for _, installedApp := range list {
		res.InstalledApps = append(res.InstalledApps, presenter.ToInstalledAppPresenter(installedApp))
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, &res)
}

// @tags installed-apps
// @summary Create installed app.
// @description Create installed app.
// @security BearerJWT
// @id create-installed-app
// @success 201
// @param payload body presenter.InstalledAppPayload true "Installed app that wants to be stored."
// @produce json
// @router /installed-apps [post]
func (c *InstalledAppController) Create(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(ContextKeyUserID).(string)
	if !ok {
		err := errors.New("user not in context")
		CheckError(err, w, r)
		return
	}

	var data presenter.InstalledAppPayload
	if err := render.Bind(r, &data); err != nil {
		CheckError(err, w, r)
		return
	}

	newInstalledApp := entity.InstalledAppPayload{
		UserID:           userID,
		PackageName:      data.PackageName,
		VersionName:      data.VersionName,
		VersionCode:      data.VersionCode,
		FirstInstallTime: data.FirstInstallTime,
		LastUpdateTime:   data.LastUpdateTime,
		AppName:          data.AppName,
		Icon:             data.Icon,
		ApkDIR:           data.ApkDIR,
		Size:             data.Size,
		BatteryLevel:     data.BatteryLevel,
		CreationDate:     time.Now().In(time.Local),
	}

	err := c.InstalledAppService.CreateInstalledApp(newInstalledApp)
	if err != nil {
		CheckError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
