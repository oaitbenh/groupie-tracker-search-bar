package groupie

type ApiUrls struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relations string `json:"relation"`
}

type ArtistList []ArtistInfo

type ArtistInfo struct {
	ID           int      `json:"id"`
	ImageURL     string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type LocationInfo struct {
	Index []struct {
		ID       int      `json:"id"`
		Location []string `json:"locations"`
	} `json:"index"`
}

type DateInfo struct {
	Index []struct {
		ID   int      `json:"id"`
		Date []string `json:"dates"`
	} `json:"index"`
}

type RelationInfo struct {
	Index []struct {
		ID           int                 `json:"id"`
		DateLocation map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

type ArtistDetails struct {
	Artist    ArtistInfo
	Locations []string
	Dates     []string
	Relations map[string][]string
}

type ArtistSearch struct {
	Artist    []ArtistInfo
	AllArtist []ArtistInfo
	Locations []struct {
		ID       int      `json:"id"`
		Location []string `json:"locations"`
	}
}

type ErrorData struct {
	StatusCode int
	Title      string
	Message    string
}
