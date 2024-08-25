package main

import (
	"fmt"
	"net/http"

	Handler "groupie/handlers"
	utils "groupie/utils"
)

func main() {
	utils.FetchAllData("https://groupietrackers.herokuapp.com/api")
	http.HandleFunc("/", Handler.Home)
	http.HandleFunc("/artists/{id}", Handler.Artist)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	fmt.Println("Listen to Localhost at Port 4000 http://localhost:4000")
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
