package delivery

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"groupie-tracker/internal/models"
)

var allArtists []models.Artist

func HomePage(page http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(page, http.StatusNotFound, ErrNotFound)
		return
	}
	if r.Method != http.MethodGet {
		ErrorHandler(page, http.StatusBadRequest, ErrBadRequest)
		return
	}
	response, _ := http.Get("https://groupietrackers.herokuapp.com/api")
	responseData, _ := ioutil.ReadAll(response.Body)
	var responseObject models.Response
	json.Unmarshal(responseData, &responseObject)

	s, _ := http.Get(responseObject.ArtistsUrl)
	d, _ := ioutil.ReadAll(s.Body)
	json.Unmarshal(d, &allArtists)
	// fmt.Println(allArtists)
	MainTmpl.Execute(page, allArtists)
}

func ArtistPage(page http.ResponseWriter, r *http.Request) {
	var (
		artist    models.Artist
		locations models.Locations
		dates     models.Dates
	)

	id := r.FormValue("button")
	response, _ := http.Get("https://groupietrackers.herokuapp.com/api")
	responseData, _ := ioutil.ReadAll(response.Body)

	var responseObject models.Response
	json.Unmarshal(responseData, &responseObject)

	artistInfo, _ := http.Get(responseObject.ArtistsUrl + "/" + id)
	d, _ := ioutil.ReadAll(artistInfo.Body)
	json.Unmarshal(d, &artist)

	// info.Image, info.Name, info.Members, info.CreationDate, info.FirstAlbum = artist.Image, artist.Name, artist.Members, artist.CreationDate, artist.FirstAlbum

	locationsInfo, _ := http.Get(artist.Locations)
	d, _ = ioutil.ReadAll(locationsInfo.Body)
	json.Unmarshal(d, &locations)

	datesInfo, _ := http.Get(locations.Dates)
	d, _ = ioutil.ReadAll(datesInfo.Body)
	json.Unmarshal(d, &dates)

	info := models.Info{
		Image:        artist.Image,
		Name:         artist.Name,
		Members:      artist.Members,
		CreationDate: artist.CreationDate,
		FirstAlbum:   artist.FirstAlbum,
		Locations:    locations.Locations,
		ConcertDates: dates.Dates,
	}

	// info.Locations = locations.Locations

	// fmt.Println(allArtists)
	ArtistTmpl.Execute(page, info)
}
