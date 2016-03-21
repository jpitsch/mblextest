package controllers

import (
	"github.com/astaxie/beego"
	"github.com/jpitsch/mblextest/models"
)

type AdminController struct {
	beego.Controller
}

/* This will persist to all other methods */
func (this *AdminController) Prepare() {
	this.Layout = "main-layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.tpl"

}

func (this *AdminController) Get() {
	this.TplName = "admin/admin-home.tpl"
	//flash := beego.ReadFromRequest(&this.Controller)
}

func (this *AdminController) GetTestAdminMenu() {
	this.TplName = "admin/test-admin.tpl"
}

/*
	This section is part of the admin of tests
*/
func (this *AdminController) GetCreateTest() {
	this.TplName = "admin/create-test.tpl"
	//done for now
}

func (this *AdminController) GetStartAddingQuestions() {
	this.TplName = "admin/create-question.tpl"

	test := Test{}
	test.Name = this.GetString("testName")
	test.Type = this.GetString("testType")

	models.SaveTest(test)

	this.SetSession("testname", test.Name)
}

func (this *AdminController) AddNextQuestion() {
	this.TplName = "admin/create-question.tpl"
	testname := this.GetSession("testname")

	answers := make([]Answer, 4)
	answers[0] = Answer{Text: this.GetString("answer1"), Position: 0}
	answers[1] = Answer{Text: this.GetString("answer2"), Position: 1}
	answers[2] = Answer{Text: this.GetString("answer3"), Position: 2}
	answers[3] = Answer{Text: this.GetString("answer4"), Position: 3}

	question := Question{Question: this.GetString("question"), Answers: answers}

	test := models.LoadTest(testname)

}

func (this *AdminController) GetEditTest() {
	this.TplName = "admin/test-admin.tpl"

	////TODO
}

func (this *AdminController) GetViewTests() {
	this.TplName = "admin/test-admin.tpl"

	////TODO
}

func (this *AdminController) GetNextQuestion() {
	this.TplName = "admin/test-admin.tpl"

	////TODO
}

func (this *AdminController) GetPreviousQuestion() {
	this.TplName = "admin/test-admin.tpl"

	////TODO
}

func (this *AdminController) FinalizeTestCreation() {
	this.TplName = "admin/test-admin.tpl"

	////TODO
}
