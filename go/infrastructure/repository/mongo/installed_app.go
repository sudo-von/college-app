package mongo

import (
	"errors"
	"time"

	"freelancer/college-app/go/entity"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type InstalledAppModel struct {
	ID               bson.ObjectId `bson:"_id"`
	User             BasicUser     `bson:"user"`
	PackageName      string        `bson:"package_name"`
	VersionName      string        `bson:"version_name"`
	VersionCode      string        `bson:"version_code"`
	FirstInstallTime string        `bson:"first_install_time"`
	LastUpdateTime   string        `bson:"last_update_ime"`
	AppName          string        `bson:"app_name"`
	Icon             string        `bson:"icon"`
	ApkDIR           string        `bson:"apk_dir"`
	Size             string        `bson:"size"`
	BatteryLevel     string        `bson:"battery_level"`
	CreationDate     time.Time     `bson:"creation_date"`
}

func toEntityInstalledApp(installedApp InstalledAppModel) entity.InstalledApp {

	user := entity.BasicUser{
		ID:    installedApp.User.ID.Hex(),
		Name:  installedApp.User.Name,
		Email: installedApp.User.Email,
	}

	return entity.InstalledApp{
		ID:               installedApp.ID.Hex(),
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
		CreationDate:     installedApp.CreationDate,
	}
}

type InstalledAppPayloadModel struct {
	ID               bson.ObjectId `bson:"_id"`
	UserID           bson.ObjectId `bson:"user_id"`
	PackageName      string        `bson:"package_name"`
	VersionName      string        `bson:"version_name"`
	VersionCode      string        `bson:"version_code"`
	FirstInstallTime string        `bson:"first_install_time"`
	LastUpdateTime   string        `bson:"last_update_ime"`
	AppName          string        `bson:"app_name"`
	Icon             string        `bson:"icon"`
	ApkDIR           string        `bson:"apk_dir"`
	Size             string        `bson:"size"`
	BatteryLevel     string        `bson:"battery_level"`
	CreationDate     time.Time     `bson:"creation_date"`
}

func toInstalledAppPayloadModel(installedApp entity.InstalledAppPayload) InstalledAppPayloadModel {

	var installedAppID bson.ObjectId
	if installedApp.ID != "" {
		installedAppID = bson.ObjectIdHex(installedApp.ID)
	} else {
		installedAppID = bson.NewObjectId()
	}

	var userID bson.ObjectId
	if installedApp.UserID != "" {
		userID = bson.ObjectIdHex(installedApp.UserID)
	} else {
		userID = bson.NewObjectId()
	}

	return InstalledAppPayloadModel{
		ID:               installedAppID,
		UserID:           userID,
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
		CreationDate:     installedApp.CreationDate,
	}
}

type InstalledAppRepository struct {
	Session      *mgo.Session
	DatabaseName string
}

func NewInstalledAppRepository(repository *Repository) *InstalledAppRepository {
	return &InstalledAppRepository{
		Session:      repository.Session,
		DatabaseName: repository.DatabaseName,
	}
}

func (r *InstalledAppRepository) GetInstalledAppByName(userID, appName string) (*entity.InstalledApp, error) {

	if !bson.IsObjectIdHex(userID) {
		return nil, errors.New("given user_id is not a valid hex")
	}

	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("installed_apps")

	currentDate := time.Now()
	initialDate := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 0, 0, time.Local)
	searchQuery := bson.M{
		"user.status": entity.ActiveStatus,
		"user_id":     bson.ObjectIdHex(userID),
		"app_name":    appName,
		"creation_date": bson.M{
			"$gte": initialDate,
			"$lt":  initialDate.AddDate(0, 0, 1),
		},
	}

	pipes := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "users",
				"localField":   "user_id",
				"foreignField": "_id",
				"as":           "user",
			},
		},
		{"$unwind": "$user"},
		{"$match": searchQuery},
	}

	var installedAppM InstalledAppModel
	pipe := com.Pipe(pipes)
	err := pipe.One(&installedAppM)
	if err != nil {
		return nil, err
	}

	installedApp := toEntityInstalledApp(installedAppM)
	return &installedApp, nil
}

func (r *InstalledAppRepository) GetInstalledAppsByUserID(userID string) ([]entity.InstalledApp, *int, error) {

	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("installed_apps")

	searchQuery := bson.M{
		"user.status": entity.ActiveStatus,
		"user_id":     bson.ObjectIdHex(userID),
	}

	currentDate := time.Now()
	initialDate := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 0, 0, time.Local)
	searchQuery["creation_date"] = bson.M{
		"$gte": initialDate,
		"$lt":  initialDate.AddDate(0, 0, 1),
	}

	pipes := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "users",
				"localField":   "user_id",
				"foreignField": "_id",
				"as":           "user",
			},
		},
		{"$unwind": "$user"},
		{"$match": searchQuery},
	}

	var installedAppsM []InstalledAppModel
	pipe := com.Pipe(pipes)
	err := pipe.All(&installedAppsM)
	if err != nil {
		return nil, nil, err
	}
	total := len(installedAppsM)
	installedApss := make([]entity.InstalledApp, 0)
	for _, installedApp := range installedAppsM {
		installedApss = append(installedApss, toEntityInstalledApp(installedApp))
	}

	return installedApss, &total, nil
}

func (r *InstalledAppRepository) CreateInstalledApp(newInstalledApp entity.InstalledAppPayload) error {

	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("installed_apps")

	installedAppM := toInstalledAppPayloadModel(newInstalledApp)
	err := com.Insert(&installedAppM)
	if err != nil {
		return err
	}

	return nil
}
