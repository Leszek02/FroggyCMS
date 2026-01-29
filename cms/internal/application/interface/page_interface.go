package interfaces

import "prestaprest/internal/domain/entity"

type PageInterface interface {
	CreatePage()
	UpdatePage()
	DeletePage()
	GetPage() (entity.Page, error)
	GetAllPages()
}
