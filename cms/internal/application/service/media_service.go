package service

import (
	"fmt"
	"prestaprest/internal/domain/entity"
	"prestaprest/internal/infrastructure/postgres"

	"gorm.io/gorm"
)

type MediaService struct {
	PostgresRepository      postgres.PostgresRepository[entity.Media]
	ProductsMediaRepository postgres.PostgresRepository[entity.Products_Medias]
}

func NewMediaService(conn *gorm.DB) MediaService {
	return MediaService{
		PostgresRepository:      postgres.NewPostgresRepository[entity.Media](conn),
		ProductsMediaRepository: postgres.NewPostgresRepository[entity.Products_Medias](conn),
	}
}

func (m *MediaService) GetAllMedia() ([]entity.Media, error) {
	allMedia, err := m.PostgresRepository.GetAll("Media_ID")
	if err != nil {
		return nil, err
	}

	return allMedia, nil
}

func (m *MediaService) GetMedia(media_ID uint) (*entity.Media, error) {
	media, err := m.PostgresRepository.Get("navigation_id", media_ID)

	fmt.Println("GetNavigation: after repository")
	if err != nil {
		return nil, err
	}
	return media, nil
}

func (m *MediaService) CreateMedia(media *entity.Media) error {
	err := m.PostgresRepository.Create(media)
	if err != nil {
		return err
	}
	return nil
}

func (m *MediaService) UpdateProductMedia(productsMedia *entity.Products_Medias) error {
	err := m.ProductsMediaRepository.Create(productsMedia)
	if err != nil {
		return err
	}

	return nil
}

func (m *MediaService) DeleteMedia(media_ID uint) error {
	err := m.PostgresRepository.Delete("Media_ID", uint64(media_ID))
	if err != nil {
		return err
	}
	return nil
}
