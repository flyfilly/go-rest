package main

import (
	"github.com/flyfilly/routes"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"net/http"
)

func main() {
	r := httprouter.New()

	userRoutes := routes.NewUserRouter(getSession())
	r.GET("/user", userRoutes.ReadAll)
	r.GET("/user/:id", userRoutes.Read)
	r.POST("/user", userRoutes.Create)
	r.PUT("/user/:id", userRoutes.Update)
	r.DELETE("/user/:id", userRoutes.Delete)

	http.ListenAndServe(":8080", r)
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return session
}
