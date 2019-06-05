package main

import (
	"context"
	"html/template"
	"log"
	"net/http"

	"github.com/fusco2k/go-tab/items"

	"github.com/fusco2k/go-tab/menus"

	"github.com/fusco2k/go-tab/config"
	"github.com/fusco2k/go-tab/sessions"

	"github.com/julienschmidt/httprouter"
)

func main() {
	//dependency config
	cl := config.NewSession("mongodb://localhost:27017")
	defer cl.Disconnect(context.Background())
	mEnv := &config.Data{CL: cl.Database("go-tab").Collection("menu")}
	iEnv := &config.Data{CL: cl.Database("go-tab").Collection("item")}
	tEnv := &config.Data{CL: cl.Database("go-tab").Collection("tab")}
	sEnv := &config.Data{CL: cl.Database("go-tab").Collection("session")}
	//template manager
	tpl := config.TplManager()
	//new httprouter
	router := httprouter.New()
	//session routes
	router.GET("/api/rest/new", sessions.NewSS(tpl))
	router.POST("/api/rest/new", sessions.FindSS(sEnv))
	//tab routes
	router.GET("/api/rest/tabs/:tab", index(tEnv, tpl))
	//admin menu routes
	router.GET("/api/rest/admin/menus", menus.Index(mEnv))
	router.POST("/api/rest/admin/menus", menus.Create(mEnv))
	router.GET("/api/rest/admin/menu/:menu", menus.Get(mEnv))
	router.PUT("/api/rest/admin/menu/:menu", menus.Edit(mEnv))
	router.DELETE("/api/rest/admin/menu/:menu", menus.Delete(mEnv))
	//admin items routes
	router.GET("/api/rest/admin/items", items.Index(iEnv))
	router.POST("/api/rest/admin/items", items.Create(iEnv))
	router.GET("/api/rest/admin/item/:item", items.Get(iEnv))
	router.PUT("/api/rest/admin/item/:item", items.Edit(iEnv))
	router.DELETE("/api/rest/admin/item/:item", items.Delete(iEnv))

	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(sEnv *config.Data, tpl *template.Template) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		err := tpl.ExecuteTemplate(w, "menu.html", nil)
		if err != nil {
			log.Printf("Couldn't execute index.html: %s", err)
		}
	}
}
