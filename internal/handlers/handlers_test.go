package handlers

import (
	"context"
	"learn-golang/internal/models"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	// {
	// 	name:               "home",
	// 	url:                "/",
	// 	method:             "GET",
	// 	params:             []postData{},
	// 	expectedStatusCode: http.StatusOK,
	// },
	// {
	// 	name:               "about",
	// 	url:                "/about",
	// 	method:             "GET",
	// 	params:             []postData{},
	// 	expectedStatusCode: http.StatusOK,
	// },
	// {
	// 	name:               "generals-quarters",
	// 	url:                "/generals-quarters",
	// 	method:             "GET",
	// 	params:             []postData{},
	// 	expectedStatusCode: http.StatusOK,
	// },
	// {
	// 	name:               "majors-suite",
	// 	url:                "/majors-suite",
	// 	method:             "GET",
	// 	params:             []postData{},
	// 	expectedStatusCode: http.StatusOK,
	// },
	// {
	// 	name:               "search-availability",
	// 	url:                "/search-availability",
	// 	method:             "GET",
	// 	params:             []postData{},
	// 	expectedStatusCode: http.StatusOK,
	// },
	// {
	// 	name:               "contact",
	// 	url:                "/contact",
	// 	method:             "GET",
	// 	params:             []postData{},
	// 	expectedStatusCode: http.StatusOK,
	// },
	// {
	// 	name:               "make-reservation",
	// 	url:                "/make-reservation",
	// 	method:             "GET",
	// 	params:             []postData{},
	// 	expectedStatusCode: http.StatusOK,
	// },
	// {
	// 	name:   "post-search-availability",
	// 	url:    "/search-availability",
	// 	method: "POST",
	// 	params: []postData{
	// 		{key: "start", value: "2020-01-01"},
	// 		{key: "end", value: "2020-01-01"},
	// 	},
	// 	expectedStatusCode: http.StatusOK,
	// },
	// {
	// 	name:   "post-search-availability-json",
	// 	url:    "/search-availability-json",
	// 	method: "POST",
	// 	params: []postData{
	// 		{key: "start", value: "2020-01-01"},
	// 		{key: "end", value: "2020-01-01"},
	// 	},
	// 	expectedStatusCode: http.StatusOK,
	// },
	// {
	// 	name:   "post-make-reservation",
	// 	url:    "/make-reservation",
	// 	method: "POST",
	// 	params: []postData{
	// 		{key: "first_name", value: "John"},
	// 		{key: "last_name", value: "Smith"},
	// 		{key: "email", value: "john.smith@gmail.com"},
	// 		{key: "phone", value: "555-444-666"},
	// 	},
	// 	expectedStatusCode: http.StatusOK,
	// },
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		testUrl := ts.URL + e.url

		var err error
		var resp *http.Response

		if e.method == "GET" {
			resp, err = ts.Client().Get(testUrl)
		} else if e.method == "POST" {
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}
			resp, err = ts.Client().PostForm(testUrl, values)
		}

		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if resp.StatusCode != e.expectedStatusCode {
			t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
		}
	}
}

func TestRepository_Reservation(t *testing.T) {
	reservation := models.Reservation{
		RoomID: 1,
		Room: models.Room{
			ID:        1,
			RoomName:  "General's Quarters",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	}

	req, _ := http.NewRequest("GET", "/make-reservation", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)

	handler := http.HandlerFunc(Repo.Reservation)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Reservation handler returns wrong status code: got %d, want %d", rr.Code, http.StatusOK)
	}
}

func getCtx(req *http.Request) context.Context {
	ctx, err := session.Load(req.Context(), req.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}
	return ctx
}
