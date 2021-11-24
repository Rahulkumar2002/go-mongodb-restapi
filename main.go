package main

import (
	"net/http" //We will use this to create a server using http.ListenAndServe method .

	"github.com/Rahulkumar2002/go-mongodb-restapi/controllers"
	"github.com/julienschmidt/httprouter" //We will use this to create simple routers.
	"gopkg.in/mgo.v2"                     //Used to connect with mongodb
)

func main() {
	r := httprouter.New()                             //New() function is creating a new instance of httprouter which is r .
	uc := controllers.NewUserController(getSession()) //we have created NewUserControllers in our
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}
