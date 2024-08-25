package groupie

import (
	"net/http"
	"strconv"
	"text/template"

	utils "groupie/utils"
)

func Artist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 || id > len(utils.AllArtists) {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	Temp, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	data := utils.ArtistDetails{
		Artist:    utils.AllArtists[id-1],
		Locations: utils.AllLocations.Index[id-1].Location,
		Dates:     utils.AllDates.Index[id-1].Date,
		Relations: utils.AllRelations.Index[id-1].DateLocation,
	}
	err = Temp.Execute(w, data)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}
