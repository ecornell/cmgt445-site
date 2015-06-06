package main

import (
	"fmt"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
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
	// m.Use(martini.Logger())
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
	fmt.Println(" CMGT/445 - Wk5 - Implementation Plan Website")
	fmt.Println(" Elijah Cornell - 2015-06-08")
	fmt.Println("")
	fmt.Println(" Site URL:  http://localhost:3000")
	fmt.Println("")
	fmt.Println(" Ctrl-C to Exit")
	fmt.Println("----------------------------------------------")

	m := InitServer()

	m.Get("/", func(r render.Render) {
		vars := map[string]string{
			"title": "Home",
		}
		r.HTML(200, "home", vars)
	})

	m.Get("/plan", func(r render.Render) {
		vars := map[string]string{
			"title": "Implementation Plan",
		}
		r.HTML(200, "plan", vars)
	})

	m.Get("/glossary", func(r render.Render) {
		vars := map[string]string{
			"title": "Glossary",
		}
		r.HTML(200, "glossary", vars)
	})

	m.Get("/references", func(r render.Render) {
		vars := map[string]string{
			"title": "References",
		}
		r.HTML(200, "references", vars)
	})

	m.Get("/about", func(r render.Render) {
		vars := map[string]string{
			"title": "About",
		}
		r.HTML(200, "about", vars)
	})

	m.Get("/bug", func(r render.Render) {
		vars := map[string]string{
			"title": "Submit Bug",
		}
		r.HTML(200, "bug", vars)
	})

	type BugForm struct {
		Description string `form:"desc"`
		Severity    string `form:"severity"`
		Steps       string `form:"steps"`
	}

	m.Post("/bug", binding.Bind(BugForm{}), func(bug BugForm, r render.Render) {

		vars := map[string]interface{}{
			"title": "Submit Bug",
		}

		errors := []string{}
		if len(bug.Description) == 0 {
			errors = append(errors, "Short Desciption Is Required")
		}

		if len(errors) == 0 {
			vars["note"] = "Bug Report Submitted Successfully"

			fmt.Println("----------------------------------------------")
			fmt.Println(" Bug Report - Success")
			fmt.Println("   Description: " + bug.Description)
			fmt.Println("   Severity: " + bug.Severity)
			fmt.Println("   Steps: " + bug.Steps)

		} else {
			vars["data"] = bug
			vars["errors"] = errors

			fmt.Println("----------------------------------------------")
			fmt.Println(" Bug Report - Error")
		}

		r.HTML(200, "bug", vars)
	})

	m.Get("/rfc", func(r render.Render) {
		vars := map[string]string{
			"title": "Request for Change",
		}
		r.HTML(200, "rfc", vars)
	})

	type RfcForm struct {
		Requestor string `form:"requestor"`

		CategoryCA string `form:"cat-ca"`
		CategoryPA string `form:"cat-pa"`
		CategoryDR string `form:"cat-dr"`
		CategoryU  string `form:"cat-u"`
		CategoryO  string `form:"cat-o"`

		AffectSchedule string `form:"affect-schedule"`
		AffectCost     string `form:"affect-cost"`
		AffectScope    string `form:"affect-scope"`
		AffectReq      string `form:"affect-req"`
		AffectTest     string `form:"affect-test"`
		AffectResource string `form:"affect-resource"`

		Details      string `form:"detail"`
		Reason       string `form:"reason"`
		Alternatives string `form:"alternatives"`
		Risks        string `form:"risks"`
		Costs        string `form:"costs"`
		Quality      string `form:"quality"`
	}

	m.Post("/rfc", binding.Bind(RfcForm{}), func(rfc RfcForm, r render.Render) {

		vars := map[string]interface{}{
			"title": "Request for Change",
			"data":  rfc,
		}

		errors := []string{}
		if len(rfc.Requestor) == 0 {
			errors = append(errors, "Requestor Is Required")
		}
		if len(rfc.Details) == 0 {
			errors = append(errors, "Description of Change Is Required")
		}
		if len(rfc.Reason) == 0 {
			errors = append(errors, "Reason for Change Is Required")
		}

		if len(errors) == 0 {
			vars["note"] = "RFC Submitted Successfully"

			fmt.Println("----------------------------------------------")
			fmt.Println(" RFC - Success")
			fmt.Println("   Requestor: " + rfc.Requestor)

		} else {
			vars["data"] = rfc
			vars["errors"] = errors
			fmt.Println("----------------------------------------------")
			fmt.Println(" RFC - Error")
		}

		r.HTML(200, "rfc", vars)
	})

	m.Run()
}
