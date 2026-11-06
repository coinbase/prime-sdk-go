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

type GetWebhookDeliveryAttemptsRequest struct {
	PortfolioId    string `json:"-"`
	SubscriptionId string `json:"-"`
	EventId        string `json:"-"`
}

type GetWebhookDeliveryAttemptsResponse struct {
	EventId        string                          `json:"event_id"`
	EventType      string                          `json:"event_type"`
	SubscriptionId string                          `json:"subscription_id"`
	DeliveryStatus model.WebhookDeliveryStatus     `json:"delivery_status"`
	Attempts       []*model.WebhookDeliveryAttempt `json:"attempts"`

	// EventPayload is the raw JSON payload that was delivered to the subscriber. The API sends it
	// base64-encoded; encoding/json decodes it into raw bytes.
	EventPayload []byte                             `json:"event_payload,omitempty"`
	Request      *GetWebhookDeliveryAttemptsRequest `json:"-"`
}

// GetWebhookDeliveryAttempts lists the delivery attempts of a single webhook event.
func (s *webhooksServiceImpl) GetWebhookDeliveryAttempts(
	ctx context.Context,
	request *GetWebhookDeliveryAttemptsRequest,
) (*GetWebhookDeliveryAttemptsResponse, error) {

	path := fmt.Sprintf(
		"/portfolios/%s/webhooks/subscriptions/%s/deliveries/%s/attempts",
		request.PortfolioId,
		request.SubscriptionId,
		request.EventId,
	)

	response := &GetWebhookDeliveryAttemptsResponse{Request: request}

	if err := client.HttpGet(
		ctx,
		s.client,
		path,
		core.EmptyQueryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		response,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return response, nil
}
