package device

import (
	"freelancer/college-app/go/entity"
)

type DeviceWriterMock struct{}

func (a DeviceWriterMock) CreateDevice(device entity.DevicePayload) error {
	return nil
}

type DeviceRepositoryMock struct {
	DeviceWriterMock
}
