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
	"fmt"

	"github.com/coinbase/core-go"
	"github.com/coinbase/prime-sdk-go/client"
	"github.com/coinbase/prime-sdk-go/model"
)

type GetWebhookSubscriptionStatsRequest struct {
	PortfolioId    string                   `json:"-"`
	SubscriptionId string                   `json:"-"`
	Period         model.WebhookStatsPeriod `json:"-"`
}

// GetWebhookSubscriptionStatsResponse contains delivery statistics for a subscription.
// Counts and latencies are int64 values serialized as strings.
type GetWebhookSubscriptionStatsResponse struct {
	SubscriptionId     string                              `json:"subscription_id"`
	HealthStatus       model.WebhookHealthStatus           `json:"health_status"`
	TotalEvents        string                              `json:"total_events"`
	Successful         string                              `json:"successful"`
	Failed             string                              `json:"failed"`
	SuccessRate        float64                             `json:"success_rate"`
	AvgLatencyMs       string                              `json:"avg_latency_ms"`
	P95LatencyMs       string                              `json:"p95_latency_ms"`
	P99LatencyMs       string                              `json:"p99_latency_ms"`
	EventTypeBreakdown []*model.WebhookEventTypeStats      `json:"event_type_breakdown"`
	RecentFailures     []*model.WebhookRecentFailure       `json:"recent_failures"`
	Request            *GetWebhookSubscriptionStatsRequest `json:"-"`
}

// GetWebhookSubscriptionStats retrieves delivery statistics for a webhook subscription.
func (s *webhooksServiceImpl) GetWebhookSubscriptionStats(
	ctx context.Context,
	request *GetWebhookSubscriptionStatsRequest,
) (*GetWebhookSubscriptionStatsResponse, error) {

	path := fmt.Sprintf(
		"/portfolios/%s/webhooks/subscriptions/%s/stats",
		request.PortfolioId,
		request.SubscriptionId,
	)

	queryParams := core.EmptyQueryParams

	if request.Period != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "period", string(request.Period))
	}

	response := &GetWebhookSubscriptionStatsResponse{Request: request}

	if err := client.HttpGet(
		ctx,
		s.client,
		path,
		queryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		response,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return response, nil
}
