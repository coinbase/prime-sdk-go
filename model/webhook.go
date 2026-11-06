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

package model

// WebhookAuthType is the authentication scheme used when delivering webhooks.
type WebhookAuthType string

const (
	WebhookAuthTypeHmac      WebhookAuthType = "WEBHOOK_AUTH_TYPE_HMAC"
	WebhookAuthTypeOauth     WebhookAuthType = "WEBHOOK_AUTH_TYPE_OAUTH"
	WebhookAuthTypeMtls      WebhookAuthType = "WEBHOOK_AUTH_TYPE_MTLS"
	WebhookAuthTypeOauthMtls WebhookAuthType = "WEBHOOK_AUTH_TYPE_OAUTH_MTLS"
)

// WebhookDeliveryStatus is the delivery state of a webhook event.
type WebhookDeliveryStatus string

const (
	WebhookDeliveryStatusDelivered WebhookDeliveryStatus = "WEBHOOK_DELIVERY_STATUS_DELIVERED"
	WebhookDeliveryStatusRetrying  WebhookDeliveryStatus = "WEBHOOK_DELIVERY_STATUS_RETRYING"
	WebhookDeliveryStatusFailed    WebhookDeliveryStatus = "WEBHOOK_DELIVERY_STATUS_FAILED"
	WebhookDeliveryStatusPending   WebhookDeliveryStatus = "WEBHOOK_DELIVERY_STATUS_PENDING"
)

// WebhookHealthStatus is the health of a webhook subscription.
type WebhookHealthStatus string

const (
	WebhookHealthStatusHealthy  WebhookHealthStatus = "WEBHOOK_HEALTH_STATUS_HEALTHY"
	WebhookHealthStatusDegraded WebhookHealthStatus = "WEBHOOK_HEALTH_STATUS_DEGRADED"
	WebhookHealthStatusFailing  WebhookHealthStatus = "WEBHOOK_HEALTH_STATUS_FAILING"
	WebhookHealthStatusDisabled WebhookHealthStatus = "WEBHOOK_HEALTH_STATUS_DISABLED"
)

// WebhookErrorType classifies the outcome of a webhook delivery attempt.
type WebhookErrorType string

const (
	WebhookErrorTypeSuccess         WebhookErrorType = "WEBHOOK_ERROR_TYPE_SUCCESS"
	WebhookErrorTypeHttp4xx         WebhookErrorType = "WEBHOOK_ERROR_TYPE_HTTP_4XX"
	WebhookErrorTypeHttp5xx         WebhookErrorType = "WEBHOOK_ERROR_TYPE_HTTP_5XX"
	WebhookErrorTypeTimeout         WebhookErrorType = "WEBHOOK_ERROR_TYPE_TIMEOUT"
	WebhookErrorTypeConnectionError WebhookErrorType = "WEBHOOK_ERROR_TYPE_CONNECTION_ERROR"
)

// WebhookStatsPeriod is the time window for webhook statistics.
type WebhookStatsPeriod string

const (
	WebhookStatsPeriod24h WebhookStatsPeriod = "WEBHOOK_STATS_PERIOD_24H"
	WebhookStatsPeriod7d  WebhookStatsPeriod = "WEBHOOK_STATS_PERIOD_7D"
	WebhookStatsPeriod30d WebhookStatsPeriod = "WEBHOOK_STATS_PERIOD_30D"
)

// WebhookUpdateMask selects which fields an update applies to.
//
// The live Prime API decodes update_mask as an object with a paths array of snake_case field names,
// not the plain string shown in the OpenAPI spec. Supported paths: event_types, target_url,
// auth_type, oauth_config, custom_headers, description, is_enabled.
type WebhookUpdateMask struct {
	Paths []string `json:"paths"`
}

// WebhookOAuthConfigInput is the OAuth client configuration supplied when creating or updating a subscription.
type WebhookOAuthConfigInput struct {
	TokenUrl     string   `json:"token_url,omitempty"`
	ClientId     string   `json:"client_id,omitempty"`
	ClientSecret string   `json:"client_secret,omitempty"`
	Scopes       []string `json:"scopes,omitempty"`
}

// WebhookOAuthConfig is the OAuth client configuration returned on a subscription.
// The client secret is never returned.
type WebhookOAuthConfig struct {
	TokenUrl string   `json:"token_url,omitempty"`
	ClientId string   `json:"client_id,omitempty"`
	Scopes   []string `json:"scopes,omitempty"`
}

