package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	models "sitepointgoapp/models"
	"strconv"
)

type ManageController struct {
	beego.Controller
}

func (manage *ManageController) Home() {
	manage.Layout = "basic-layout.tpl"
	manage.LayoutSections = make(map[string]string)
	manage.LayoutSections["Header"] = "header.tpl"
	manage.LayoutSections["Footer"] = "footer.tpl"
	manage.TplNames = "manage/home.tpl"
}

/**
 * Delete the user, by their id, if that user's available.
 */
func (manage *ManageController) Delete() {
	manage.Layout = "basic-layout.tpl"
	manage.LayoutSections = make(map[string]string)
	manage.LayoutSections["Header"] = "header.tpl"
	manage.LayoutSections["Footer"] = "footer.tpl"
	manage.TplNames = "manage/home.tpl"

	// convert the string value to an int
	articleId, _ := strconv.Atoi(manage.Ctx.Input.Param(":id"))

	o := orm.NewOrm()
	o.Using("default")

	article := models.Article{}

	// Check if the article exists first
	if exist := o.QueryTable(article.TableName()).Filter("Id", articleId).Exist(); exist {
		if num, err := o.Delete(&models.Article{Id: articleId}); err == nil {
			beego.Info("Record Deleted. ", num)
		} else {
			beego.Error("Record couldn't be deleted. Reason: ", err)
		}
	} else {
		beego.Info("Record Doesn't exist.")
	}
}

func (manage *ManageController) Update() {
	o := orm.NewOrm()
	o.Using("default")
	flash := beego.NewFlash()

	// convert the string value to an int
	if articleId, err := strconv.Atoi(manage.Ctx.Input.Param(":id")); err == nil {
		article := models.Article{Id: articleId}
		// attempt to load the record from the database
		if o.Read(&article) == nil {
			// set the Client and Url properties to arbitrary values
			article.Client = "Sitepoint"
			article.Url = "http://www.google.com"
			if num, err := o.Update(&article); err == nil {
				flash.Notice("Record Was Updated.")
				flash.Store(&manage.Controller)
				beego.Info("Record Was Updated. ", num)
			}
		} else {
			flash.Notice("Record Was NOT Updated.")
			flash.Store(&manage.Controller)
			beego.Error("Couldn't find article matching id: ", articleId)
		}
	} else {
		flash.Notice("Record Was NOT Updated.")
		flash.Store(&manage.Controller)
		beego.Error("Couldn't convert id from a string to a number. ", err)
	}

	// redirect afterwards
	manage.Redirect("/manage/view", 302)
}

func (manage *ManageController) View() {
	manage.Layout = "basic-layout.tpl"
	manage.LayoutSections = make(map[string]string)
	manage.LayoutSections["Header"] = "header.tpl"
	manage.LayoutSections["Footer"] = "footer.tpl"
	manage.TplNames = "manage/view.tpl"

	flash := beego.ReadFromRequest(&manage.Controller)

	if ok := flash.Data["error"]; ok != "" {
		// Display error messages
		manage.Data["errors"] = ok
	}

	if ok := flash.Data["notice"]; ok != "" {
		// Display error messages
		manage.Data["notices"] = ok
	}

	o := orm.NewOrm()
	o.Using("default")

	var articles []*models.Article
	num, err := o.QueryTable("articles").All(&articles)

	if err != orm.ErrNoRows && num > 0 {
		manage.Data["records"] = articles
	}
}

func (manage *ManageController) Add() {
	manage.Data["Form"] = &models.Article{}
	manage.Layout = "basic-layout.tpl"
	manage.LayoutSections = make(map[string]string)
	manage.LayoutSections["Header"] = "header.tpl"
	manage.LayoutSections["Footer"] = "footer.tpl"
	manage.TplNames = "manage/add.tpl"

	flash := beego.ReadFromRequest(&manage.Controller)

	if ok := flash.Data["error"]; ok != "" {
		// Display error messages
		manage.Data["flash"] = ok
	}

	o := orm.NewOrm()
	o.Using("default")

	article := models.Article{}

	if err := manage.ParseForm(&article); err != nil {
		beego.Error("Couldn't parse the form. Reason: ", err)
	} else {
		manage.Data["Articles"] = article
		valid := validation.Validation{}
		isValid, _ := valid.Valid(article)

		if manage.Ctx.Input.Method() == "POST" {
			if !isValid {
				manage.Data["Errors"] = valid.ErrorsMap
				beego.Error("Form didn't validate.")
			} else {
				searchArticle := models.Article{Name: article.Name}
				beego.Debug("Article name supplied:", article.Name)
				err = o.Read(&searchArticle)
				beego.Debug("Err:", err)
				flash := beego.NewFlash()

				if err == orm.ErrNoRows || err == orm.ErrMissPK {
					beego.Debug("No article found matching details supplied. Attempting to insert article: ", article)
					id, err := o.Insert(&article)
					if err == nil {
						msg := fmt.Sprintf("Article inserted with id:", id)
						beego.Debug(msg)
						flash.Notice(msg)
						flash.Store(&manage.Controller)
					} else {
						msg := fmt.Sprintf("Couldn't insert new article. Reason: ", err)
						beego.Debug(msg)
						flash.Error(msg)
						flash.Store(&manage.Controller)
					}
				} else {
					beego.Debug("Article found matching details supplied. Cannot insert")
				}
			}
		}
	}
}
