package Controller

import (
	"net/http"
	"fmt"
	"common/helper"
	"Model"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, helper.OutputJson(1000, "success", "123"))
}

func GetNameHandler(w http.ResponseWriter, r *http.Request) {
	info := Model.GetUser()
	fmt.Fprintln(w, helper.OutputJson(1000, "success",info))
}
