package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

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

func (app *Application) GetResponse(w http.ResponseWriter) (responseData []ResponseData) {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		app.errorLog.Println("Error getting data from API:", err)
		app.ServerError(w, fmt.Errorf("failed to fetch data from API"))

		return
	}

	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&responseData); err != nil {
		app.errorLog.Println("Error decoding JSON:", err)
		app.ServerError(w, fmt.Errorf("failed to decode JSON response"))
	}

	return
}

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.NotFound(w)
		return
	}

	if r.Method != http.MethodGet {
		app.MethodNotAllowed(w)
		return
	}

	responseData := app.GetResponse(w)

	err = templates.ExecuteTemplate(w, "home.html", responseData)
	if err != nil {
		app.ServerError(w, err)
	}
}

func (app *Application) Artist(w http.ResponseWriter, r *http.Request) {
	responseData := app.GetResponse(w)

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id > len(responseData) || id <= 0 || r.URL.Query().Get("id")[0] == '0' || len(r.URL.Query()) != 1 {
		app.NotFound(w)
		return
	}

	if r.Method != http.MethodGet {
		app.MethodNotAllowed(w)
		return
	}

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
		app.ServerError(w, fmt.Errorf("failed to fetch data from API"))

		return
	}

	defer relationsResponse.Body.Close()

	if err := json.NewDecoder(relationsResponse.Body).Decode(&artist); err != nil {
		app.errorLog.Println("Error decoding JSON:", err)
		app.ServerError(w, fmt.Errorf("failed to decode JSON response"))

		return
	}

	err = templates.ExecuteTemplate(w, "artist.html", artist)
	if err != nil {
		app.ServerError(w, err)
	}
}
