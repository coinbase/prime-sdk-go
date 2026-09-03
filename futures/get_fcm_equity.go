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

package futures

import (
	"context"
	"fmt"

	"github.com/coinbase/core-go"
	"github.com/coinbase/prime-sdk-go/client"
)

type GetFcmEquityRequest struct {
	EntityId string `json:"-"`
}

type GetFcmEquityResponse struct {
	EodAccountEquity     string               `json:"eod_account_equity"`
	EodUnrealizedPnl     string               `json:"eod_unrealized_pnl"`
	CurrentExcessDeficit string               `json:"current_excess_deficit"`
	AvailableToSweep     string               `json:"available_to_sweep"`
	Request              *GetFcmEquityRequest `json:"-"`
}

func (s *futuresServiceImpl) GetFcmEquity(
	ctx context.Context,
	request *GetFcmEquityRequest,
) (*GetFcmEquityResponse, error) {

	path := fmt.Sprintf("/entities/%s/futures/equity", request.EntityId)

	response := &GetFcmEquityResponse{Request: request}

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
