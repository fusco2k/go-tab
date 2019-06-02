package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/fusco2k/go-tab/tabs"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/fusco2k/go-tab/config"
	"github.com/fusco2k/go-tab/sessions"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	//dependency config
	cl := config.NewSession("mongodb://localhost:27017")
	defer cl.Disconnect(context.Background())
	ss := &config.Session{CL: cl.Database("go-tab").Collection("session")}

	router.GET("/api/rest/:tab/new", newSession(ss))

	log.Fatal(http.ListenAndServe(":8080", router))
}

func newSession(ss *config.Session) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//parse the tab
		tab, err := strconv.Atoi(ps.ByName("tab"))
		if err != nil {
			log.Printf("Error converting the string: %s/n", err)
		}
		session := &sessions.Session{}
		//check if the tab has an opened session, case not, create a new one
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		err = ss.CL.FindOne(ctx, bson.M{"tab": bson.M{"number": tab}}).Decode(&session)
		if err != nil && err.Error() == "mongo: no documents in result" {
			log.Printf("Creating a new session")
			createSession(w, r, tab, session, ss)
		} else {
			log.Printf("Error decoding the new session cursor: %s/n", err)
			cancel()
		}
		//print the session for debug
		fmt.Println(session)
	}
}
func createSession(w http.ResponseWriter, r *http.Request, tab int, session *sessions.Session, ss *config.Session) {
	//configure session
	session.ID = primitive.NewObjectID()
	session.SecretToken = uuid.New()
	session.Tab = tabs.Tab{
		ID:         primitive.NewObjectID(),
		Number:     12,
		TimeOpened: time.Now(),
		Table:      15,
	}
	session.Time = 1
	//check if the cookie exists,and if so, clear it and set a new one
	if cookie, err := r.Cookie("SessionID"); cookie != nil {
		http.SetCookie(w, &http.Cookie{Name: "SessionID", MaxAge: -1})
	} else if err != nil {
		log.Printf("Could not set cookie: %s/n", err)
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
