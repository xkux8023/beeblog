package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
	"path"
	"strings"
)

type TopicController struct {
	BaseController
}

func (this *TopicController) Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true
	this.TplName = "topic.html"

	topics, err := models.GetAllTopics("", "", false)
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["Topics"] = topics
	}

}

func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	tid := this.Input().Get("tid")
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	category := this.Input().Get("category")
	label := this.Input().Get("label")

	//上传附件
	_, fh, err := this.GetFile("attachment")
	if err != nil {
		beego.Error(err)
	}

	var attachment string
	if fh != nil {
		//保存附件
		attachment = fh.Filename
		beego.Info(attachment)
		err = this.SaveToFile("attachment", path.Join("attachment", attachment))
		if err != nil {
			beego.Error(err)
		}
	}

	if len(tid) == 0 {
		err = models.AddTopic(title, content, category, label, attachment)
	} else {
		err = models.ModifyTopic(tid, title, content, category, label, attachment)
	}

	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 302)
}
func (this *TopicController) Add() {
	this.TplName = "topic_add.html"
	this.Data["CategoryList"], _ = models.GetAllCategories()
}

func (this *TopicController) View() {
	this.TplName = "topic_view.html"

	tid := this.Ctx.Input.Param("0")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Data["Topic"] = topic
	this.Data["Tid"] = tid
	this.Data["Labels"] = strings.Split(topic.Labels, " ")

	replies, err := models.GetAllReplies(tid)
	if err != nil {
		beego.Error(err)
		return
	}

	this.Data["Replies"] = replies
	this.Data["IsLogin"] = checkAccount(this.Ctx)
}

func (this *TopicController) Modify() {
	this.TplName = "topic_modify.html"
	tid := this.Input().Get("tid")

	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Data["Topic"] = topic
	this.Data["Tid"] = tid

	this.Data["CategoryList"], _ = models.GetAllCategories()
}

func (this *TopicController) Delete() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	err := models.DeleteTopic(this.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 302)
}
