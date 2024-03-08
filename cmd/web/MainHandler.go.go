package web

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"text/template"
)

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	ImageURL     string   `json:"image"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Members      []string `json:"members"`
	Relations
}

type Relations struct {
	Rel map[string][]string `json:"datesLocations"`
}

const (
	urlArtists   = "https://groupietrackers.herokuapp.com/api/artists"
	urlRelations = "https://groupietrackers.herokuapp.com/api/relation/"
)

func (app *application) MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		app.ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		app.ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	Artists, err := GetArtists()
	if err != nil {
		app.ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, Artists)
	if err != nil {
		app.ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
}

func GetArtists() ([]Artist, error) {
	artistResp, err := http.Get(urlArtists)
	if err != nil {
		return nil, err
	}
	artists := []Artist{}

	defer artistResp.Body.Close()
	body, err := io.ReadAll(artistResp.Body)
	err = json.Unmarshal(body, &artists)
	if err != nil {
		return nil, err
	}
	return artists, nil
}

func GetArtist(ID int) (Artist, error) {
	var artist Artist
	artists, err := GetArtists()
	if err != nil {
		return artist, err
	}
	if ID > len(artists) {
		return artist, errors.New("ID does not exist")
	}
	artist = artists[ID-1]

	artist.Relations, err = GetRelation(ID)
	if err != nil {
		return artist, err
	}
	return artist, nil
}

func GetRelation(ID int) (Relations, error) {
	url := urlRelations + strconv.Itoa(ID)
	relResp, err := http.Get(url)
	if err != nil {
		return Relations{}, err
	}
	defer relResp.Body.Close()
	var artistsRel Relations
	err = json.NewDecoder(relResp.Body).Decode(&artistsRel)
	if err != nil {
		return Relations{}, err
	}
	return artistsRel, nil
}
