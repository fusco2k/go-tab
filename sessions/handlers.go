package sessions

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/fusco2k/go-tab/config"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
)

//NewSS opens a new session entering a tab number
func NewSS(tpl *template.Template) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		err := tpl.ExecuteTemplate(w, "open.html", nil)
		if err != nil {
			log.Printf("Couldn't execute open.gohtml: %s", err)
		}
	}
}

//FindSS looks for the tab on any active session
func FindSS(sEnv *config.Data) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//parse the tab
		tab, err := strconv.Atoi(r.FormValue("tab"))
		if err != nil {
			log.Printf("Error converting the string: %s\n", err)
		}
		ss := &Session{}
		//check if the tab has an opened session, case not, create a new one
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		err = sEnv.CL.FindOne(ctx, bson.M{"tab.number": tab}).Decode(&ss)
		if err != nil && err.Error() == "mongo: no documents in result" {
			log.Printf("Error decoding the new session cursor: %s\n", err)
			log.Printf("Creating a new session")
			CreateSS(w, r, tab, ss, sEnv)
		} else {
			log.Println("Checking the session")
			CheckSS(w, r, ss)
		}
	}
}
