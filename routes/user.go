package routes

import (
	"encoding/json"
	"fmt"
	"github.com/flyfilly/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

const (
	database   string = "go-rest"
	collection string = "users"
)

type UserRouter struct {
	session *mgo.Session
	collection *mgo.Collection
}

func NewUserRouter(s *mgo.Session) *UserRouter {
	return &UserRouter{s, s.DB(database).C(collection)}
}

func (this UserRouter) ReadAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var all []models.User

	err := this.collection.Find(nil).All(&all)

	if err != nil {
		respond(w, http.StatusInternalServerError, nil)
		return
	}

	uj, _ := json.Marshal(all)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", uj)
}

func (this UserRouter) ReadOne(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	//Bad Request
	defer func() {
		recover()
	}()

	if !bson.IsObjectIdHex(id) {
		respond(w, http.StatusBadRequest, nil)
		return
	}

	oid := bson.ObjectIdHex(id)

	u := models.User{}

	if err := this.collection.FindId(oid).One(&u); err != nil {
		respond(w, http.StatusNoContent, nil)
		return
	}

	uj, _ := json.Marshal(u)

	respond(w, http.StatusOK, uj)
}

//TODO Update for PUT
func (this UserRouter) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)

	uj, _ := json.Marshal(u)

	respond(w, http.StatusOK, uj)
}

func (this UserRouter) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)
	u.Id = bson.NewObjectId()
	this.collection.Insert(u)
	uj, _ := json.Marshal(u)

	respond(w, http.StatusCreated, uj)
}

func (this UserRouter) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	//Bad Request
	defer func() {
		recover()
	}()

	if !bson.IsObjectIdHex(id) {
		respond(w, http.StatusBadRequest, nil)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := this.collection.RemoveId(oid); err != nil {
		respond(w, http.StatusNoContent, nil)
		return
	}

	respond(w, http.StatusOK, nil)
}

func respond(w http.ResponseWriter, code int, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprintf(w, "%s", response)
	return
}