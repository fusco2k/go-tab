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
	router.GET("/api/rest/new", newSS(tpl))
	router.POST("/api/rest/new", findSS(sEnv))

	log.Fatal(http.ListenAndServe(":8080", router))
}

func newSS(tpl *template.Template) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		err := tpl.ExecuteTemplate(w, "open.html", nil)
		if err != nil {
			log.Printf("Couldn't execute open.gohtml: %s", err)
		}
	}
}

func findSS(sEnv *config.Session) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//parse the tab
		tab, err := strconv.Atoi(r.FormValue("tab"))
		if err != nil {
			log.Printf("Error converting the string: %s\n", err)
		}
		ss := &sessions.Session{}
		//check if the tab has an opened session, case not, create a new one
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		err = sEnv.CL.FindOne(ctx, bson.M{"tab.number": tab}).Decode(&ss)
		if err != nil && err.Error() == "mongo: no documents in result" {
			log.Printf("Error decoding the new session cursor: %s\n", err)
			log.Printf("Creating a new session")
			createSS(w, r, tab, ss, sEnv)
		} else {
			log.Println("Checking the session")
			checkSS(w, r, ss)
		}
	}
}

func checkSS(w http.ResponseWriter, r *http.Request, ss *sessions.Session) {
	cookie, err := r.Cookie("SessionID")
	if err != nil {
		log.Printf("Error getting cookie: %s\n", err)
	}
	switch cookie.Value == ss.SecretToken.String() {
	case false:
		fmt.Println("here")
		//todo deal if the tab reports an opened session and the client already has a valid and opened session. Check why it has an opened session, that should not happen
	case true:
		r.Method = "GET"
		http.Redirect(w, r, "/api/rest/"+strconv.Itoa(ss.Tab.Number), http.StatusFound)
	}
}

func createSS(w http.ResponseWriter, r *http.Request, tab int, session *sessions.Session, ss *config.Session) {
	//configure session
	session.SecretToken = uuid.New()
	session.Tab = tabs.Tab{
		Number:     tab,
		TimeOpened: time.Now(),
		Table:      10,
	}
	session.Time = 1
	//check if the cookie exists,and if so, clear it and set a new one
	if cookie, err := r.Cookie("SessionID"); cookie != nil {
		http.SetCookie(w, &http.Cookie{Name: "SessionID", MaxAge: -1})
		log.Printf("Clearing cookie: %s\n", err)
	} else if err != nil {
		log.Printf("Empty cookie: %s\n", err)
	}
	http.SetCookie(w, &http.Cookie{Name: "SessionID", Value: session.SecretToken.String()})
	//write to the db
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := ss.CL.InsertOne(ctx, session)
	if err != nil {
		log.Printf("Could not insert session on db: %s", err)
	}
	//print the ObjID inserted
	fmt.Println(result.InsertedID)
}
