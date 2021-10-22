package installed_app

import (
	"errors"
	"freelancer/college-app/go/api/presenter"
	"freelancer/college-app/go/entity"
	"freelancer/college-app/go/usecase/user"
)

type Service struct {
	installedAppRepository InstalledAppRepository
	userRepository         UserRepository
}

func NewService(installedAppRepository InstalledAppRepository, userRepository user.UserRepository) *Service {
	return &Service{
		installedAppRepository,
		userRepository,
	}
}

func (s Service) GetInstalledAppByName(userID, appName string) (*entity.InstalledApp, error) {

	_, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		if err.Error() == "not found" {
			return nil, entity.NewErrorNotFound(err, presenter.ErrUserNotFound)
		}
		return nil, entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	installedApp, err := s.installedAppRepository.GetInstalledAppByName(userID, appName)
	if err != nil {
		return nil, entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	return installedApp, nil
}

func (s Service) GetInstalledAppsByUserID(userID string) ([]entity.InstalledApp, *int, error) {

	_, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		if err.Error() == "not found" {
			return nil, nil, entity.NewErrorNotFound(err, presenter.ErrUserNotFound)
		}
		return nil, nil, entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	installedApps, total, err := s.installedAppRepository.GetInstalledAppsByUserID(userID)
	if err != nil {
		return nil, nil, entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	return installedApps, total, nil
}

func (s Service) CreateInstalledApp(newInstalledApp entity.InstalledAppPayload) error {

	_, err := s.userRepository.GetUserByID(newInstalledApp.UserID)
	if err != nil {
		if err.Error() == "not found" {
			return entity.NewErrorNotFound(err, presenter.ErrUserNotFound)
		}
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}

	_, err = s.installedAppRepository.GetInstalledAppByName(newInstalledApp.UserID, newInstalledApp.AppName)
	if err != nil && err.Error() != "not found" {
		return err
	} else if err == nil {
		return entity.NewErrorConflict(errors.New("installed app has already been registered today"), presenter.ErrInstalledAppAlreadyRegistered)
	}

	err = s.installedAppRepository.CreateInstalledApp(newInstalledApp)
	if err != nil {
		return entity.NewErrorInternalServer(err, presenter.ErrIntServError)
	}
	return nil
}
