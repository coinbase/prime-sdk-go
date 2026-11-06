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

package webhooks

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coinbase/prime-sdk-go/client"
	"github.com/coinbase/prime-sdk-go/credentials"
	"github.com/coinbase/prime-sdk-go/model"
)

type capturedRequest struct {
	method string
	path   string
	query  string
	body   string
}

// newTestService returns a WebhooksService wired to an httptest server that records the
// last request and replies with the given JSON body.
func newTestService(t *testing.T, responseBody string) (WebhooksService, *capturedRequest) {
	t.Helper()

	captured := &capturedRequest{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		captured.method = r.Method
		captured.path = r.URL.Path
		captured.query = r.URL.RawQuery
		captured.body = string(b)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(responseBody))
	}))
	t.Cleanup(srv.Close)

	cl := client.NewRestClient(&credentials.Credentials{
		AccessKey:  "test-key",
		Passphrase: "test-pass",
		SigningKey: "test-signing-key",
	}, *srv.Client())
	cl.SetBaseUrl(srv.URL)

	return NewWebhooksService(cl), captured
}

func assertRequest(t *testing.T, got *capturedRequest, method, path, query string) {
	t.Helper()
	if got.method != method {
		t.Errorf("method = %s, want %s", got.method, method)
	}
	if got.path != path {
		t.Errorf("path = %s, want %s", got.path, path)
	}
	if got.query != query {
		t.Errorf("query = %q, want %q", got.query, query)
	}
}

func TestGetWebhookEventTypes(t *testing.T) {
	svc, got := newTestService(t, `{"event_types":[{"name":"ORDER_UPDATED","domain_object":"order","description":"d"}]}`)

	resp, err := svc.GetWebhookEventTypes(context.Background(), &GetWebhookEventTypesRequest{})
	if err != nil {
		t.Fatal(err)
	}
	assertRequest(t, got, http.MethodGet, "/webhooks/event-types", "")
	if len(resp.EventTypes) != 1 || resp.EventTypes[0].Name != "ORDER_UPDATED" {
		t.Errorf("unexpected response: %+v", resp)
	}
}

func TestListWebhookSubscriptions(t *testing.T) {
	svc, got := newTestService(t, `{"subscriptions":[{"subscription_id":"s1","is_enabled":true}],"next_cursor":"c2","has_next":true}`)

	resp, err := svc.ListWebhookSubscriptions(context.Background(), &ListWebhookSubscriptionsRequest{
		PortfolioId:     "p1",
		IncludeDisabled: true,
		EventTypeFilter: "ORDER_UPDATED",
		Cursor:          "c1",
	})
	if err != nil {
		t.Fatal(err)
	}
	assertRequest(t, got, http.MethodGet, "/portfolios/p1/webhooks/subscriptions",
		"include_disabled=true&event_type_filter=ORDER_UPDATED&cursor=c1")
	if !resp.HasNext || resp.NextCursor != "c2" || len(resp.Subscriptions) != 1 {
		t.Errorf("unexpected response: %+v", resp)
	}

	if _, err := resp.Next(context.Background()); err != nil {
		t.Fatal(err)
	}
	assertRequest(t, got, http.MethodGet, "/portfolios/p1/webhooks/subscriptions",
		"include_disabled=true&event_type_filter=ORDER_UPDATED&cursor=c2")
}

func TestListWebhookSubscriptionsNextWithoutMorePages(t *testing.T) {
	svc, _ := newTestService(t, `{"subscriptions":[],"has_next":false}`)

	resp, err := svc.ListWebhookSubscriptions(context.Background(), &ListWebhookSubscriptionsRequest{PortfolioId: "p1"})
	if err != nil {
		t.Fatal(err)
	}
	next, err := resp.Next(context.Background())
	if err != nil || next != nil {
		t.Errorf("Next() = %v, %v; want nil, nil", next, err)
	}
}

