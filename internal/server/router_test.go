package server

//nolint

import (
	"encoding/json"
	"github.com/Ladence/golang_base_kubernetes/internal/api"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	r := NewRouter()
	testServer := httptest.NewServer(r)
	defer testServer.Close()

	res, err := http.Get(testServer.URL + "/home")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code for /home is wrong. Have: %v, want: %v", res.StatusCode, http.StatusOK)
	}

	res, err = http.Post(testServer.URL+"/home", "text/plain", nil)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status code for POST /home is wrong. Have: %v, want: %v", res.StatusCode, http.StatusMethodNotAllowed)
	}

	res, err = http.Get(testServer.URL + "/not-exist")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Status code for GET /not-exist is wrong. Have: %v, want: %v", res.StatusCode, http.StatusNotFound)
	}
}

func TestHomeHandler(t *testing.T) {
	w := httptest.NewRecorder()
	home(w, nil)

	resp := w.Result()
	if have, want := resp.StatusCode, http.StatusOK; have != want {
		t.Errorf("Status code is wrong. Have: %v, want: %v", have, want)
	}

	r := &api.HomeResponse{}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(bytes, r)
	if err != nil {
		t.Fatal(err)
	}
}
