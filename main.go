package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/fusco2k/go-tab/tabs"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/fusco2k/go-tab/config"
	"github.com/fusco2k/go-tab/sessions"

	"github.com/julienschmidt/httprouter"
)

func main() {
	//dependency config
	cl := config.NewSession("mongodb://localhost:27017")
	defer cl.Disconnect(context.Background())
	sEnv := &config.Session{CL: cl.Database("go-tab").Collection("session")}
	//template manager
	tpl := config.TplManager()
	//new httprouter
	router := httprouter.New()
	//routes
	router.GET("/api/rest/new", sessions.NewSS(tpl))
	router.POST("/api/rest/new", sessions.FindSS(sEnv))

	log.Fatal(http.ListenAndServe(":8080", router))
}
