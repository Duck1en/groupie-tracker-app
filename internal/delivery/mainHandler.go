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
	response, err := http.Get("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		ErrorHandler(page, http.StatusInternalServerError, ErrServer)
		return
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		ErrorHandler(page, http.StatusInternalServerError, ErrServer)
		return
	}
	var responseObject models.Response
	json.Unmarshal(responseData, &responseObject)

	s, err := http.Get(responseObject.ArtistsUrl)
	if err != nil {
		ErrorHandler(page, http.StatusInternalServerError, ErrServer)
		return
	}
	d, err := ioutil.ReadAll(s.Body)
	if err != nil {
		ErrorHandler(page, http.StatusInternalServerError, ErrServer)
		return
	}
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
	response, err := http.Get("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		ErrorHandler(page, http.StatusInternalServerError, ErrServer)
		return
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		ErrorHandler(page, http.StatusInternalServerError, ErrServer)
		return
	}
	var responseObject models.Response
	json.Unmarshal(responseData, &responseObject)

	artistInfo, err := http.Get(responseObject.ArtistsUrl + "/" + id)
	if err != nil {
		ErrorHandler(page, http.StatusInternalServerError, ErrServer)
		return
	}

	d, err := ioutil.ReadAll(artistInfo.Body)
	if err != nil {
		ErrorHandler(page, http.StatusInternalServerError, ErrServer)
		return
	}
	json.Unmarshal(d, &artist)

	// info.Image, info.Name, info.Members, info.CreationDate, info.FirstAlbum = artist.Image, artist.Name, artist.Members, artist.CreationDate, artist.FirstAlbum

	locationsInfo, err := http.Get(artist.Locations)
	if err != nil {
		ErrorHandler(page, http.StatusNotFound, ErrNotFound)
		return
	}
	d, err = ioutil.ReadAll(locationsInfo.Body)
	if err != nil {
		ErrorHandler(page, http.StatusInternalServerError, ErrServer)
		return
	}
	json.Unmarshal(d, &locations)

	datesInfo, err := http.Get(locations.Dates)
	if err != nil {
		ErrorHandler(page, http.StatusInternalServerError, ErrServer)
		return
	}
	d, err = ioutil.ReadAll(datesInfo.Body)
	if err != nil {
		ErrorHandler(page, http.StatusInternalServerError, ErrServer)
		return
	}
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
