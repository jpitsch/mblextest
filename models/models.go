package main

import {
	"time"
	"github.com/astaxie/beego/orm"
}

type User struct {
	Id 			int
	FirstName 	string
	LastName 	string
	UserName	string
	Password	string
	Created		time.Time
	Updated		time.Time
	UserProps	*UserProps
}

type UserProps struct {
	Id		int
}

type Test struct {
	Id		int
}

type Question struct {
	Id		int
}


// This will register the models to the ORM
func init() {
	orm.RegisterModel(new(User), new(Test))
}