package repository

import (
	"github.com/mayron1806/go-clover-core"
	"github.com/mayron1806/go-clover-core/db"
	"github.com/mayron1806/go-clover-core/logger"
)

type UserRepository struct {
	*clover.Repository
}

func NewUserRepository(db *db.Database, logger *logger.Logger) *UserRepository {
	return &UserRepository{
		Repository: clover.NewRepository(db, logger),
	}
}