func TestCreateWebhookSubscription(t *testing.T) {
	svc, got := newTestService(t, `{"subscription":{"subscription_id":"s1"},"hmac_secret":"secret"}`)

	req := &CreateWebhookSubscriptionRequest{
		PortfolioId: "p1",
		EventTypes:  []string{"ORDER_UPDATED"},
		TargetUrl:   "https://example.com/hook",
		AuthType:    model.WebhookAuthTypeHmac,
	}
	resp, err := svc.CreateWebhookSubscription(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	assertRequest(t, got, http.MethodPost, "/portfolios/p1/webhooks/subscriptions", "")
	if req.IdempotencyKey == "" {
		t.Error("expected an idempotency key to be generated")
	}

	var body map[string]any
	if err := json.Unmarshal([]byte(got.body), &body); err != nil {
		t.Fatal(err)
	}
	if body["target_url"] != "https://example.com/hook" || body["auth_type"] != "WEBHOOK_AUTH_TYPE_HMAC" || body["idempotency_key"] == "" {
		t.Errorf("unexpected body: %s", got.body)
	}
	if _, ok := body["portfolio_id"]; ok {
		t.Error("portfolio_id must not be sent in the body")
	}
	if resp.HmacSecret != "secret" || resp.Subscription.SubscriptionId != "s1" {
		t.Errorf("unexpected response: %+v", resp)
	}
}

func TestGetWebhookSubscription(t *testing.T) {
	svc, got := newTestService(t, `{"subscription":{"subscription_id":"s1","version":"3"}}`)

	resp, err := svc.GetWebhookSubscription(context.Background(), &GetWebhookSubscriptionRequest{PortfolioId: "p1", SubscriptionId: "s1"})
	if err != nil {
		t.Fatal(err)
	}
	assertRequest(t, got, http.MethodGet, "/portfolios/p1/webhooks/subscriptions/s1", "")
	if resp.Subscription.Version != "3" {
		t.Errorf("unexpected response: %+v", resp)
	}
}

func TestUpdateWebhookSubscription(t *testing.T) {
	svc, got := newTestService(t, `{"subscription":{"subscription_id":"s1"}}`)

	disabled := false
	_, err := svc.UpdateWebhookSubscription(context.Background(), &UpdateWebhookSubscriptionRequest{
		PortfolioId:     "p1",
		SubscriptionId:  "s1",
		IsEnabled:       &disabled,
		UpdateMask:      "is_enabled",
		ExpectedVersion: "3",
	})
	if err != nil {
		t.Fatal(err)
	}
	assertRequest(t, got, http.MethodPatch, "/portfolios/p1/webhooks/subscriptions/s1", "")

	var body map[string]any
	if err := json.Unmarshal([]byte(got.body), &body); err != nil {
		t.Fatal(err)
	}
	if v, ok := body["is_enabled"]; !ok || v != false {
		t.Errorf("is_enabled=false must be sent explicitly, body: %s", got.body)
	}
	if body["update_mask"] != "is_enabled" || body["expected_version"] != "3" {
		t.Errorf("unexpected body: %s", got.body)
	}
}

func TestDeleteWebhookSubscription(t *testing.T) {
	svc, got := newTestService(t, `{}`)

	if _, err := svc.DeleteWebhookSubscription(context.Background(), &DeleteWebhookSubscriptionRequest{PortfolioId: "p1", SubscriptionId: "s1"}); err != nil {
		t.Fatal(err)
	}
	assertRequest(t, got, http.MethodDelete, "/portfolios/p1/webhooks/subscriptions/s1", "")
}

func TestRotateWebhookHmacSecret(t *testing.T) {
	svc, got := newTestService(t, `{"subscription":{"subscription_id":"s1"},"hmac_secret":"new","previous_secret_expires_at":"2026-01-01T00:00:00Z"}`)

	resp, err := svc.RotateWebhookHmacSecret(context.Background(), &RotateWebhookHmacSecretRequest{
		PortfolioId:               "p1",
		SubscriptionId:            "s1",
		RevokePreviousImmediately: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	assertRequest(t, got, http.MethodPost, "/portfolios/p1/webhooks/subscriptions/s1/rotate-hmac-secret", "")
	if got.body != `{"revoke_previous_immediately":true}` {
		t.Errorf("unexpected body: %s", got.body)
	}
	if resp.HmacSecret != "new" || resp.PreviousSecretExpiresAt == "" {
		t.Errorf("unexpected response: %+v", resp)
	}
}

func TestSendWebhookSimulation(t *testing.T) {
	svc, got := newTestService(t, `{"event_id":"e1"}`)

	resp, err := svc.SendWebhookSimulation(context.Background(), &SendWebhookSimulationRequest{
		PortfolioId:    "p1",
		SubscriptionId: "s1",
		EventType:      "ORDER_UPDATED",
	})
	if err != nil {
		t.Fatal(err)
	}
	assertRequest(t, got, http.MethodPost, "/portfolios/p1/webhooks/subscriptions/s1/simulate", "")
	if got.body != `{"event_type":"ORDER_UPDATED"}` || resp.EventId != "e1" {
		t.Errorf("unexpected body %s / response %+v", got.body, resp)
	}
}

func TestGetWebhookDeliveryLogs(t *testing.T) {
	svc, got := newTestService(t, `{"entries":[{"delivery_id":"d1","delivery_status":"WEBHOOK_DELIVERY_STATUS_FAILED","status_code":500}],"next_cursor":"n","has_next":true}`)

	resp, err := svc.GetWebhookDeliveryLogs(context.Background(), &GetWebhookDeliveryLogsRequest{
		PortfolioId:    "p1",
		SubscriptionId: "s1",
		StatusFilter:   model.WebhookDeliveryStatusFailed,
	})
	if err != nil {
		t.Fatal(err)
	}
	assertRequest(t, got, http.MethodGet, "/portfolios/p1/webhooks/subscriptions/s1/deliveries",
		"status_filter=WEBHOOK_DELIVERY_STATUS_FAILED")
	if len(resp.Entries) != 1 || resp.Entries[0].StatusCode != 500 || !resp.HasNext {
		t.Errorf("unexpected response: %+v", resp)
	}

	if _, err := resp.Next(context.Background()); err != nil {
		t.Fatal(err)
	}
	assertRequest(t, got, http.MethodGet, "/portfolios/p1/webhooks/subscriptions/s1/deliveries",
		"status_filter=WEBHOOK_DELIVERY_STATUS_FAILED&cursor=n")
}

func TestGetWebhookDeliveryAttempts(t *testing.T) {
	// "e30=" is base64 for "{}"
	svc, got := newTestService(t, `{"event_id":"e1","attempts":[{"attempt_number":1}],"event_payload":"e30="}`)

	resp, err := svc.GetWebhookDeliveryAttempts(context.Background(), &GetWebhookDeliveryAttemptsRequest{
		PortfolioId:    "p1",
		SubscriptionId: "s1",
		EventId:        "e1",
	})
	if err != nil {
		t.Fatal(err)
	}
	assertRequest(t, got, http.MethodGet, "/portfolios/p1/webhooks/subscriptions/s1/deliveries/e1/attempts", "")
	if string(resp.EventPayload) != "{}" || len(resp.Attempts) != 1 {
		t.Errorf("unexpected response: %+v", resp)
	}
}

func TestGetWebhookSubscriptionStats(t *testing.T) {
	svc, got := newTestService(t, `{"subscription_id":"s1","total_events":"10","success_rate":0.9}`)

	resp, err := svc.GetWebhookSubscriptionStats(context.Background(), &GetWebhookSubscriptionStatsRequest{
		PortfolioId:    "p1",
		SubscriptionId: "s1",
		Period:         model.WebhookStatsPeriod24h,
	})
	if err != nil {
		t.Fatal(err)
	}
	assertRequest(t, got, http.MethodGet, "/portfolios/p1/webhooks/subscriptions/s1/stats", "period=WEBHOOK_STATS_PERIOD_24H")
	if resp.TotalEvents != "10" || resp.SuccessRate != 0.9 {
		t.Errorf("unexpected response: %+v", resp)
	}
}

func TestGetWebhookOwnerStats(t *testing.T) {
	svc, got := newTestService(t, `{"total_subscriptions":"2","overall_success_rate":1}`)

	resp, err := svc.GetWebhookOwnerStats(context.Background(), &GetWebhookOwnerStatsRequest{
		PortfolioId: "p1",
		Period:      model.WebhookStatsPeriod30d,
	})
	if err != nil {
		t.Fatal(err)
	}
	assertRequest(t, got, http.MethodGet, "/portfolios/p1/webhooks/stats", "period=WEBHOOK_STATS_PERIOD_30D")
	if resp.TotalSubscriptions != "2" {
		t.Errorf("unexpected response: %+v", resp)
	}
}

func TestReplayWebhookEvent(t *testing.T) {
	svc, got := newTestService(t, `{"deliveries_queued":1,"delivery_ids":["d1"],"delivery_ids_truncated":false}`)

	resp, err := svc.ReplayWebhookEvent(context.Background(), &ReplayWebhookEventRequest{
		PortfolioId:    "p1",
		EventId:        "e1",
		SubscriptionId: "s1",
	})
	if err != nil {
		t.Fatal(err)
	}
	assertRequest(t, got, http.MethodPost, "/portfolios/p1/webhooks/events/e1/replay", "")
	if got.body != `{"subscription_id":"s1"}` || resp.DeliveriesQueued != 1 {
		t.Errorf("unexpected body %s / response %+v", got.body, resp)
	}
}
