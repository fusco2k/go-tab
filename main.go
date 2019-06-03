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
	ss := &config.Session{CL: cl.Database("go-tab").Collection("session")}
	//template manager
	tpl := config.TplManager()
	//new httprouter
	router := httprouter.New()
	//routes
	router.GET("/api/rest/open", openSS(tpl))
	router.POST("/api/rest/new", newSS(ss))

	log.Fatal(http.ListenAndServe(":8080", router))
}

func openSS(tpl *template.Template) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		err := tpl.ExecuteTemplate(w, "open.gohtml", nil)
		if err != nil {
			log.Printf("Couldn't execute open.gohtml: %s", err)
		}
	}
}

func newSS(ss *config.Session) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//parse the tab
		tab, err := strconv.Atoi(ps.ByName("tab"))
		if err != nil {
			log.Printf("Error converting the string: %s\n", err)
		}
		session := &sessions.Session{}
		//check if the tab has an opened session, case not, create a new one
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		err = ss.CL.FindOne(ctx, bson.M{"tab.number": 15}).Decode(&session)
		if err != nil && err.Error() == "mongo: no documents in result" {
			log.Printf("Creating a new session")
			fmt.Println(tab)
			createSession(w, r, tab, session, ss)
		}
		if err != nil {
			log.Printf("Error decoding the new session cursor: %s\n", err)
			cancel()
		}
		//print the session for debug
		fmt.Println(session)
	}
}

func createSession(w http.ResponseWriter, r *http.Request, tab int, session *sessions.Session, ss *config.Session) {
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
