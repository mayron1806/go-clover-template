package main

import (
	"net/http"

	"github.com/mayron1806/go-clover-core"
)

func main() {
	c := clover.NewClover()

	c.ConfigureServer(&http.Server{}, false)

	c.Run()
}
