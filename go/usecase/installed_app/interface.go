package installed_app

import (
	"freelancer/college-app/go/entity"
)

type UserReader interface {
	GetUserByID(userID string) (*entity.User, error)
}

type UserRepository interface {
	UserReader
}

type InstalledAppReader interface {
	GetInstalledAppsByUserID(userID string) ([]entity.InstalledApp, *int, error)
	GetInstalledAppByName(userID, appName string) (*entity.InstalledApp, error)
}

type InstalledAppWriter interface {
	CreateInstalledApp(newInstalledApp entity.InstalledAppPayload) error
}

type InstalledAppRepository interface {
	InstalledAppReader
	InstalledAppWriter
}
