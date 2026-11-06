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

	"github.com/coinbase/core-go"
	"github.com/coinbase/prime-sdk-go/client"
	"github.com/coinbase/prime-sdk-go/model"
)

type GetWebhookEventTypesRequest struct{}

type GetWebhookEventTypesResponse struct {
	EventTypes []*model.WebhookEventType    `json:"event_types"`
	Request    *GetWebhookEventTypesRequest `json:"-"`
}

// GetWebhookEventTypes lists the event types that can be subscribed to.
func (s *webhooksServiceImpl) GetWebhookEventTypes(
	ctx context.Context,
	request *GetWebhookEventTypesRequest,
) (*GetWebhookEventTypesResponse, error) {

	path := "/webhooks/event-types"

	response := &GetWebhookEventTypesResponse{Request: request}

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
