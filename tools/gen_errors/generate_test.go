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

package main

import (
	"strings"
	"testing"
)

func TestParseSpecFixture(t *testing.T) {
	spec := []byte(`
openapi: 3.0.1
tags:
- name: PrimeRESTAPI
- name: Error Codes
  x-error-codes:
  - name: RESOURCE_NOT_FOUND
    httpStatus: 404
    description: The requested resource could not be found.
  - name: INTERNAL_ERROR
    httpStatus: 500
    description: "An unexpected internal error occurred while processing the request.\
      \ If this persists, contact support."
  x-subcodes:
  - name: ORDER_NOT_FOUND
    errorCode: RESOURCE_NOT_FOUND
    httpStatus: 404
    description: No order was found matching the provided order_id.
  - name: AUTH_UNAUTHENTICATED
    errorCode: AUTHENTICATION_FAILED
    httpStatus: 401
    description: The request could not be authenticated. Check that the credentials
      supplied with the request are valid.
paths:
  /v1/orders:
    get: {}
`)

	codes, subcodes, err := parseSpec(spec)
	if err != nil {
		t.Fatalf("parseSpec: %v", err)
	}
	if len(codes) != 2 {
		t.Fatalf("codes = %d, want 2", len(codes))
	}
	if codes[0].Name != "RESOURCE_NOT_FOUND" || codes[0].HTTPStatus != 404 {
		t.Fatalf("first code = %+v", codes[0])
	}
	if !strings.Contains(codes[1].Description, "unexpected internal error") {
		t.Fatalf("quoted description not unescaped: %q", codes[1].Description)
	}
	if len(subcodes) != 2 {
		t.Fatalf("subcodes = %d, want 2", len(subcodes))
	}
	if subcodes[1].Name != "AUTH_UNAUTHENTICATED" {
		t.Fatalf("second subcode = %+v", subcodes[1])
	}
	if !strings.Contains(subcodes[1].Description, "credentials supplied") {
		t.Fatalf("folded description not joined: %q", subcodes[1].Description)
	}
}

func TestDomainForSubcode(t *testing.T) {
	cases := map[string]string{
		"ORDER_NOT_FOUND":                     "order",
		"CROSS_MARGIN_NOT_ENABLED":            "financing",
		"ADDRESS_BOOK_NOT_FOUND":              "addressbook",
		"ADVANCED_TRANSFER_NOT_FOUND":         "advanced_transfer",
		"ONCHAIN_TRANSACTION_REQUEST_INVALID": "onchain",
		"FCM_EQUITY_NOT_FOUND":                "futures",
		"AUTH_UNAUTHENTICATED":                "common",
	}
	for name, want := range cases {
		got, err := domainForSubcode(name)
		if err != nil {
			t.Errorf("%s: %v", name, err)
			continue
		}
		if got != want {
			t.Errorf("%s: domain %q, want %q", name, got, want)
		}
	}
}

func TestGeneratedGodoc(t *testing.T) {
	out := renderSubcodes("order", []subcodeDef{{
		Name:        "ORDER_NOT_FOUND",
		ErrorCode:   "RESOURCE_NOT_FOUND",
		HTTPStatus:  404,
		Description: "No order was found matching the provided order_id.",
	}})
	if !strings.Contains(out, "// SubcodeOrderNotFound is the subcode ORDER_NOT_FOUND.") {
		t.Fatalf("missing hover godoc:\n%s", out)
	}
	if !strings.Contains(out, "// HTTP status: 404. Error code: RESOURCE_NOT_FOUND.") {
		t.Fatalf("missing status godoc:\n%s", out)
	}
}
