package postgres

import (
	"context"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(ctx context.Context, dsn string) (*gorm.DB, error) {
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return conn, nil
}
