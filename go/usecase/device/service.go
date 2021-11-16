package device

import (
	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/usecase/user"
)

type Service struct {
	deviceRepository DeviceRepository
	userRepository   UserRepository
}

func NewService(deviceRepository DeviceRepository, userRepository user.UserRepository) *Service {
	return &Service{
		deviceRepository,
		userRepository,
	}
}

func (s Service) CreateDevice(device entity.DevicePayload) error {

	_, err := s.userRepository.GetUserByID(device.UserID)
	if err != nil {
		if err.Error() == "not found" {
			return entity.NewErrorNotFound(err, presenter.ErrUserNotFound)
		}
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	err = s.deviceRepository.CreateDevice(device)
	if err != nil {
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}
	return nil
}
