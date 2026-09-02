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

package financing

import (
	"context"
	"fmt"

	"github.com/coinbase/core-go"
	"github.com/coinbase/prime-sdk-go/client"
	"github.com/coinbase/prime-sdk-go/model"
)

type GetPortfolioRewardsRateRequest struct {
	PortfolioId string `json:"-"`
}

type GetPortfolioRewardsRateResponse struct {
	CurrentRate    string                          `json:"current_rate"`
	AvailableRates []*model.RewardsRateTier        `json:"available_rates"`
	Request        *GetPortfolioRewardsRateRequest `json:"-"`
}

func (s *financingServiceImpl) GetPortfolioRewardsRate(
	ctx context.Context,
	request *GetPortfolioRewardsRateRequest,
) (*GetPortfolioRewardsRateResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/rewards/rate", request.PortfolioId)

	response := &GetPortfolioRewardsRateResponse{Request: request}

	if err := core.HttpGet(
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
