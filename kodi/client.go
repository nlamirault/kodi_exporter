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

	"github.com/prometheus/common/log"
)

// Client defines the Kodi API client
type Client struct {
	URI      string
	Username string
	Password string
	Client   *http.Client
}

// NewClient defines a new client for the Kodi JSONRPC API
func NewClient(address string, username string, password string) *Client {
	return &Client{
		URI:      address,
		Username: username,
		Password: password,
		Client:   &http.Client{},
	}

}

func (k *Client) performRequest(request *Request) (*http.Response, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	log.Debugf("Kodi Request : %v\n", string(body))

	req, err := http.NewRequest("POST", k.URI, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(k.Username, k.Password)
	response, err := k.Client.Do(req)
	log.Debugf("Kodi Response : %v %v\n", response, err)
	return response, err
}

func (k *Client) rpc(method string, params interface{}, response interface{}) error {
	resp, err := k.performRequest(&Request{
		Jsonrpc: "2.0",
		Method:  method,
		ID:      1,
		Params:  params})
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Debugf("KODI Json : %v\n", string(b))
	dec := json.NewDecoder(bytes.NewBuffer(b))
	err = dec.Decode(response)
	// log.Debugf("KODI entity : %v\n", response)
	return err
}

// ShowNotification make a RPC call to shows a GUI notification
func (k *Client) ShowNotification(title string, message string) (*ShowNotificationResponse, error) {
	resp := &ShowNotificationResponse{}
	params := map[string]interface{}{
		`title`:   title,
		`message`: message,
	}
	err := k.rpc("GUI.ShowNotification", params, resp)
	return resp, err
}

// AudioGetArtists make a RPC call to retrieve all artists
func (k *Client) AudioGetArtists() (*AudioGetArtistsResponse, error) {
	resp := &AudioGetArtistsResponse{}
	params := map[string]interface{}{}
	err := k.rpc("AudioLibrary.GetArtists", params, resp)
	return resp, err
}

// AudioGetAlbums make a RPC call to retrieve all albums
func (k *Client) AudioGetAlbums() (*AudioGetAlbumsResponse, error) {
	resp := &AudioGetAlbumsResponse{}
	params := map[string]interface{}{}
	err := k.rpc("AudioLibrary.GetAlbums", params, resp)
	return resp, err
}

// AudioGetSongs make a RPC call to retrieve all songs
func (k *Client) AudioGetSongs() (*AudioGetSongsResponse, error) {
	resp := &AudioGetSongsResponse{}
	params := map[string]interface{}{}
	err := k.rpc("AudioLibrary.GetSongs", params, resp)
	return resp, err
}

// VideoGetMovies make a RPC call to retrieve all movies
func (k *Client) VideoGetMovies() (*VideoGetMoviesResponse, error) {
	resp := &VideoGetMoviesResponse{}
	params := map[string]interface{}{}
	err := k.rpc("VideoLibrary.GetMovies", params, resp)
	return resp, err
}

// VideoGetTVShows make a RPC call to retrieve all TV shows
func (k *Client) VideoGetTVShows() (*VideoGetTVShowsResponse, error) {
	resp := &VideoGetTVShowsResponse{}
	params := map[string]interface{}{}
	err := k.rpc("VideoLibrary.GetTVShows", params, resp)
	return resp, err
}

func (k *Client) videoGetGenres(videotype string) (*VideoGetGenresResponse, error) {
	resp := &VideoGetGenresResponse{}
	params := map[string]interface{}{
		`type`: videotype,
	}
	err := k.rpc("VideoLibrary.GetGenres", params, resp)
	return resp, err
}

// VideoGetTVShowsGenres make a RPC call to retrieve all genres for TV shows
func (k *Client) VideoGetTVShowsGenres() (*VideoGetGenresResponse, error) {
	return k.videoGetGenres("tvshow")
}

// VideoGetMoviesGenres make a RPC call to retrieve all genres for movies
func (k *Client) VideoGetMoviesGenres() (*VideoGetGenresResponse, error) {
	return k.videoGetGenres("movie")
}
