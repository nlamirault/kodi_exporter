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

// type Result string

type ShowNotificationResponse struct {
	ResponseBase
	Result string `json:"result,omitempty"`
}
