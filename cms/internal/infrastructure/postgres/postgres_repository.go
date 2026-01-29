package postgres

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type PostgresRepository[T any] struct {
	db     *gorm.DB
	entity T
}

func NewPostgresRepository[T any](db *gorm.DB) PostgresRepository[T] {
	return PostgresRepository[T]{
		db: db,
	}
}

func (repo *PostgresRepository[T]) GetDB() *gorm.DB {
	return repo.db
}

func (repo *PostgresRepository[T]) Create(entity *T) error {
	c := context.Background()
	err := gorm.G[T](repo.db).Create(c, entity)
	return err
}

func (repo *PostgresRepository[T]) Update(entity *T) error {
	c := context.Background()
	err := repo.db.WithContext(c).Save(entity).Error
	return err
}

func (repo *PostgresRepository[T]) Get(param string, query any) (*T, error) {
	c := context.Background()
	fmt.Println(param, query)
	res, err := gorm.G[T](repo.db).Where(fmt.Sprintf("%s = ?", param), query).Take(c)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Element %s not found", param)
		}
		return nil, err
	}
	return &res, nil
}

func (repo *PostgresRepository[T]) GetAll(order string) ([]T, error) {
	c := context.Background()
	res, err := gorm.G[T](repo.db).Order(fmt.Sprintf("%s ASC", order)).Find(c)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (repo *PostgresRepository[T]) GetAllFiltered(order string, param string, query any) ([]T, error) {
	c := context.Background()
	res, err := gorm.G[T](repo.db).Where(fmt.Sprintf("%s = ?", param), query).Order(fmt.Sprintf("%s ASC", order)).Find(c)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (repo *PostgresRepository[T]) Delete(param string, ID uint64) error {
	c := context.Background()
	_, err := gorm.G[T](repo.db).Where(fmt.Sprintf("%s = ?", param), ID).Delete(c)
	return err
}

func (repo *PostgresRepository[T]) GetWithRelation(param string, query any, loader func(*T) error) (*T, error) {
	c := context.Background()

	res, err := gorm.G[T](repo.db).
		Where(fmt.Sprintf("%s = ?", param), query).
		Take(c)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Element %s not found", param)
		}
		return nil, err
	}

	if err := loader(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (repo *PostgresRepository[T]) GetAllWithRelation(param string, query any, order string, loader func(*T) error) ([]T, error) {
	c := context.Background()

	res, err := gorm.G[T](repo.db).
		Where(fmt.Sprintf("%s = ?", param), query).
		Order(fmt.Sprintf("%s ASC", order)).
		Find(c)

	if err != nil {
		return nil, err
	}

	for i := range res {
		if err := loader(&res[i]); err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (repo *PostgresRepository[T]) GetAllWithLoader(order string, loader func(*T) error) ([]T, error) {
	c := context.Background()

	res, err := gorm.G[T](repo.db).
		Order(fmt.Sprintf("%s ASC", order)).
		Find(c)

	if err != nil {
		return nil, err
	}

	for i := range res {
		if err := loader(&res[i]); err != nil {
			return nil, err
		}
	}

	return res, nil
}
