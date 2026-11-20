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

package financing_test

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coinbase/prime-sdk-go/client"
	"github.com/coinbase/prime-sdk-go/credentials"
	"github.com/coinbase/prime-sdk-go/financing"
)

func TestGetCrossMarginPrimeOverviewSignsV2Path(t *testing.T) {
	const (
		accessKey  = "test-access-key"
		signingKey = "test-signing-key"
		wantPath   = "/v2/entities/test-entity/cross_margin/prime"
	)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != wantPath {
			t.Errorf("request = %s %s, want GET %s", r.Method, r.URL.Path, wantPath)
		}
		if got := r.Header.Get("X-CB-ACCESS-KEY"); got != accessKey {
			t.Errorf("access key = %q, want %q", got, accessKey)
		}
		timestamp := r.Header.Get("X-CB-ACCESS-TIMESTAMP")
		if timestamp == "" {
			t.Error("missing access timestamp")
		}
		mac := hmac.New(sha256.New, []byte(signingKey))
		_, _ = io.WriteString(mac, timestamp+http.MethodGet+wantPath)
		wantSignature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
		if got := r.Header.Get("X-CB-ACCESS-SIGNATURE"); got != wantSignature {
			t.Errorf("signature = %q, want %q for %s", got, wantSignature, wantPath)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, "{}")
	}))
	defer server.Close()

	v1BaseURL := server.URL + "/v1"
	cl := client.NewRestClient(&credentials.Credentials{
		AccessKey:  accessKey,
		SigningKey: signingKey,
	}, *server.Client()).SetBaseUrl(v1BaseURL)
	svc := financing.NewFinancingService(cl)
	_, err := svc.GetCrossMarginPrimeOverview(context.Background(), &financing.GetCrossMarginPrimeOverviewRequest{
		EntityId: "test-entity",
	})
	if err != nil {
		t.Fatal(err)
	}
	if got := cl.HttpBaseUrl(); got != v1BaseURL {
		t.Errorf("shared base URL = %q, want %q", got, v1BaseURL)
	}
}
