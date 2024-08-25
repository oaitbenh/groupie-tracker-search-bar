package groupie

import (
	"net/http"
	"text/template"

	utils "groupie/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if utils.AllArtists == nil {
		utils.FetchAllData("https://groupietrackers.herokuapp.com/api")
	}
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	Temp, err := template.ParseFiles("templates/home.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	search := r.FormValue("query")
	var (
		ArtistToShow []utils.ArtistInfo
		ids          []int
	)
	if len(search) != 0 {
		ids = utils.Search(search, w, r)
		for _, id := range ids {
			ArtistToShow = append(ArtistToShow, utils.AllArtists[id])
		}
	} else {
		ArtistToShow = utils.AllArtists
	}

	data := utils.ArtistSearch{
		Artist:    ArtistToShow,
		AllArtist: utils.AllArtists,
		Locations: utils.AllLocations.Index,
	}
	err = Temp.Execute(w, data)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}
