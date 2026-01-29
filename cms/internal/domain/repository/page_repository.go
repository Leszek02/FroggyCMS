package repository

import (
	"prestaprest/internal/domain/entity"

	"github.com/labstack/echo/v4"
)

type PageRepository interface {
	Create(page *entity.Page) (*entity.Page, error)
	Update(page *entity.Page) (*entity.Page, error)
	Get(pageName string, ctx echo.Context) (*entity.Page, error)
	GetAll() ([]*entity.Page, error)
	Delete(page_ID int) error
}
