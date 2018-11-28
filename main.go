package main

import (
	"fmt"
	"net/http"

	"./models"
	"./controllers"
	"strconv"
)

func main() {
	fmt.Println("Starting....")

	//读取yaml配置文件
	conf,_ := models.GetConf()

	http.HandleFunc("/", controllers.IndexHandler)
	http.ListenAndServe(":" + strconv.Itoa(conf.Port), nil)
}

