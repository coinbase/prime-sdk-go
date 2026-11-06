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

type GetWebhookOwnerStatsRequest struct {
	PortfolioId string                   `json:"-"`
	Period      model.WebhookStatsPeriod `json:"-"`
}

// GetWebhookOwnerStatsResponse contains webhook statistics aggregated across a portfolio's subscriptions.
// Counts and latencies are int64 values serialized as strings.
type GetWebhookOwnerStatsResponse struct {
	TotalSubscriptions  string                              `json:"total_subscriptions"`
	ActiveSubscriptions string                              `json:"active_subscriptions"`
	TotalEvents         string                              `json:"total_events"`
	OverallSuccessRate  float64                             `json:"overall_success_rate"`
	AvgLatencyMs        string                              `json:"avg_latency_ms"`
	Subscriptions       []*model.WebhookSubscriptionSummary `json:"subscriptions"`
	RecentFailures      []*model.WebhookRecentFailure       `json:"recent_failures"`
	Request             *GetWebhookOwnerStatsRequest        `json:"-"`
}

// GetWebhookOwnerStats retrieves webhook statistics for a portfolio.
func (s *webhooksServiceImpl) GetWebhookOwnerStats(
	ctx context.Context,
	request *GetWebhookOwnerStatsRequest,
) (*GetWebhookOwnerStatsResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/webhooks/stats", request.PortfolioId)

	queryParams := core.EmptyQueryParams

	if request.Period != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "period", string(request.Period))
	}

	response := &GetWebhookOwnerStatsResponse{Request: request}

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
