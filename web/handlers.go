package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	if err = templates.ExecuteTemplate(w, "home.html", app.getResponse(w)); err != nil {
		app.serverError(w, err)
	}
}

func (app *application) artist(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.notFound(w)
		return
	}

	responseData := app.getResponse(w)

	artist := Artist{
		Name:         responseData[id-1].Name,
		Image:        responseData[id-1].Image,
		CreationDate: responseData[id-1].CreationDate,
		FirstAlbum:   responseData[id-1].FirstAlbum,
		Members:      responseData[id-1].Members,
	}

	relationsResponse, err := http.Get(responseData[id-1].Relations)
	if err != nil {
		app.errorLog.Println("Error getting data from API:", err)
		app.serverError(w, fmt.Errorf("failed to fetch data from API"))

		return
	}

	defer relationsResponse.Body.Close()

	if err := json.NewDecoder(relationsResponse.Body).Decode(&artist); err != nil {
		app.errorLog.Println("Error decoding JSON:", err)
		app.serverError(w, fmt.Errorf("failed to decode JSON response"))

		return
	}

	if err = templates.ExecuteTemplate(w, "artist.html", artist); err != nil {
		app.serverError(w, err)
	}
}

func (app *application) getResponse(w http.ResponseWriter) (responseData []ResponseData) {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		app.errorLog.Println("Error getting data from API:", err)
		app.serverError(w, fmt.Errorf("failed to fetch data from API"))

		return
	}

	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&responseData); err != nil {
		app.errorLog.Println("Error decoding JSON:", err)
		app.serverError(w, fmt.Errorf("failed to decode JSON response"))
	}

	return
}

type ResponseData struct {
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Id           float64  `json:"id"`
	CreationDate float64  `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Members      []string `json:"members"`
	Relations    string   `json:"relations"`
}

type Artist struct {
	Name         string
	Image        string
	CreationDate float64
	FirstAlbum   string
	Members      []string
	Relations    map[string][]string `json:"datesLocations"`
}
