package main

import (
	"fmt"
	"net/http"

	"github.com/eshun/video/models"
	"github.com/eshun/video/controllers"
	"strconv"
	"log"
)

func main() {
	//读取yaml配置文件
	conf, _ := models.GetConf()
	fmt.Printf("Your application is running here: http://localhost:%s\n", strconv.Itoa(conf.Port))


	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
	http.Handle("/", http.FileServer(http.Dir("./public")))

	http.HandleFunc("/video", controllers.VideoHandler)
	err := http.ListenAndServe(":"+strconv.Itoa(conf.Port), nil)
	if err != nil {
		log.Fatalf("ListenAndServe err: %v\n", err)
	}
}
