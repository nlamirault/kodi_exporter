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

// Request define the object sended to the server for the JSONRPC call
type Request struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	ID      int64       `json:"id"`
	Params  interface{} `json:"params,omitempty"`
}

// ErrorStack define sonme errors informations
type ErrorStack struct {
	Name     string      `json:"name,omitempty"`
	Type     string      `json:"type,omitempty"`
	Message  string      `json:"message,omitempty"`
	Property *ErrorStack `json:"property,omitempty"`
}

// ErrorData define some additional information about the error.
type ErrorData struct {
	Method string      `json:"method"`
	Stack  *ErrorStack `json:"stack"`
}

// ResponseError define the error member of the response when a rpc call encounters an error
type ResponseError struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    *ErrorData `json:"data,omitempty"`
}

// ResponseBase define the base JSON object response after a JSONRPC call
type ResponseBase struct {
	Jsonrpc string         `json:"jsonrpc,omitempty"`
	Method  string         `json:"method,omitempty"`
	ID      int64          `json:"id,omitempty"`
	Error   *ResponseError `json:"error,omitempty"`
}
