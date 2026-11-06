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

type RotateWebhookHmacSecretRequest struct {
	PortfolioId    string `json:"-"`
	SubscriptionId string `json:"-"`

	// RevokePreviousImmediately invalidates the old secret immediately instead of keeping it valid
	// for the server grace window. Use when the secret may be compromised.
	RevokePreviousImmediately bool `json:"revoke_previous_immediately,omitempty"`
}

type RotateWebhookHmacSecretResponse struct {
	Subscription *model.WebhookSubscription `json:"subscription"`

	// HmacSecret is only returned when the secret is rotated. Store it securely and never log it.
	HmacSecret string `json:"hmac_secret"`

	// PreviousSecretExpiresAt is when the previous secret stops being accepted.
	PreviousSecretExpiresAt string                          `json:"previous_secret_expires_at,omitempty"`
	Request                 *RotateWebhookHmacSecretRequest `json:"-"`
}

// RotateWebhookHmacSecret rotates the HMAC signing secret of a webhook subscription.
func (s *webhooksServiceImpl) RotateWebhookHmacSecret(
	ctx context.Context,
	request *RotateWebhookHmacSecretRequest,
) (*RotateWebhookHmacSecretResponse, error) {

	path := fmt.Sprintf(
		"/portfolios/%s/webhooks/subscriptions/%s/rotate-hmac-secret",
		request.PortfolioId,
		request.SubscriptionId,
	)

	response := &RotateWebhookHmacSecretResponse{Request: request}

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