// WebhookSubscription is a webhook subscription for a portfolio.
type WebhookSubscription struct {
	SubscriptionId string              `json:"subscription_id,omitempty"`
	EventTypes     []string            `json:"event_types,omitempty"`
	TargetUrl      string              `json:"target_url,omitempty"`
	AuthType       WebhookAuthType     `json:"auth_type,omitempty"`
	OauthConfig    *WebhookOAuthConfig `json:"oauth_config,omitempty"`
	CustomHeaders  map[string]string   `json:"custom_headers,omitempty"`
	IsEnabled      bool                `json:"is_enabled"`
	Description    string              `json:"description,omitempty"`
	CreatedBy      string              `json:"created_by,omitempty"`
	UpdatedBy      string              `json:"updated_by,omitempty"`
	CreatedAt      string              `json:"created_at,omitempty"`
	UpdatedAt      string              `json:"updated_at,omitempty"`
	HealthStatus   WebhookHealthStatus `json:"health_status,omitempty"`

	// Optimistic-concurrency version (int64 serialized as a string). Pass it as ExpectedVersion on update.
	Version string `json:"version,omitempty"`
}

// WebhookEventType describes an event type that can be subscribed to.
type WebhookEventType struct {
	Name         string `json:"name,omitempty"`
	DomainObject string `json:"domain_object,omitempty"`
	Description  string `json:"description,omitempty"`
}

// WebhookDeliveryLogEntry is a single delivery record for a webhook event.
type WebhookDeliveryLogEntry struct {
	DeliveryId     string                `json:"delivery_id,omitempty"`
	EventId        string                `json:"event_id,omitempty"`
	EventType      string                `json:"event_type,omitempty"`
	DeliveryStatus WebhookDeliveryStatus `json:"delivery_status,omitempty"`
	StatusCode     int32                 `json:"status_code,omitempty"`
	LatencyMs      int32                 `json:"latency_ms,omitempty"`
	AttemptNumber  int32                 `json:"attempt_number,omitempty"`
	MaxRetries     int32                 `json:"max_retries,omitempty"`
	ErrorType      WebhookErrorType      `json:"error_type,omitempty"`
	ErrorMessage   string                `json:"error_message,omitempty"`
	IsTest         bool                  `json:"is_test,omitempty"`
	IsReplay       bool                  `json:"is_replay,omitempty"`
	OccurredAt     string                `json:"occurred_at,omitempty"`
	AttemptedAt    string                `json:"attempted_at,omitempty"`
}

// WebhookDeliveryAttempt is a single delivery attempt for a webhook event.
type WebhookDeliveryAttempt struct {
	DeliveryAttemptId string           `json:"delivery_attempt_id,omitempty"`
	DeliveryId        string           `json:"delivery_id,omitempty"`
	AttemptNumber     int32            `json:"attempt_number,omitempty"`
	StatusCode        int32            `json:"status_code,omitempty"`
	LatencyMs         int32            `json:"latency_ms,omitempty"`
	ErrorType         WebhookErrorType `json:"error_type,omitempty"`
	ErrorMessage      string           `json:"error_message,omitempty"`
	AttemptedAt       string           `json:"attempted_at,omitempty"`
}

// WebhookEventTypeStats is the delivery breakdown for a single event type.
// Counts are int64 values serialized as strings.
type WebhookEventTypeStats struct {
	EventType    string  `json:"event_type,omitempty"`
	Total        string  `json:"total,omitempty"`
	Successful   string  `json:"successful,omitempty"`
	SuccessRate  float64 `json:"success_rate,omitempty"`
	AvgLatencyMs string  `json:"avg_latency_ms,omitempty"`
}

// WebhookRecentFailure is a recent failed webhook delivery.
type WebhookRecentFailure struct {
	DeliveryId     string           `json:"delivery_id,omitempty"`
	SubscriptionId string           `json:"subscription_id,omitempty"`
	EventId        string           `json:"event_id,omitempty"`
	EventType      string           `json:"event_type,omitempty"`
	ErrorType      WebhookErrorType `json:"error_type,omitempty"`
	AttemptNumber  int32            `json:"attempt_number,omitempty"`
	AttemptedAt    string           `json:"attempted_at,omitempty"`
}

// WebhookSubscriptionSummary is a per-subscription summary in the owner stats.
// EventsInPeriod is an int64 value serialized as a string.
type WebhookSubscriptionSummary struct {
	SubscriptionId string              `json:"subscription_id,omitempty"`
	Description    string              `json:"description,omitempty"`
	HealthStatus   WebhookHealthStatus `json:"health_status,omitempty"`
	EventsInPeriod string              `json:"events_in_period,omitempty"`
	SuccessRate    float64             `json:"success_rate,omitempty"`
}
