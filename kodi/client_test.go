// Copyright (C) 2016 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kodi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	// "net/http/httputil"
	"testing"

	"github.com/prometheus/common/log"
)

func Test_CantCreateKodiClientWithInvalidURI(t *testing.T) {
	client, err := NewClient("foo.bar", "", "")
	if err == nil {
		t.Fatalf("Invalid Kodi client uri: %v", client)
	}
}

func Test_CanCreateKodiClientWithAuthentication(t *testing.T) {
	client, err := NewClient("http://localhost:8080", "foo", "bar")
	if err != nil {
		t.Fatalf("Kodi client failed: %v", client)
	}
	if client.Username != "foo" || client.Password != "bar" {
		t.Fatalf("Kodi invalid authentication: %v", client)
	}
	if client.URI != "http://localhost:8080/jsonrpc" {
		t.Fatalf("Kodi invalud JSONRPC Uri: %s", client.URI)
	}
}

type kodiserver struct {
	*httptest.Server
}

func newKodiServer(req *Request) *kodiserver {
	h := &kodiserver{}
	h.Server = httptest.NewServer(handler(h, req))
	return h
}

func handler(ks *kodiserver, req *Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		dec := json.NewDecoder(bytes.NewBuffer(b))
		err = dec.Decode(req)
		log.Infof("%q", req)
		log.Infof("URI: %s %s", r.URL, req.Method)
		var resp string
		switch req.Method {
		case "JSONRPC.Ping":
			resp = `{"id":1,"jsonrpc":"2.0","result":"pong"}`
		case "GUI.ShowNotification":
			resp = `{"id":1,"jsonrpc":"2.0","result":"OK"}`
		case "VideoLibrary.GetMovies":
			resp = `{"id":1,"jsonrpc":"2.0","result":{"limits":{"end":3,"start":0,"total":3},"movies":[{"label":"108 Rois-Démons","movieid":1},{"label":"1001 pattes","movieid":2},{"label":"Aladdin","movieid":3}]}}`
		case "VideoLibrary.GetTVShows":
			resp = `{"id":1,"jsonrpc":"2.0","result":{"limits":{"end":4,"start":0,"total":4},"tvshows":[{"label":"Better Call Saul","tvshowid":1},{"label":"Star Wars - The Clone Wars","tvshowid":2},{"label":"Star Wars Rebels","tvshowid":3},{"label":"Deutschland 83","tvshowid":4}]}}`
		case "AudioLibrary.GetSongs":
			resp = `{"id":1,"jsonrpc":"2.0","result":{"limits":{"end":3095,"start":0,"total":3095},"songs":[{"label":"When the Going Gets Tough, the Tough Get Karazzee","songid":1},{"label":"Pardon My Freedom","songid":2},{"label":"Dear Can","songid":3},{"label":"King's Weed","songid":4}],"limits":{"end":4,"start":0,"total":4}}}`
		case "AudioLibrary.GetArtists":
			resp = `{"id":1,"jsonrpc":"2.0","result":{"artists":[{"artist":"!!!","artistid":1,"label":"!!!"},{"artist":"69","artistid":2,"label":"69"},{"artist":"ABBA","artistid":3,"label":"ABBA"},{"artist":"Adele","artistid":4,"label":"Adele"},{"artist":"Alain Souchon","artistid":5,"label":"Alain Souchon"},{"artist":"Alela Diane","artistid":6,"label":"Alela Diane"},{"artist":"Alpha Blondy","artistid":7,"label":"Alpha Blondy"}],"limits":{"end":7,"start":0,"total":7}}}`
		case "AudioLibrary.GetAlbums":
			resp = `{"id":1,"jsonrpc":"2.0","result":{"albums":[{"albumid":1,"label":"Louden Up Now"},{"albumid":2,"label":"Myth Takes"},{"albumid":3,"label":"[non-album tracks]"},{"albumid":4,"label":"Gold: Greatest Hits"},{"albumid":5,"label":"Rolling in the Deep"}],"limits":{"end":5,"start":0,"total":5}}}`
		}
		w.Write([]byte(resp))
	}
}

