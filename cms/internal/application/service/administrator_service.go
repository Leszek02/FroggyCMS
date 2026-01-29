package service

import (
	"fmt"
	"prestaprest/internal/domain/entity"
	"prestaprest/internal/infrastructure/auth"
	"prestaprest/internal/infrastructure/postgres"

	"gorm.io/gorm"
)

type AdministratorService struct {
	PostgresRepository postgres.PostgresRepository[entity.Administrator]
	passwordManager    *auth.PasswordHasher
}

func NewAdministratorService(conn *gorm.DB) AdministratorService {
	return AdministratorService{
		PostgresRepository: postgres.NewPostgresRepository[entity.Administrator](conn),
		passwordManager:    auth.NewPasswordHasher(),
	}
}

func (s *AdministratorService) GetAllAdministrators() ([]entity.Administrator, error) {
	allAdministrators, err := s.PostgresRepository.GetAll("Administrator_ID")
	if err != nil {
		return nil, err
	}
	return allAdministrators, nil
}

func (s *AdministratorService) GetAdministrator(administratorID uint) (*entity.Administrator, error) {
	administrator, err := s.PostgresRepository.Get("Administrator_ID", administratorID)
	fmt.Println("Am here (service)?")
	if err != nil {
		return nil, err
	}
	return administrator, nil
}

func (s *AdministratorService) CreateAdministrator(administrator *entity.Administrator) error {
	fmt.Println("Creating administrator: ", administrator)
	hashedPassword, err := s.passwordManager.HashPassword(administrator.Password)
	if err != nil {
		return err
	}
	fmt.Println("Hashed password: ", hashedPassword)
	administrator.Password = hashedPassword
	err = s.PostgresRepository.Create(administrator)
	if err != nil {
		return err
	}
	return nil
}

func (s *AdministratorService) UpdateAdministrator(administrator *entity.Administrator) error {
	err := s.PostgresRepository.Update(administrator)
	if err != nil {
		return err
	}
	return nil
}

func (s *AdministratorService) DeleteAdministrator(administrator_ID uint) error {
	err := s.PostgresRepository.Delete("Administrator_ID", uint64(administrator_ID))
	if err != nil {
		return err
	}
	return nil
}

func (s *AdministratorService) Authenticate(email, password string) (*entity.Administrator, error) {
	administrator, err := s.PostgresRepository.Get("email", email)
	if err != nil {
		return nil, err
	}
	if administrator.Password != password {
		return nil, fmt.Errorf("invalid credentials")
	}
	return administrator, nil
}
