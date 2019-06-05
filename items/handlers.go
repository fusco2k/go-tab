package items

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fusco2k/go-tab/config"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Index handler to show all items
func Index(iEnv *config.Data) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//get all users from db
		items := AllData(iEnv.CL)
		//loop throht users slice and prints
		for _, item := range items {
			fmt.Fprintf(w, "%s, %s, %d, %t\n", item.ID, item.Name, item.Price, item.Visible)
		}
	}
}

//Get handler to get a especific item
func Get(iEnv *config.Data) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//parse the param id and ask for the correspondent item
		id, _ := primitive.ObjectIDFromHex(ps.ByName("id"))
		item := OneData(iEnv.CL, id)
		//prints the item
		fmt.Fprintf(w, "%s, %s, %d, %t\n", item.ID, item.Name, item.Price, item.Visible)
	}
}

//Create handler to create a new item
func Create(iEnv *config.Data) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		i := Item{}
		json.NewDecoder(r.Body).Decode(&i)
		obj := CreateData(iEnv.CL, i)
		//prints the id of the created item
		fmt.Fprintf(w, "create a item with the following id: %s", obj.Hex())
	}
}

//Edit handler to edit a existing item
func Edit(iEnv *config.Data) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		i := []Item{}
		//decode the json with the a slice with the item to be modified
		json.NewDecoder(r.Body).Decode(&i)
		obj := ModifyData(iEnv.CL, i)
		//prints the id of the modified item
		fmt.Fprintf(w, "modified the item with the following id: %s", obj.Hex())

	}
}

//Delete handler to delete a item
func Delete(iEnv *config.Data) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//parse the param id and ask for the correspondent item
		id, _ := primitive.ObjectIDFromHex(ps.ByName("id"))
		res := DeleteData(iEnv.CL, id)
		//prints the number of documents deleted
		fmt.Fprintf(w, "number of documents deleted: %v", res)
	}
}
