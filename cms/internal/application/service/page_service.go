package service

import (
	"fmt"
	"prestaprest/internal/domain/entity"
	"prestaprest/internal/infrastructure/postgres"

	"gorm.io/gorm"
)

type PageService struct {
	PostgresRepository postgres.PostgresRepository[entity.Page]
}

func NewPageService(conn *gorm.DB) PageService {
	return PageService{
		PostgresRepository: postgres.NewPostgresRepository[entity.Page](conn),
	}
}

func (s *PageService) GetAllPages() ([]entity.Page, error) {
	allPages, err := s.PostgresRepository.GetAll("type")
	if err != nil {
		return nil, err
	}
	return allPages, nil
}

func (s *PageService) GetPage(pageName string) (*entity.Page, error) {
	page, err := s.PostgresRepository.Get("Name", pageName)
	fmt.Println("Am here (service)?")
	if err != nil {
		return nil, err
	}
	return page, nil
}

func (s *PageService) GetPageByID(page_ID uint) (*entity.Page, error) {
	page, err := s.PostgresRepository.Get("Page_ID", uint64(page_ID))
	if err != nil {
		return nil, err
	}
	return page, nil
}

func (s *PageService) CreatePage(page *entity.Page) error {
	err := s.PostgresRepository.Create(page)
	if err != nil {
		return err
	}
	return nil
}

func (s PageService) UpdatePage(page *entity.Page) error {
	err := s.PostgresRepository.Update(page)
	if err != nil {
		return err
	}
	return nil
}

func (s *PageService) DeletePage(page_ID uint) error {
	err := s.PostgresRepository.Delete("Page_ID", uint64(page_ID))
	if err != nil {
		return err
	}
	return nil
}
