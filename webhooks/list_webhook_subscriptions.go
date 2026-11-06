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

type ListWebhookSubscriptionsRequest struct {
	PortfolioId     string `json:"-"`
	IncludeDisabled bool   `json:"-"`
	EventTypeFilter string `json:"-"`

	// Cursor is the NextCursor of a previous response.
	Cursor string `json:"-"`
}

// ListWebhookSubscriptionsResponse holds one page of subscriptions. Webhook list endpoints return
// flat next_cursor / has_next fields rather than the nested pagination object used elsewhere.
type ListWebhookSubscriptionsResponse struct {
	Subscriptions []*model.WebhookSubscription     `json:"subscriptions"`
	NextCursor    string                           `json:"next_cursor"`
	HasNext       bool                             `json:"has_next"`
	Request       *ListWebhookSubscriptionsRequest `json:"-"`
	service       WebhooksService
}

// Next fetches the next page of results. Returns nil, nil if there are no more pages.
func (r *ListWebhookSubscriptionsResponse) Next(ctx context.Context) (*ListWebhookSubscriptionsResponse, error) {
	if !r.HasNext || r.NextCursor == "" {
		return nil, nil
	}

	nextReq := *r.Request
	nextReq.Cursor = r.NextCursor

	return r.service.ListWebhookSubscriptions(ctx, &nextReq)
}

// ListWebhookSubscriptions lists the webhook subscriptions of a portfolio.
func (s *webhooksServiceImpl) ListWebhookSubscriptions(
	ctx context.Context,
	request *ListWebhookSubscriptionsRequest,
) (*ListWebhookSubscriptionsResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/webhooks/subscriptions", request.PortfolioId)

	queryParams := core.EmptyQueryParams

	if request.IncludeDisabled {
		queryParams = core.AppendHttpQueryParam(queryParams, "include_disabled", "true")
	}
	if request.EventTypeFilter != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "event_type_filter", request.EventTypeFilter)
	}
	if request.Cursor != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "cursor", request.Cursor)
	}

	response := &ListWebhookSubscriptionsResponse{Request: request, service: s}

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
