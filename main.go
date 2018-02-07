package main

import (
	"net/http"
	"log"
	"Controller"
	"fmt"
	"strings"
	"io/ioutil"
)

func main() {
	
	router()
	err := http.ListenAndServe("127.0.0.1:9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func router(){
	http.HandleFunc("/user", Controller.IndexHandler)
	http.HandleFunc("/name", Controller.GetNameHandler)
}