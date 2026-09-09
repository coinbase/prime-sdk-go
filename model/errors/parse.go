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

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ParseBody unmarshals a Prime REST API error JSON body.
func ParseBody(body []byte) (Response, error) {
	var resp Response
	if err := json.Unmarshal(body, &resp); err != nil {
		return Response{}, err
	}
	return resp, nil
}

// From extracts a Prime ApiError from err using errors.As.
func From(err error) (*ApiError, bool) {
	if err == nil {
		return nil, false
	}
	var apiErr *ApiError
	if errors.As(err, &apiErr) {
		return apiErr, true
	}
	return nil, false
}

// IsCode reports whether err is an ApiError with the given error code.
func IsCode(err error, code ErrorCode) bool {
	apiErr, ok := From(err)
	return ok && apiErr.Code == code
}

// IsSubcode reports whether err is an ApiError with the given subcode.
func IsSubcode(err error, sub Subcode) bool {
	apiErr, ok := From(err)
	return ok && apiErr.Subcode == sub
}

// Format returns a human-readable summary including spec description when known.
func (e *ApiError) Format() string {
	if e == nil {
		return ""
	}
	var b strings.Builder
	fmt.Fprintf(&b, "code=%s", e.Code)
	if e.Subcode != "" {
		fmt.Fprintf(&b, " subcode=%s", e.Subcode)
	}
	if e.Message != "" {
		fmt.Fprintf(&b, " message=%s", e.Message)
	}
	if info, ok := e.SubcodeInfo(); ok && info.Description != "" {
		fmt.Fprintf(&b, " description=%s", info.Description)
	} else if info, ok := LookupErrorCode(e.Code); ok && info.Description != "" {
		fmt.Fprintf(&b, " description=%s", info.Description)
	}
	if e.TraceID != "" {
		fmt.Fprintf(&b, " trace_id=%s", e.TraceID)
	}
	if e.StatusCode != 0 {
		fmt.Fprintf(&b, " status=%d", e.StatusCode)
	}
	if e.URL != "" {
		fmt.Fprintf(&b, " url=%s", e.URL)
	}
	return b.String()
}

// SubcodeInfo returns spec metadata for this error's subcode, if known.
func (e *ApiError) SubcodeInfo() (SubcodeInfo, bool) {
	if e == nil || e.Subcode == "" {
		return SubcodeInfo{}, false
	}
	return LookupSubcode(e.Subcode)
}

// HTTPStatus returns the HTTP status recorded on the error.
func (e *ApiError) HTTPStatus() int {
	if e == nil {
		return 0
	}
	return e.StatusCode
}

// Retryable reports whether the error is typically safe to retry after backoff.
func (e *ApiError) Retryable() bool {
	if e == nil {
		return false
	}
	switch e.Code {
	case ErrorCodeRateLimitExceeded, ErrorCodeServiceUnavailable:
		return true
	default:
		return false
	}
}
