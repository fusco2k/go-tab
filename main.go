package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

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
		tab, err := strconv.Atoi(ps.ByName("tab"))
		if err != nil {
			log.Printf("Error converting the string: %s", err)
		}
		session := &sessions.Session{
			// Tab: tabs.Tab{Number: int8(tab)},
		}
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		err = ss.CL.FindOne(ctx, bson.M{"tab": bson.M{"number": tab}}).Decode(&session)
		if err != nil && err.Error() == "mongo: no documents in result" {
			log.Printf("Creating a new session")
		} else {
			log.Printf("Error decoding the new session cursor: %s", err)
			cancel()
		}
		fmt.Println(session)
	}
}
