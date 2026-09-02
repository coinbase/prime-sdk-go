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

type GetXMLiquidationRequest struct {
	EntityId      string `json:"-"`
	LiquidationId string `json:"liquidation_id,omitempty"`
}

type GetXMLiquidationResponse struct {
	Liquidation *model.XMLiquidationDetail `json:"liquidation"`
	Request     *GetXMLiquidationRequest   `json:"-"`
}

func (s *financingServiceImpl) GetXMLiquidation(
	ctx context.Context,
	request *GetXMLiquidationRequest,
) (*GetXMLiquidationResponse, error) {

	path := fmt.Sprintf("/entities/%s/cross_margin/liquidation", request.EntityId)

	queryParams := core.EmptyQueryParams
	if request.LiquidationId != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "liquidation_id", request.LiquidationId)
	}

	response := &GetXMLiquidationResponse{Request: request}

	if err := core.HttpGet(
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
