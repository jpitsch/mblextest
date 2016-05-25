package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var (
	mgoSession   *mgo.Session
	databaseName = "pitsch_test"
	databaseURI  = "mtest:password@ds015398.mongolab.com:15398/mtest_dev"
)

type User struct {
	FirstName string    `bson:"firstname"`
	LastName  string    `bson:"lastname"`
	UserName  string    `bson:"username"`
	Password  string    `bson:"password"`
	Created   time.Time `bson:"created"`
	Updated   time.Time `bson:"updated"`
	Email     string    `bson:"email"`
	UserProps
}

type UserProps struct {
	Id int
}

type Test struct {
	Id        int
	Name      string
	TestType  string 	 	
	Publish   bool
	Question  []Question
}

type Question struct {
	Id        	   int	
	Selected       string 			//`bson:"selected"`
	CorrectAnswer  string 			//`bson:"correctAnswer"`
	Text       	   string 			//`bson:"text"`
	Number		   int 	  			//`bson:"number"`
	Answer 		   []Answer
}

type Answer struct {
	Id		 int
	Position int    	//`bson:"position"`
	Text     string 	//`bson:"text"`
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

func SaveTest(t Test) {
	session := getSession()
	defer session.Close()

	c := session.DB("mtest_dev").C("testcollection")
	err := c.Insert(t)

	if err != nil {
		fmt.Println("Error while trying to save Test to Mongo: " + t.Name)
		fmt.Println(err)
	}
}

func SaveTestQuestion(t Test, q Question) {
	session := getSession()
	defer session.Close()

	c := session.DB("mtest_dev").C("testcollection")

	c.update({"name": t.Name}, {"$addToSet": {Question: q}})

	if err != nil {
		fmt.Println("Error while trying to add a new question to: " + t.Name)
		fmt.Println(err)
	}	
}

func LoadTest(name string) Test {
	session := getSession()
	defer session.Close()

	tests := session.DB("mtest_dev").C("testcollection")
	testresult := Test{}
	err := tests.Find(bson.M{"name": name}).One(&testresult)

	if err != nil {
		fmt.Println("No test by that name " + name)
	}

	return testresult
}

func NewTest() Test {
	return Test{}
}
