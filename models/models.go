package models

import (
	"fmt"
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	mgoSession   *mgo.Session
	databaseName = "pitsch_test"
	databaseURI = "mtest:password@ds015398.mongolab.com:15398/mtest_dev"
)

type User struct {
	Id 			int 		`bson:"_id"`
	FirstName 	string		`bson:"firstname"` 
	LastName 	string		`bson:"lastname"`
	UserName	string 		`bson:"username"`
	Password	string 		`bson:"password"`
	Created		time.Time   `bson:"created"`
	Updated		time.Time   `bson:"updated"`
	Email		string		`bson:"email"`
	UserProps	*UserProps  
}

type UserProps struct {
	Id		int
}

type Test struct {
	Id			int 		`bson:"_id"`
	Name 		string		`bson:"name"`
	Questions	[]Question
}

type Question struct {
	Id			int 		`bson:"_id"`
	Selected	int 		`bson:"selected"`
	Question	string 		`bson:"question"`
	Answers		[]Answer

	Test 		*Test
}

type Answer struct {
	Id 			int 		`bson:"_id"`
	Position	int 		`bson:"position"`
	Text		string 		`bson:"text"`

	Questions 	*Question 
}


//Only get the copy of the session for Mongo database
func getSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(databaseURI)
		if err != nil {
			fmt.Println("Error connecting to Mongo: ", err)
		}
	}
	return mgoSession.Copy()
}

//Load user information for Login
func LoadUser(user string) User {
	session := getSession()
	defer session.Close()

	users := session.DB("mtest_dev").C("usercollection")

	userresult := User{}
	err := users.Find(bson.M{"username": user}).One(&userresult)
	if err != nil {
		fmt.Println("No such user: ")
	}
	fmt.Println("User found.")
	return userresult
}