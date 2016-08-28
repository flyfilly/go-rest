package main

import (
	"github.com/flyfilly/routes"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"
)

func main() {
	r := httprouter.New()

	userRoutes := routes.NewUserRouter(getSession())
	r.GET("/user", userRoutes.ReadAll)
	r.GET("/user/:id", userRoutes.ReadOne)
	r.POST("/user", userRoutes.Create)
	r.PUT("/user/:id", userRoutes.Update)
	r.DELETE("/user/:id", userRoutes.Delete)

	//For https
	// http.ListenAndServeTLS("addr", "certFilePath", "keyFilePath", r)
	
	//For development
	http.ListenAndServe(":8080", r)
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return session
}
