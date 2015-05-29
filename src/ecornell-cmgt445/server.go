package main

import (
	"fmt"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

type Server struct {
	*martini.Martini
	martini.Router
}

func InitServer() *Server {
	r := martini.NewRouter()
	m := martini.New()
	m.Use(martini.Recovery())
	m.Use(martini.Static("assets", martini.StaticOptions{
		SkipLogging: true,
	}))
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)
	m.Use(render.Renderer(render.Options{
		Layout:  "master",
		Charset: "UTF-8",
	}))
	return &Server{m, r}
}

func main() {

	fmt.Println("----------------------------------------------")
	fmt.Println(" CMGT/445 - Wk4 - Implementation Plan Website")
	fmt.Println(" Elijah Cornell - 2015-06-01")
	fmt.Println("")
	fmt.Println(" Site URL:  http://localhost:3000")
	fmt.Println("----------------------------------------------")

	m := InitServer()

	m.Get("/", func(r render.Render) {

		vars := map[string]string{
			"title": "Home",
		}

		r.HTML(200, "home", vars)

	})

	m.Run()
}
