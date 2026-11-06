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

type UpdateWebhookSubscriptionRequest struct {
	PortfolioId    string `json:"-"`
	SubscriptionId string `json:"-"`

	EventTypes    []string                       `json:"event_types,omitempty"`
	TargetUrl     string                         `json:"target_url,omitempty"`
	AuthType      model.WebhookAuthType          `json:"auth_type,omitempty"`
	OauthConfig   *model.WebhookOAuthConfigInput `json:"oauth_config,omitempty"`
	CustomHeaders map[string]string              `json:"custom_headers,omitempty"`
	Description   string                         `json:"description,omitempty"`

	// IsEnabled is a pointer so that false can be sent explicitly.
	IsEnabled *bool `json:"is_enabled,omitempty"`

	// UpdateMask lists the fields to update (field-mask paths).
	UpdateMask string `json:"update_mask,omitempty"`

	// ExpectedVersion is the subscription Version the update is based on (optimistic concurrency).
	ExpectedVersion string `json:"expected_version,omitempty"`
}

type UpdateWebhookSubscriptionResponse struct {
	Subscription *model.WebhookSubscription        `json:"subscription"`
	Request      *UpdateWebhookSubscriptionRequest `json:"-"`
}

// UpdateWebhookSubscription updates a webhook subscription.
func (s *webhooksServiceImpl) UpdateWebhookSubscription(
	ctx context.Context,
	request *UpdateWebhookSubscriptionRequest,
) (*UpdateWebhookSubscriptionResponse, error) {

	path := fmt.Sprintf(
		"/portfolios/%s/webhooks/subscriptions/%s",
		request.PortfolioId,
		request.SubscriptionId,
	)

	response := &UpdateWebhookSubscriptionResponse{Request: request}

	if err := client.HttpPatch(
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
