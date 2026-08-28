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
	"github.com/coinbase/prime-sdk-go/model"
)

type GetDerivativesCurrencySummaryRequest struct {
	PortfolioId string `json:"-"`
}

type GetDerivativesCurrencySummaryResponse struct {
	PortfolioId string                                `json:"portfolio_id"`
	Balances    []*model.DerivativesCurrencyBalance   `json:"balances"`
	Request     *GetDerivativesCurrencySummaryRequest `json:"-"`
}

func (s *futuresServiceImpl) GetDerivativesCurrencySummary(
	ctx context.Context,
	request *GetDerivativesCurrencySummaryRequest,
) (*GetDerivativesCurrencySummaryResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/derivatives/currency_summary", request.PortfolioId)

	response := &GetDerivativesCurrencySummaryResponse{Request: request}

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
