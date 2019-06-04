package sessions

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/fusco2k/go-tab/config"
	"github.com/fusco2k/go-tab/tabs"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Session represents a session struct
type Session struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	SecretToken uuid.UUID          `json:"secrettoken,omitempty" bson:"secrettoken,omitempty"`
	Tab         tabs.Tab           `json:"tab,omitempty" bson:"tab,omitempty"`
	Time        primitive.DateTime `json:"time,omitempty" bson:"time,omitempty"`
}

//CheckSS check if the tab correspond to any existing client session
func CheckSS(w http.ResponseWriter, r *http.Request, ss *Session) {
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

//CreateSS configure and creates a session with given parameters
func CreateSS(w http.ResponseWriter, r *http.Request, tab int, ss *Session, sEnv *config.Session) {
	//configure session
	ss.SecretToken = uuid.New()
	ss.Tab = tabs.Tab{
		Number:     tab,
		TimeOpened: time.Now(),
		Table:      10,
	}
	ss.Time = 1
	//check if the cookie exists,and if so, clear it and set a new one
	if cookie, err := r.Cookie("SessionID"); cookie != nil {
		http.SetCookie(w, &http.Cookie{Name: "SessionID", MaxAge: -1})
		log.Printf("Clearing cookie: %s\n", err)
	} else if err != nil {
		log.Printf("Empty cookie: %s\n", err)
	}
	http.SetCookie(w, &http.Cookie{Name: "SessionID", Value: ss.SecretToken.String()})
	//write to the db
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := sEnv.CL.InsertOne(ctx, ss)
	if err != nil {
		log.Printf("Could not insert session on db: %s", err)
	}
	//print the ObjID inserted
	fmt.Println(result.InsertedID)
}
