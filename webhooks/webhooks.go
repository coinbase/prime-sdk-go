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

	"github.com/coinbase/prime-sdk-go/client"
	"github.com/coinbase/prime-sdk-go/model"
)

// WebhooksService manages webhook subscriptions, delivery logs, and statistics (beta).
type WebhooksService interface {
	GetWebhookEventTypes(ctx context.Context, request *GetWebhookEventTypesRequest) (*GetWebhookEventTypesResponse, error)
	ListWebhookSubscriptions(ctx context.Context, request *ListWebhookSubscriptionsRequest) (*ListWebhookSubscriptionsResponse, error)
	CreateWebhookSubscription(ctx context.Context, request *CreateWebhookSubscriptionRequest) (*CreateWebhookSubscriptionResponse, error)
	GetWebhookSubscription(ctx context.Context, request *GetWebhookSubscriptionRequest) (*GetWebhookSubscriptionResponse, error)
	UpdateWebhookSubscription(ctx context.Context, request *UpdateWebhookSubscriptionRequest) (*UpdateWebhookSubscriptionResponse, error)
	DeleteWebhookSubscription(ctx context.Context, request *DeleteWebhookSubscriptionRequest) (*DeleteWebhookSubscriptionResponse, error)
	RotateWebhookHmacSecret(ctx context.Context, request *RotateWebhookHmacSecretRequest) (*RotateWebhookHmacSecretResponse, error)
	SendWebhookSimulation(ctx context.Context, request *SendWebhookSimulationRequest) (*SendWebhookSimulationResponse, error)
	GetWebhookDeliveryLogs(ctx context.Context, request *GetWebhookDeliveryLogsRequest) (*GetWebhookDeliveryLogsResponse, error)
	GetWebhookDeliveryAttempts(ctx context.Context, request *GetWebhookDeliveryAttemptsRequest) (*GetWebhookDeliveryAttemptsResponse, error)
	GetWebhookSubscriptionStats(ctx context.Context, request *GetWebhookSubscriptionStatsRequest) (*GetWebhookSubscriptionStatsResponse, error)
	GetWebhookOwnerStats(ctx context.Context, request *GetWebhookOwnerStatsRequest) (*GetWebhookOwnerStatsResponse, error)
	ReplayWebhookEvent(ctx context.Context, request *ReplayWebhookEventRequest) (*ReplayWebhookEventResponse, error)
	ServiceConfig() *model.ServiceConfig
}

// NewWebhooksService creates a new WebhooksService with default config.
func NewWebhooksService(c client.RestClient) WebhooksService {
	return &webhooksServiceImpl{
		client:        c,
		serviceConfig: model.DefaultServiceConfig(),
	}
}

// NewWebhooksServiceWithConfig creates a new WebhooksService with custom config.
func NewWebhooksServiceWithConfig(c client.RestClient, config *model.ServiceConfig) WebhooksService {
	if config == nil {
		config = model.DefaultServiceConfig()
	}
	return &webhooksServiceImpl{
		client:        c,
		serviceConfig: config,
	}
}

type webhooksServiceImpl struct {
	client        client.RestClient
	serviceConfig *model.ServiceConfig
}

func (s *webhooksServiceImpl) ServiceConfig() *model.ServiceConfig {
	return s.serviceConfig
}
