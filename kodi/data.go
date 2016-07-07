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

// ListLimitsReturned define a Kodi entity for list informations
type ListLimitsReturned struct {
	Total int `json:"total"`
	Start int `json:"start,omitempty"`
	End   int `json:"end,omitempty"`
}

// type Result string

// GUI

// ShowNotificationResponse define a response after a ShowNotification RPC call
type ShowNotificationResponse struct {
	ResponseBase
	Result string `json:"result,omitempty"`
}

// Audio Library

// Artist define the Kodi artist entity
type Artist struct {
	Artist   string `json:"artist,omitempty"`
	ArtistID int    `json:"artistid"`
	Label    string `json:"label,omitempty"`
}

// ArtistsResponse define the Kodi artists list response
type ArtistsResponse struct {
	Artists []Artist            `json:"artists,omitempty"`
	Limits  *ListLimitsReturned `json:"limits,omitempty"`
}

// AudioGetArtistsResponse define the response to the GetArtists RPC call
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

type MoviesResponse struct {
	Movies []Movie             `json:"movies,omitempty"`
	Limits *ListLimitsReturned `json:"limits,omitempty"`
}

type VideoGetMoviesResponse struct {
	ResponseBase
	Result MoviesResponse `json:"result,omitempty"`
}

type Genre struct {
	GenreID int    `json:"genreid"`
	Label   string `json:"label,omitempty"`
}

type GenresResponse struct {
	Genres []Genre             `json:"genres,omitempty"`
	Limits *ListLimitsReturned `json:"limits,omitempty"`
}

type VideoGetGenresResponse struct {
	ResponseBase
	Result GenresResponse `json:"result,omitempty"`
}
