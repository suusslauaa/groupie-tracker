package web

import (
	"encoding/json"
	"io"
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

func (app *Application) GetResponse() (responseData []ResponseData, srverr error) {
	response, err := http.Get(app.Config.ArtistsURL)
	if err != nil {
		return nil, err
	}
	
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(bytes, &responseData); err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return responseData, nil
}

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		app.MethodNotAllowed(w)
		return
	}

	if r.URL.Path != "/" {
		app.NotFound(w)
		return
	}

	responseData, err := app.GetResponse()
	if err != nil {
		app.InternalServerError(w, err)
		return
	}

	if err = templates.ExecuteTemplate(w, "home.html", responseData); err != nil {
		app.InternalServerError(w, err)
		return
	}
}

func (app *Application) Artist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		app.MethodNotAllowed(w)
		return
	}

	responseData, err := app.GetResponse()
	if err != nil {
		app.InternalServerError(w, err)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id > len(responseData) || id <= 0 || r.URL.Query().Get("id")[0] == '0' || len(r.URL.Query()) != 1 {
		app.BadRequest(w)
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
		app.InternalServerError(w, err)
		return
	}

	bytes, err := io.ReadAll(relationsResponse.Body)
	if err != nil {
		app.InternalServerError(w, err)
		return
	}

	if err = json.Unmarshal(bytes, &artist); err != nil {
		app.InternalServerError(w, err)
		return
	}

	defer relationsResponse.Body.Close()

	if err = templates.ExecuteTemplate(w, "artist.html", artist); err != nil {
		app.InternalServerError(w, err)
		return
	}
}
