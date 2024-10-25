package main

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/mayron1806/go-clover-core"
	"github.com/mayron1806/go-clover-core/config"
	"github.com/mayron1806/go-clover-core/db"
	"github.com/mayron1806/go-clover-template/internal/ci"
	"github.com/mayron1806/go-clover-template/internal/router"
)

func main() {
	c := clover.NewClover()

	c.ConfigureServer(&http.Server{Addr: ":8080"}, false)

	dbOpts, err := config.NewEnvLoader[db.DatabaseOptions]().LoadEnv()
	if err != nil {
		panic(err)
	}
	database, err := c.ConfigureDatabase(*dbOpts)
	if err != nil {
		panic(err)
	}

	ci.InitContainer(clover.NewContainer(database))

	router.RegisterRoutes(c.Router())

	c.Run()
}
