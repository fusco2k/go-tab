package menus

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fusco2k/go-tab/config"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Index handler to show all menus
func Index(mEnv *config.Data) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//get all users from db
		items := AllData(mEnv.CL)
		//loop throht users slice and prints
		for _, menu := range items {
			fmt.Fprintf(w, "%s, %s, %T\n", menu.ID, menu.Category, menu.Items)
		}
	}
}

//Get handler to get a especific menu
func Get(mEnv *config.Data) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//parse the param id and ask for the correspondent menu
		id, _ := primitive.ObjectIDFromHex(ps.ByName("id"))
		menu := OneData(mEnv.CL, id)
		//prints the menu
		fmt.Fprintf(w, "%s, %s, %T\n", menu.ID, menu.Category, menu.Items)
	}
}

//Create handler to create a new menu
func Create(mEnv *config.Data) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		m := Menu{}
		json.NewDecoder(r.Body).Decode(&m)
		obj := CreateData(mEnv.CL, m)
		//prints the id of the created menu
		fmt.Fprintf(w, "create a menu with the following id: %s", obj.Hex())
	}
}

//Edit handler to edit a existing menu
func Edit(mEnv *config.Data) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		m := []Menu{}
		//decode the json with the a slice with the menu to be modified
		json.NewDecoder(r.Body).Decode(&m)
		obj := ModifyData(mEnv.CL, m)
		//prints the id of the modified menu
		fmt.Fprintf(w, "modified the menu with the following id: %s", obj.Hex())

	}
}

//Delete handler to delete a menu
func Delete(mEnv *config.Data) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//parse the param id and ask for the correspondent menu
		id, _ := primitive.ObjectIDFromHex(ps.ByName("id"))
		res := DeleteData(mEnv.CL, id)
		//prints the number of documents deleted
		fmt.Fprintf(w, "number of documents deleted: %v", res)
	}
}
