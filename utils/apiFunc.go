package groupie

import (
	"io"
	"net/http"
)

func FetchAllData(Url string) {
	MainApi = MainApi.FethData(Url)
	AllArtists = AllArtists.FethData(MainApi.Artists)
	AllDates = AllDates.FethData(MainApi.Dates)
	AllLocations = AllLocations.FethData(MainApi.Locations)
	AllRelations = AllRelations.FethData(MainApi.Relations)
}

func GetJson(Url string) []byte {
	Response_Data, err := http.Get(Url)
	if err != nil {
		ErrorBeforeRun = err
		return nil
	}
	Json_Data, err := io.ReadAll(Response_Data.Body)
	if err != nil {
		ErrorBeforeRun = err
		return nil
	}
	return Json_Data
}
