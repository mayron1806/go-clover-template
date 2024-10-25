package ci

import (
	"github.com/mayron1806/go-clover-core"
	"github.com/mayron1806/go-clover-core/logger"
	"github.com/mayron1806/go-clover-template/internal/repository"
)

type Container struct {
	*clover.Container

	userRepository *repository.UserRepository
}

func (c *Container) Init() {
	c.InitOnce.Do(func() {
		c.userRepository = repository.NewUserRepository(
			c.Database(),
			logger.NewLogger(logger.LoggerOptions{Prefix: "[USER REPOSITORY]"}),
		)
	})
}
func (c *Container) UserRepository() *repository.UserRepository {
	c.Init()
	return c.userRepository
}
