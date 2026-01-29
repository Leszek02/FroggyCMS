package service

import (
	"fmt"
	"prestaprest/internal/domain/entity"
	"prestaprest/internal/infrastructure/postgres"
	"strconv"

	"gorm.io/gorm"
)

type ClientService struct {
	PostgresRepository postgres.PostgresRepository[entity.Client]
}

func NewClientService(conn *gorm.DB) ClientService {
	return ClientService{
		PostgresRepository: postgres.NewPostgresRepository[entity.Client](conn),
	}
}

func (s *ClientService) GetAllClient() ([]entity.Client, error) {
	allClients, err := s.PostgresRepository.GetAll("Client_ID")
	if err != nil {
		return nil, err
	}

	return allClients, nil
}

func (s *ClientService) GetClient(client_ID int64) (*entity.Client, error) {
	client, err := s.PostgresRepository.Get("Client_ID", strconv.FormatUint(uint64(client_ID), 10))
	fmt.Println("GetClient: after repository")
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (s *ClientService) CreateClient(client *entity.Client) error {
	err := s.PostgresRepository.Create(client)
	if err != nil {
		return err
	}
	return nil
}

func (s *ClientService) UpdateClient(client *entity.Client) error {
	err := s.PostgresRepository.Update(client)
	if err != nil {
		return err
	}
	return nil
}

func (s *ClientService) DeleteClient(client_ID uint64) error {
	err := s.PostgresRepository.Delete("Client_ID", client_ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *ClientService) Authenticate(email, password string) (*entity.Client, error) {
	client, err := s.PostgresRepository.Get("email", email)
	if err != nil {
		return nil, err
	}
	if client.Password != password {
		return nil, fmt.Errorf("invalid credentials")
	}
	return client, nil
}
