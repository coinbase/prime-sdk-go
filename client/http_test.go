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

package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coinbase/core-go"
	"github.com/coinbase/prime-sdk-go/credentials"
	apierrors "github.com/coinbase/prime-sdk-go/model/errors"
)

func testRestClient(t *testing.T, handler http.HandlerFunc) (RestClient, *httptest.Server) {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	cl := NewRestClient(&credentials.Credentials{
		AccessKey:  "test-key",
		Passphrase: "test-pass",
		SigningKey: "test-signing-key",
	}, *srv.Client())
	cl.SetBaseUrl(srv.URL)
	return cl, srv
}

func TestHttpGetParsesPrimeError(t *testing.T) {
	cl, _ := testRestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"code":     "RESOURCE_NOT_FOUND",
			"message":  "order missing",
			"subcode":  "ORDER_NOT_FOUND",
			"trace_id": "trc-1",
		})
	})

	var resp struct {
		OK bool `json:"ok"`
	}
	err := HttpGet(context.Background(), cl, "/orders/1", "", DefaultSuccessHttpStatusCodes, struct{}{}, &resp, cl.HeadersFunc())
	if err == nil {
		t.Fatal("expected error")
	}
	apiErr, ok := apierrors.From(err)
	if !ok {
		t.Fatalf("From() = false, err=%v (%T)", err, err)
	}
	if apiErr.Code != apierrors.ErrorCodeResourceNotFound {
		t.Errorf("code = %s", apiErr.Code)
	}
	if apiErr.Subcode != apierrors.SubcodeOrderNotFound {
		t.Errorf("subcode = %s", apiErr.Subcode)
	}
	if apiErr.TraceID != "trc-1" {
		t.Errorf("trace_id = %s", apiErr.TraceID)
	}
	if apiErr.StatusCode != http.StatusNotFound {
		t.Errorf("status = %d", apiErr.StatusCode)
	}
	if !apierrors.IsSubcode(err, apierrors.SubcodeOrderNotFound) {
		t.Error("IsSubcode failed")
	}
}

func TestHttpGetMalformedErrorBody(t *testing.T) {
	cl, _ := testRestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
		_, _ = w.Write([]byte("not-json"))
	})

	var resp struct{}
	err := HttpGet(context.Background(), cl, "/x", "", DefaultSuccessHttpStatusCodes, struct{}{}, &resp, cl.HeadersFunc())
	apiErr, ok := apierrors.From(err)
	if !ok {
		t.Fatalf("From() = false, err=%v", err)
	}
	if apiErr.Message != "not-json" {
		t.Errorf("message = %q", apiErr.Message)
	}
}

func TestHttpGetSuccess(t *testing.T) {
	cl, _ := testRestClient(t, func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	var resp struct {
		OK bool `json:"ok"`
	}
	if err := HttpGet(context.Background(), cl, "/ok", core.EmptyQueryParams, DefaultSuccessHttpStatusCodes, struct{}{}, &resp, cl.HeadersFunc()); err != nil {
		t.Fatal(err)
	}
	if !resp.OK {
		t.Fatal("expected ok")
	}
}
