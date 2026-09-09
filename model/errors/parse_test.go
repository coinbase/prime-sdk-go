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
	"fmt"
	"strings"
	"testing"
)

func TestParseBodyAndFrom(t *testing.T) {
	body := []byte(`{"code":"VALIDATION_ERROR","message":"bad","subcode":"ORDER_REQUEST_INVALID","trace_id":"abc"}`)
	resp, err := ParseBody(body)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Code != ErrorCodeValidationError || resp.Subcode != SubcodeOrderRequestInvalid {
		t.Fatalf("parsed %+v", resp)
	}

	apiErr := &ApiError{Response: resp, StatusCode: 400, URL: "https://example/orders"}
	wrapped := fmt.Errorf("create order: %w", apiErr)
	got, ok := From(wrapped)
	if !ok || got.TraceID != "abc" {
		t.Fatalf("From = (%v, %v)", got, ok)
	}
	if !IsCode(wrapped, ErrorCodeValidationError) {
		t.Fatal("IsCode")
	}
	if !IsSubcode(wrapped, SubcodeOrderRequestInvalid) {
		t.Fatal("IsSubcode")
	}
	if apiErr.Retryable() {
		t.Fatal("validation should not be retryable")
	}
	gotFmt := apiErr.Format()
	for _, part := range []string{"VALIDATION_ERROR", "ORDER_REQUEST_INVALID", "abc"} {
		if !strings.Contains(gotFmt, part) {
			t.Fatalf("Format missing %q: %q", part, gotFmt)
		}
	}
}

func TestRetryable(t *testing.T) {
	if !(&ApiError{Response: Response{Code: ErrorCodeRateLimitExceeded}}).Retryable() {
		t.Fatal("429 should be retryable")
	}
	if !(&ApiError{Response: Response{Code: ErrorCodeServiceUnavailable}}).Retryable() {
		t.Fatal("503 should be retryable")
	}
}

func TestParseBodyUnauthorized(t *testing.T) {
	body := []byte(`{"code":"AUTHENTICATION_FAILED","message":"nope","subcode":"AUTH_UNAUTHENTICATED","trace_id":"t"}`)
	resp, err := ParseBody(body)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Subcode != SubcodeAuthUnauthenticated {
		t.Fatalf("subcode %s", resp.Subcode)
	}
}

func TestFromNil(t *testing.T) {
	if _, ok := From(nil); ok {
		t.Fatal("expected false")
	}
}
