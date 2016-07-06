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

func NewClient(address string, username string, password string) *Client {
	return &Client{
		URI:      address,
		Username: username,
		Password: password,
		Client:   &http.Client{},
	}

}

func (k *Client) performRequest(request interface{}) (*http.Response, error) {
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

func (k *Client) RPC(request interface{}, response interface{}) error {
	resp, err := k.performRequest(request)
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
	return err
}
