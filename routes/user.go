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
}

func NewUserRouter(s *mgo.Session) *UserRouter {
	return &UserRouter{s}
}

func (ur UserRouter) ReadAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var all []models.User

	err := ur.session.DB(database).C(collection).Find(nil).All(&all)

	if err != nil {
		respond(w, http.StatusInternalServerError, nil)
	}

	uj, _ := json.Marshal(all)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", uj)
}

func (ur UserRouter) Read(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	//Bad Request
	defer func() {
		recover()
	}()

	if !bson.IsObjectIdHex(id) {
		respond(w, http.StatusBadRequest, nil)
	}

	oid := bson.ObjectIdHex(id)

	u := models.User{}

	if err := ur.session.DB(database).C(collection).FindId(oid).One(&u); err != nil {
		respond(w, http.StatusNotFound, nil)
	}

	uj, _ := json.Marshal(u)

	respond(w, http.StatusOK, uj)
}

func (ur UserRouter) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)

	uj, _ := json.Marshal(u)

	respond(w, http.StatusOK, uj)
}

// CreateUser creates a new user resource
func (ur UserRouter) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)
	u.Id = bson.NewObjectId()
	ur.session.DB(database).C(collection).Insert(u)
	uj, _ := json.Marshal(u)

	respond(w, http.StatusCreated, uj)
}

// RemoveUser removes an existing user resource
func (ur UserRouter) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		respond(w, http.StatusBadRequest, nil)
	}

	oid := bson.ObjectIdHex(id)

	if err := ur.session.DB(database).C(collection).RemoveId(oid); err != nil {
		respond(w, http.StatusNotFound, nil)
	}

	respond(w, http.StatusOK, nil)
}

func respond(w http.ResponseWriter, code int, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprintf(w, "%s", response)
}
