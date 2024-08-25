package groupie

import (
	"net/http"
	"text/template"

	utils "groupie/utils"
)

func ErrorHandler(w http.ResponseWriter, n int) {
	w.WriteHeader(n)
	template.New("Error")
	Temp, err := template.ParseFiles("./templates/error.html")
	if err != nil {
		http.Error(w, Errors[n].Title, n)
		return
	}
	err = Temp.Execute(w, Errors[n])
	if err != nil {
		http.Error(w, Errors[n].Title, n)
		return
	}
}

var Errors map[int]utils.ErrorData = map[int]utils.ErrorData{
	400: {
		StatusCode: 400,
		Title:      "Bad Request",
		Message:    "The server could not understand the request due to invalid syntax.",
	},
	404: {
		StatusCode: 404,
		Title:      "Page Not Found",
		Message:    "The page you are looking for might have been removed, had its name changed, or is temporarily unavailable.",
	},
	405: {
		StatusCode: 405,
		Title:      "Method Not Allowed",
		Message:    "The method specified in the request is not allowed for the resource identified by the request URI.",
	},
	500: {
		StatusCode: 500,
		Title:      "Internal Server Error",
		Message:    "The server encountered an unexpected condition that prevented it from fulfilling the request.",
	},
}
