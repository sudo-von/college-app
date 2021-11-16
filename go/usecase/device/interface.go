package device

import (
	"freelancer/college-app/go/entity"
)

type UserReader interface {
	GetUserByID(userID string) (*entity.User, error)
}

type UserRepository interface {
	UserReader
}

type DeviceWriter interface {
	CreateDevice(device entity.DevicePayload) error
}

type DeviceRepository interface {
	DeviceWriter
}
