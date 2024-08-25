package groupie

import "encoding/json"

func (Stru ApiUrls) FethData(Url string) ApiUrls {
	Json_Data := GetJson(Url)
	json.Unmarshal(Json_Data, &Stru)
	return Stru
}

func (Stru ArtistList) FethData(Url string) ArtistList {
	Json_Data := GetJson(Url)
	json.Unmarshal(Json_Data, &Stru)
	return Stru
}

func (Stru DateInfo) FethData(Url string) DateInfo {
	Json_Data := GetJson(Url)
	json.Unmarshal(Json_Data, &Stru)
	return Stru
}

func (Stru LocationInfo) FethData(Url string) LocationInfo {
	Json_Data := GetJson(Url)
	json.Unmarshal(Json_Data, &Stru)
	return Stru
}

func (Stru RelationInfo) FethData(Url string) RelationInfo {
	Json_Data := GetJson(Url)
	json.Unmarshal(Json_Data, &Stru)
	return Stru
}
