package main

import (
	"groupie-tracker/web"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// func TestHomeHandler(t *testing.T) {
// 	t.Parallel()
// 	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
// 	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

// 	app := web.NewApplication(errorLog, infoLog)

// 	// var w http.ResponseWriter
// 	// var r *http.Request

// 	// fmt.Println(app.Home(w, r))

// 	s := httptest.NewServer(http.HandlerFunc(app.Home))
// 	req, err := http.NewRequest(http.MethodDelete, s.URL, nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// defer res.Body.Close()
// 	// body, err := io.ReadAll(res.Body)
// 	// if err != nil {
// 	// 	t.Fatal(err)
// 	// }
// 	want := 405

// 	if res.StatusCode != want {
// 		t.Errorf("Unexpected body returned. Want %d, got %d", want, res.StatusCode)
// 	}
// }

// func TestHomeHandler2(t *testing.T) {
// 	t.Parallel()
// 	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
// 	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

// 	app := web.NewApplication(errorLog, infoLog)

// 	// var w http.ResponseWriter
// 	// var r *http.Request

// 	// fmt.Println(app.Home(w, r))

// 	// s := httptest.NewServer(http.HandlerFunc(app.Home))
// 	req := httptest.NewRequest(http.MethodGet, "/", nil)

// 	recorder := httptest.NewRecorder()

// 	app.Home(recorder, req)

// 	resp := recorder.Result()
// 	defer resp.Body.Close()
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// res, err := http.DefaultClient.Do(req)
// 	// if err != nil {
// 	// 	t.Fatal(err)
// 	// }
// 	// defer res.Body.Close()
// 	// body, err := io.ReadAll(res.Body)
// 	// if err != nil {
// 	// 	t.Fatal(err)
// 	// }
// 	want := "Hello World!"

// 	if string(body) != want {
// 		t.Errorf("Unexpected body returned. Want %q, got %q", want, string(body))
// 	}
// }

// func TestServer(t *testing.T) {
// 	cmd := exec.Command("go", "run", ".")
// 	err := cmd.Start()
// 	if err != nil {
// 		t.Fatalf("Failed to start main program: %v", err)
// 	}
// 	defer func() {
// 		if err := cmd.Process.Kill(); err != nil {
// 			t.Fatalf("Failed to kill process: %v", err)
// 		}
// 	}()
// }

func TestHome(t *testing.T) {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := web.NewApplication(errorLog, infoLog)

	// os.Chdir("../../")
	testTable := []struct {
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
			path:   "/home",
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
	}
	for _, testcase := range testTable {
		req := httptest.NewRequest(testcase.method, testcase.path, nil)
		w := httptest.NewRecorder()
		app.Home(w, req)
		resp := w.Result()
		if resp.StatusCode != testcase.code {
			t.Errorf("expected status %v; got %v", testcase.code, resp.Status)
		}
	}
}
func TestArtists(t *testing.T) {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := web.NewApplication(errorLog, infoLog)

	testTable := []struct {
		path   string
		method string
		code   int
	}{
		{
			path:   "/artists?id=21",
			method: "GET",
			code:   200,
		},
		{
			path:   "/artists?id=-1",
			method: "GET",
			code:   404,
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
			code:   404,
		},
	}
	for _, testcase := range testTable {
		req, err := http.NewRequest(testcase.method, testcase.path, nil)
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()
		app.Artist(w, req)
		resp := w.Result()
		if resp.StatusCode != testcase.code {
			t.Errorf("expected status %v; got %v", testcase.code, resp.Status)
		}
	}
}

// func TestCheckId(t *testing.T) {
// 	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
// 	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

// 	app := web.NewApplication(errorLog, infoLog)

// 	tests := []struct {
// 		queryParam   string
// 		expectedId   int
// 		expectedCode int
// 	}{
// 		{
// 			queryParam:   "id=123",
// 			expectedId:   123,
// 			expectedCode: 0,
// 		},
// 		{
// 			queryParam:   "id=1",
// 			expectedId:   1,
// 			expectedCode: 0,
// 		},
// 		{
// 			queryParam:   "",
// 			expectedId:   0,
// 			expectedCode: http.StatusNotFound,
// 		},
// 		{
// 			queryParam:   "id=123&name=John",
// 			expectedId:   0,
// 			expectedCode: http.StatusNotFound,
// 		},
// 		{
// 			queryParam:   "id=abc",
// 			expectedId:   0,
// 			expectedCode: http.StatusNotFound,
// 		},
// 		{
// 			queryParam:   "id=0123",
// 			expectedId:   0,
// 			expectedCode: http.StatusNotFound,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run("", func(t *testing.T) {
// 			req, err := http.NewRequest("GET", "/artists?"+tt.queryParam, nil)
// 			if err != nil {
// 				t.Fatal(err)
// 			}
// 			id, code := app.Artist(req)
// 			if id != tt.expectedId {
// 				t.Errorf("Got ID %d, expected %d", id, tt.expectedId)
// 			}
// 			if code != tt.expectedCode {
// 				t.Errorf("Got code %d, expected %d", code, tt.expectedCode)
// 			}
// 		})
// 	}
// }
