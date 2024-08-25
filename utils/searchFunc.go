package groupie

import (
	"net/http"
	"strconv"
	"strings"
)

func Search(query string, w http.ResponseWriter, r *http.Request) []int {
	query = strings.ToLower(query)
	var id []int
	queryint, _ := strconv.Atoi(query)
	for _, info := range AllArtists {
		info.Name = strings.ToLower(info.Name)
		if strings.Contains(info.Name, query) || info.CreationDate == queryint || strings.Contains(info.FirstAlbum, query) {
			id = append(id, info.ID-1)
			break
		}
		for _, member := range info.Members {
			member = strings.ToLower(member)
			if strings.Contains(member, query) {
				id = append(id, info.ID-1)
				break
			}
		}
	}
	for i := 0; i < len(AllLocations.Index); i++ {
		Loc := AllLocations.Index[i]
		for j := 0; j < len(Loc.Location); j++ {
			if strings.Contains(Loc.Location[j], query) {
				id = append(id, Loc.ID-1)
				break
			}
		}
	}
	return id
}
