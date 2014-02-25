package main

import (
		_ "<%= baseName %>/routers"
		"github.com/astaxie/beego"
)

func main(){
	beego.Run()
}