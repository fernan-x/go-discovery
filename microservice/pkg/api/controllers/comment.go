package controllers

import "net/http"

func getComments(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("comments"))
}
