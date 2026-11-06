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
	"github.com/coinbase/prime-sdk-go/utils"
)

type CreateWebhookSubscriptionRequest struct {
	PortfolioId   string                         `json:"-"`
	EventTypes    []string                       `json:"event_types,omitempty"`
	TargetUrl     string                         `json:"target_url,omitempty"`
	AuthType      model.WebhookAuthType          `json:"auth_type,omitempty"`
	OauthConfig   *model.WebhookOAuthConfigInput `json:"oauth_config,omitempty"`
	CustomHeaders map[string]string              `json:"custom_headers,omitempty"`
	Description   string                         `json:"description,omitempty"`

	// IdempotencyKey is generated automatically when empty.
	IdempotencyKey string `json:"idempotency_key,omitempty"`
}

type CreateWebhookSubscriptionResponse struct {
	Subscription *model.WebhookSubscription `json:"subscription"`

	// HmacSecret is only returned when the subscription is created. Store it securely and never log it.
	HmacSecret string                            `json:"hmac_secret"`
	Request    *CreateWebhookSubscriptionRequest `json:"-"`
}

// CreateWebhookSubscription creates a webhook subscription for a portfolio.
func (s *webhooksServiceImpl) CreateWebhookSubscription(
	ctx context.Context,
	request *CreateWebhookSubscriptionRequest,
) (*CreateWebhookSubscriptionResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/webhooks/subscriptions", request.PortfolioId)

	if len(request.IdempotencyKey) == 0 {
		request.IdempotencyKey = utils.NewUuid()
	}

	response := &CreateWebhookSubscriptionResponse{Request: request}

	if err := client.HttpPost(
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
