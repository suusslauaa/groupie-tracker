package web

import (
	"flag"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"testing"
)

func TestHome(t *testing.T) {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := NewApplication(errorLog, infoLog)

	// os.Chdir("../../")
	test_list := []struct {
		path   string
		method string
		code   int
	}{
		{
			path:   "/",
			method: "GET",
			code:   200,
		},
		{
			path:   "/main",
			method: "GET",
			code:   404,
		},
		{
			path:   "/",
			method: "DELETE",
			code:   405,
		},
		{
			path:   "/",
			method: "POST",
			code:   405,
		},
		{
			path:   "/",
			method: "PUT",
			code:   405,
		},
		{
			path:   "/Home",
			method: "POST",
			code:   405,
		},
		{
			path:   "/Home",
			method: "GET",
			code:   404,
		},
		{
			path:   "/main.go",
			method: "POST",
			code:   405,
		},
		{
			path:   "/artist",
			method: "GET",
			code:   404,
		},
	}
	for _, cases := range test_list {
		req := httptest.NewRequest(cases.method, cases.path, nil)
		w := httptest.NewRecorder()
		app.Home(w, req)
		resp := w.Result()
		if resp.StatusCode != cases.code {
			t.Errorf("Expected status %v; got %v", cases.code, resp.Status)
		}
	}
}

func TestArtists(t *testing.T) {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := NewApplication(errorLog, infoLog)

	test_list := []struct {
		path    string
		method  string
		code    int
		content string
	}{
		{
			path:    "/artists?id=1",
			method:  "GET",
			code:    200,
			content: "Queen",
		},
		{
			path:   "/artists?id=-1",
			method: "GET",
			code:   400,
		},
		{
			path:   "/artists?id=21",
			method: "DELETE",
			code:   405,
		},
		{
			path:   "/artists?id=2",
			method: "POST",
			code:   405,
		},
		{
			path:   "/artists?id=13",
			method: "PUT",
			code:   405,
		},
		{
			path:   "/artists?id=12&name=Akzhol",
			method: "GET",
			code:   400,
		},
		{
			path:   "/artists?id=715",
			method: "GET",
			code:   400,
		},
		{
			path:   "/artists?id=000001",
			method: "GET",
			code:   400,
		},
		{
			path:   "/artists?id=01",
			method: "GET",
			code:   400,
		},
		{
			path:   "/artists?id=17",
			method: "GET",
			code:   200,
			content: "Bee Gees",
		},
		{
			path:   "/artists?id=15",
			method: "DELETE",
			code:   405,
		},
		{
			path:   "/artists?id=LOOL",
			method: "GET",
			code:   400,
		},
		{
			path:   "/artists?id='A'",
			method: "GET",
			code:   400,
		},
		{
			path:   "/artists?id=12/3",
			method: "GET",
			code:   400,
		},
		{
			path:   "/artists?id===15",
			method: "GET",
			code:   400,
		},
		{
			path:   "/artists????id=15",
			method: "GET",
			code:   400,
		},
		{
			path:   "/artists?id=main.go",
			method: "GET",
			code:   400,
		},
	}
	for _, cases := range test_list {
		req, err := http.NewRequest(cases.method, cases.path, nil)
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()
		app.Artist(w, req)
		resp := w.Result()
		g, _ := io.ReadAll(resp.Body)

		if cases.content != "" {
			tr := regexp.MustCompile(cases.content).MatchString(string(g))
			if tr && resp.StatusCode != 200 {
				t.Errorf("status is %v; got %v", resp.StatusCode, cases.content)
			}
		}

		if resp.StatusCode != cases.code {
			t.Errorf("Expected status %v; got %v", cases.content, resp.StatusCode)
		}
	}
}

func TestServerError(t *testing.T) {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := NewApplication(errorLog, infoLog)

	w := httptest.NewRecorder()

	app.InternalServerError(w, nil)

	resp := w.Result()

	if resp.StatusCode != 500 {
		t.Errorf("Expected status %v; got %v", 500, resp.Status)
	}
}

func TestClientError(t *testing.T) {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := NewApplication(errorLog, infoLog)

	test_list := []struct {
		text string
		code int
	}{
		{
			text: "404 Not Found",
			code: 404,
		},
		{
			text: "400 Bad Request",
			code: 400,
		},
		{
			text: "405 Method Not Allowed",
			code: 405,
		},
	}
	for _, cases := range test_list {
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()
		app.ClientError(w, cases.code)
		resp := w.Result()
		if resp.Status != cases.text {
			t.Errorf("Expected status %v; got %v", cases.text, resp.Status)
		}
	}
}

func TestBadRequest(t *testing.T) {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := NewApplication(errorLog, infoLog)

	w := httptest.NewRecorder()

	app.BadRequest(w)

	resp := w.Result()

	if resp.StatusCode != 400 {
		t.Errorf("Expected status %v; got %v", 400, resp.Status)
	}
}

func TestMethodNotAllowed(t *testing.T) {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := NewApplication(errorLog, infoLog)

	w := httptest.NewRecorder()

	app.MethodNotAllowed(w)

	resp := w.Result()

	if resp.StatusCode != 405 {
		t.Errorf("Expected status %v; got %v", 405, resp.Status)
	}
}

func TestNotFound(t *testing.T) {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := NewApplication(errorLog, infoLog)

	w := httptest.NewRecorder()

	app.NotFound(w)

	resp := w.Result()

	if resp.StatusCode != 404 {
		t.Errorf("Expected status %v; got %v", 404, resp.Status)
	}
}

func TestErrors(t *testing.T) {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := NewApplication(errorLog, infoLog)

	test_list := []struct {
		text string
		code int
	}{
		{
			text: "404 Not Found",
			code: 404,
		},
		{
			text: "400 Bad Request",
			code: 400,
		},
		{
			text: "405 Method Not Allowed",
			code: 405,
		},
	}
	for _, cases := range test_list {
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()
		app.Errors(w, "", cases.code)
		resp := w.Result()
		if resp.Status != cases.text {
			t.Errorf("Expected status %v; got %v", cases.text, resp.Status)
		}
	}
}

func TestNewServer(t *testing.T) {
	addr := flag.String("addr", ":4000", "the network address of the web server")
	flag.Parse()

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := NewApplication(errorLog, infoLog)

	server := NewServer(addr, errorLog, app.Routes())

	if server.Addr != ":4000" {
		t.Errorf("Expected status\n%v", server.Addr)
	}
}

func TestNewApplication(t *testing.T) {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := NewApplication(errorLog, infoLog)

	if app.Config.ArtistsURL != "https://groupietrackers.herokuapp.com/api/artists" {
		t.Errorf("Expected status\n%v", app.Config.ArtistsURL)
	}
}

func TestGetResponse(t *testing.T) {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := NewApplication(errorLog, infoLog)

	resp, err := app.GetResponse()
	if err != nil {
		t.Errorf("internal Server Error while testing GetResponse func")
	}

	if resp[0].Name != "Queen" {
		t.Errorf("Expected status\n%v", resp[0].Name)
	}
}
