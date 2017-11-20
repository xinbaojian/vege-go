package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"vege-go/models"
	_ "vege-go/routers"
)

func main() {
	fmt.Println(models.ParseJson())
	beego.Run()
}
