package service

import (
	"fmt"
	"prestaprest/internal/domain/entity"
	"prestaprest/internal/infrastructure/postgres"

	"gorm.io/gorm"
)

type NavigationService struct {
	PostgresRepository postgres.PostgresRepository[entity.Navigation]
}

func NewNavigationService(conn *gorm.DB) NavigationService {
	return NavigationService{
		PostgresRepository: postgres.NewPostgresRepository[entity.Navigation](conn),
	}
}

func (s *NavigationService) GetAllNavigation() ([]*entity.Navigation, error) {
	allNavs, err := s.PostgresRepository.GetAll("position")
	if err != nil {
		return nil, err
	}

	navMap := make(map[uint]*entity.Navigation)
	var rootNavs []*entity.Navigation

	for i := range allNavs {
		// if allNavs[i].Status == false {
		// 	continue
		// }
		nav := &allNavs[i]
		navMap[nav.Navigation_ID] = nav
		nav.Children = []*entity.Navigation{}
	}

	for i := range allNavs {
		// if allNavs[i].Status == false {
		// 	continue
		// }
		nav := &allNavs[i]
		if nav.Navigation_Navigation_ID == nil {
			rootNavs = append(rootNavs, nav)
		} else {
			if parent, exists := navMap[*nav.Navigation_Navigation_ID]; exists {
				parent.Children = append(parent.Children, nav)
			}
		}
	}

	return rootNavs, nil
}

func (s *NavigationService) GetNavigation(navigation_ID uint) (*entity.Navigation, error) {
	navigation, err := s.PostgresRepository.Get("navigation_id", navigation_ID)

	fmt.Println("GetNavigation: after repository")
	if err != nil {
		return nil, err
	}
	return navigation, nil
}

func (s *NavigationService) CreateNavigation(navigation *entity.Navigation) error {
	err := s.PostgresRepository.Create(navigation)
	if err != nil {
		return err
	}
	return nil
}

func (s *NavigationService) UpdateNavigation(navigation *entity.Navigation) error {
	err := s.PostgresRepository.Update(navigation)
	if err != nil {
		return err
	}
	return nil
}

func (s *NavigationService) DeleteNavigation(navigation_ID uint) error {
	err := s.PostgresRepository.Delete("Navigation_ID", uint64(navigation_ID))
	if err != nil {
		return err
	}
	return nil
}