func getClientAndServer(t *testing.T, req *Request) (*kodiserver, *Client) {
	h := newKodiServer(req)
	client, err := NewClient(h.URL, "foo", "bar")
	if err != nil {
		t.Fatalf("%v", err)
	}
	return h, client
}

func TestKodiShowNotificationCall(t *testing.T) {
	req := &Request{}
	h, client := getClientAndServer(t, req)
	// h := newKodiServer(req)
	defer h.Close()

	// client, err := NewClient(h.URL, "foo", "bar")
	// if err != nil {
	// 	t.Fatalf("%v", err)
	// }

	resp, err := client.ShowNotification("unit", "test")
	if err != nil {
		t.Fatalf("%v", err)
	}
	log.Infof("Resp: %s", resp)
	if resp.Result != "OK" {
		t.Fatalf("Invalid GUI ShowNotification response: %s", resp)
	}
}

func TestKodiPingCall(t *testing.T) {
	req := &Request{}
	h, client := getClientAndServer(t, req)
	// h := newKodiServer(req)
	defer h.Close()

	// client, err := NewClient(h.URL, "foo", "bar")
	// if err != nil {
	// 	t.Fatalf("%v", err)
	// }

	resp, err := client.Ping()
	if err != nil {
		t.Fatalf("%v", err)
	}
	log.Infof("Resp: %s", resp)
	if resp.Result != "pong" {
		t.Fatalf("Invalid Ping response: %s", resp)
	}
}

func TestKodiGetMoviesCall(t *testing.T) {
	req := &Request{}
	h, client := getClientAndServer(t, req)
	// h := newKodiServer(req)
	defer h.Close()

	// client, err := NewClient(h.URL, "foo", "bar")
	// if err != nil {
	// 	t.Fatalf("%v", err)
	// }

	resp, err := client.VideoGetMovies()
	if err != nil {
		t.Fatalf("%v", err)
	}
	log.Infof("Resp: %s", resp)
	if resp.Result.Limits.End != 3 || resp.Result.Limits.Total != 3 {
		t.Fatalf("Invalid movies end: %s", resp)
	}
}

func TestKodiGetTVShowsCall(t *testing.T) {
	req := &Request{}
	h, client := getClientAndServer(t, req)
	// h := newKodiServer(req)
	defer h.Close()

	// client, err := NewClient(h.URL, "foo", "bar")
	// if err != nil {
	// 	t.Fatalf("%v", err)
	// }

	resp, err := client.VideoGetTVShows()
	if err != nil {
		t.Fatalf("%v", err)
	}
	log.Infof("Resp: %s", resp)
	if resp.Result.Limits.End != 4 || resp.Result.Limits.Total != 4 {
		t.Fatalf("Invalid tv shows end: %s", resp)
	}
}

func TestKodiGetArtistsCall(t *testing.T) {
	req := &Request{}
	h, client := getClientAndServer(t, req)
	defer h.Close()

	resp, err := client.AudioGetArtists()
	if err != nil {
		t.Fatalf("%v", err)
	}
	log.Infof("Resp: %s", resp)
	if resp.Result.Limits.End != 7 || resp.Result.Limits.Total != 7 {
		t.Fatalf("Invalid artists end: %s", resp)
	}
}

func TestKodiGetAlbumsCall(t *testing.T) {
	req := &Request{}
	h, client := getClientAndServer(t, req)
	defer h.Close()

	resp, err := client.AudioGetAlbums()
	if err != nil {
		t.Fatalf("%v", err)
	}
	log.Infof("Resp: %s", resp)
	if resp.Result.Limits.End != 5 || resp.Result.Limits.Total != 5 {
		t.Fatalf("Invalid albums end: %s", resp)
	}
}

func TestKodiGetSongsCall(t *testing.T) {
	req := &Request{}
	h, client := getClientAndServer(t, req)
	defer h.Close()

	resp, err := client.AudioGetSongs()
	if err != nil {
		t.Fatalf("%v", err)
	}
	log.Infof("Resp: %s", resp)
	if resp.Result.Limits.End != 4 || resp.Result.Limits.Total != 4 {
		t.Fatalf("Invalid songs end: %s", resp)
	}
}
