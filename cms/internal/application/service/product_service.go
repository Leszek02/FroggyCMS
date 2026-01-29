package service

import (
	"prestaprest/internal/domain/entity"
	"prestaprest/internal/infrastructure/postgres"
	"strings"

	"gorm.io/gorm"
)

type ProductService struct {
	PostgresRepository       postgres.PostgresRepository[entity.Product]
	MediaRepository          postgres.PostgresRepository[entity.Media]
	ProductsMediasRepository postgres.PostgresRepository[entity.Products_Medias]
}

func NewProductService(conn *gorm.DB) ProductService {
	return ProductService{
		PostgresRepository:       postgres.NewPostgresRepository[entity.Product](conn),
		MediaRepository:          postgres.NewPostgresRepository[entity.Media](conn),
		ProductsMediasRepository: postgres.NewPostgresRepository[entity.Products_Medias](conn),
	}
}

func (s *ProductService) GetAllProduct() ([]entity.Product, error) {
	loadMedia := func(product *entity.Product) error {
		var mediaItems []entity.Media

		err := s.MediaRepository.GetDB().
			Table("medias").
			Select("medias.*").
			Joins("INNER JOIN products_medias ON products_medias.medias_media_id = medias.media_id").
			Where("products_medias.products_product_id = ?", product.ProductID).
			Scan(&mediaItems).Error

		if err != nil {
			return nil
		}

		for _, media := range mediaItems {
			if strings.Contains(media.Path, "/products/main") || strings.Contains(media.Path, "products/main") {
				product.ProductMainImage = media.Path
			} else if strings.Contains(media.Path, "/products/icons") || strings.Contains(media.Path, "products/icons") {
				product.ProductIconImage = media.Path
			}
		}

		return nil
	}
	allProducts, err := s.PostgresRepository.GetAllWithLoader("product_id", loadMedia)
	if err != nil {
		return nil, err
	}

	return allProducts, nil
}

func (s *ProductService) GetProduct(productId uint64) (*entity.Product, error) {
	loadMedia := func(product *entity.Product) error {
		var mediaItems []entity.Media
		err := s.MediaRepository.GetDB().
			Table("medias").
			Select("medias.*").
			Joins("INNER JOIN products_medias ON products_medias.medias_media_id = medias.media_id").
			Where("products_medias.products_product_id = ?", productId).
			Scan(&mediaItems).Error

		if err != nil {
			return err
		}

		for _, media := range mediaItems {
			if strings.Contains(media.Path, "/products/main") || strings.Contains(media.Path, "products/main") {
				product.ProductMainImage = media.Path
			} else if strings.Contains(media.Path, "/products/icons") || strings.Contains(media.Path, "products/icons") {
				product.ProductIconImage = media.Path
			}
		}

		return nil
	}

	product, err := s.PostgresRepository.GetWithRelation("product_id", productId, loadMedia)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) CreateProduct(product *entity.Product) error {
	err := s.PostgresRepository.Create(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) UpdateProduct(product *entity.Product) error {
	err := s.PostgresRepository.Update(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) DeleteProduct(product_ID uint64) error {
	err := s.PostgresRepository.Delete("Product_ID", uint64(product_ID))
	if err != nil {
		return err
	}
	return nil
}
