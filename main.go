package main

import (
	"beeblog/models"
	_ "beeblog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
)

func init() {
	models.RegisterDB()
	//附件上传文件夹处理
	os.Mkdir("attachment", os.ModePerm)
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)

	beego.Run()
}
