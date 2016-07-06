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

type Request struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Id      int64       `json:"id"`
	Params  interface{} `json:"params,omitempty"`
}

type ErrorStack struct {
	Name     string      `json:"name,omitempty"`
	Type     string      `json:"type,omitempty"`
	Message  string      `json:"message,omitempty"`
	Property *ErrorStack `json:"property,omitempty"`
}

type ErrorData struct {
	Method string      `json:"method"`
	Stack  *ErrorStack `json:"stack"`
}

type ResponseError struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    *ErrorData `json:"data,omitempty"`
}

type ResponseBase struct {
	Jsonrpc string         `json:"jsonrpc,omitempty"`
	Method  string         `json:"method,omitempty"`
	ID      int64          `json:"id,omitempty"`
	Error   *ResponseError `json:"error,omitempty"`
}

type ListLimitsReturned struct {
	Total int `json:"total"`
	Start int `json:"start,omitempty"`
	End   int `json:"end,omitempty"`
}

// type Result string

// GUI

type ShowNotificationResponse struct {
	ResponseBase
	Result string `json:"result,omitempty"`
}

// Audio Library

type Artist struct {
	Artist   string `json:"artist,omitempty"`
	ArtistID int    `json:"artistid"`
	Label    string `json:"label,omitempty"`
}

type ArtistsResponse struct {
	Artists []Artist            `json:"artists,omitempty"`
	Limits  *ListLimitsReturned `json:"limits,omitempty"`
}

type AudioGetArtistsResponse struct {
	ResponseBase
	Result ArtistsResponse `json:"result,omitempty"`
}

type Album struct {
	AlbumID int    `json:"albumid"`
	Label   string `json:"label,omitempty"`
}

type AlbumsResponse struct {
	Albums []Album             `json:"albums,omitempty"`
	Limits *ListLimitsReturned `json:"limits,omitempty"`
}

type AudioGetAlbumsResponse struct {
	ResponseBase
	Result AlbumsResponse `json:"result,omitempty"`
}

type Song struct {
	SongID int    `json:"songid"`
	Label  string `json:"label,omitempty"`
}

type SongsResponse struct {
	Songs  []Song              `json:"songs,omitempty"`
	Limits *ListLimitsReturned `json:"limits,omitempty"`
}

type AudioGetSongsResponse struct {
	ResponseBase
	Result SongsResponse `json:"result,omitempty"`
}

// Video Library

type TVShow struct {
	TVShowID int    `json:"tvshowid"`
	Label    string `json:"label,omitempty"`
}

type TVShowsResponse struct {
	TVShows []TVShow            `json:"tvshows,omitempty"`
	Limits  *ListLimitsReturned `json:"limits,omitempty"`
}

type VideoGetTVShowsResponse struct {
	ResponseBase
	Result TVShowsResponse `json:"result,omitempty"`
}

type Movie struct {
	MovieID int    `json:"movieid"`
	Label   string `json:"label,omitempty"`
}

type MoviessResponse struct {
	Movies []Movie             `json:"movies,omitempty"`
	Limits *ListLimitsReturned `json:"limits,omitempty"`
}

type VideoGetMoviesResponse struct {
	ResponseBase
	Result TVShowsResponse `json:"result,omitempty"`
}
