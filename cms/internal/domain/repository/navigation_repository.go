package repository

import (
	"prestaprest/internal/domain/entity"
)

type NavigationRepository interface {
	Create(navigation *entity.Navigation) error
	Update(navigation *entity.Navigation) error
	Get(navigation_ID uint) (*entity.Navigation, error)
	GetAll() ([]entity.Navigation, error)
	Delete(navigation_ID uint) error
}
