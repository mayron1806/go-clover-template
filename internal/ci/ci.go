package ci

import (
	"sync"

	"github.com/mayron1806/go-clover-core"
)

var container *Container
var containerOnce sync.Once

func InitContainer(c *clover.Container) {
	containerOnce.Do(func() {
		container = &Container{
			Container: c,
		}
	})
}
func GetContainer() *Container {
	if container == nil {
		panic("container not initialized")
	}
	return container
}
