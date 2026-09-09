/**
 * Copyright 2026-present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package errors

import "fmt"

// ErrorCode is a Prime REST API error code from the OpenAPI x-error-codes list.
type ErrorCode string

// Subcode is a Prime REST API subcode from the OpenAPI x-subcodes list.
type Subcode string

// Response is the JSON error body returned by the Prime REST API.
type Response struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Subcode Subcode   `json:"subcode,omitempty"`
	TraceID string    `json:"trace_id,omitempty"`
}

// ApiError is the typed Prime REST error returned by client.Http* helpers.
type ApiError struct {
	Response
	// StatusCode is the HTTP status received from the API.
	StatusCode int
	// URL is the request URL that produced the error.
	URL string
}

func (e *ApiError) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf(
		"prime api error: code=%s subcode=%s message=%s status=%d url=%s",
		e.Code, e.Subcode, e.Message, e.StatusCode, e.URL,
	)
}

// ErrorCodeInfo is metadata for a global error code from the OpenAPI spec.
type ErrorCodeInfo struct {
	Code        ErrorCode
	HTTPStatus  int
	Description string
}

// SubcodeInfo is metadata for a subcode from the OpenAPI spec.
type SubcodeInfo struct {
	Subcode     Subcode
	ErrorCode   ErrorCode
	HTTPStatus  int
	Description string
}

var (
	errorCodeInfo map[ErrorCode]ErrorCodeInfo
	subcodeInfo   map[Subcode]SubcodeInfo
)

// LookupErrorCode returns spec metadata for an error code.
func LookupErrorCode(code ErrorCode) (ErrorCodeInfo, bool) {
	info, ok := errorCodeInfo[code]
	return info, ok
}

// LookupSubcode returns spec metadata for a subcode.
func LookupSubcode(sub Subcode) (SubcodeInfo, bool) {
	info, ok := subcodeInfo[sub]
	return info, ok
}
