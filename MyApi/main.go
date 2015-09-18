package main

import (
	_ "Users/dapeng/Desktop/temp/GoTest/MyApi/docs"
	_ "Users/dapeng/Desktop/temp/GoTest/MyApi/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
