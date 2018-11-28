package controllers

import (
	"net/http"
	"fmt"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}
