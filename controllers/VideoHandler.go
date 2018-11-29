package controllers

import (
	"fmt"
	"github.com/eshun/video/models"
	"github.com/eshun/video/utils/request"
	"net/http"
)

func VideoHandler(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	url := queryValues.Get("url")
	if len(url) == 0 || string(url) == "" {
		fmt.Printf("param url not exist\n")
	}
	route := queryValues.Get("route")
	if len(route) == 0 || string(route) == "" {
		fmt.Printf("param route not exist\n")
	}
	conf, _ := models.GetConf()

	for _, fvip := range conf.Fvip {
		//fmt.Printf("Name:%v Url:%v \n\r", fvip.Name, fvip.Url)
		if fvip.Name == string(route) {
			body := request.Get(fvip.Url)
			fmt.Println(body.Html())
		}
	}

	w.Write([]byte("ok"))
}
