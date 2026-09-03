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
	"github.com/coinbase/prime-sdk-go/utils"
)

type ListCrossMarginLiquidationsRequest struct {
	EntityId   string                    `json:"-"`
	Status     model.XMLiquidationStatus `json:"status,omitempty"`
	StartTime  string                    `json:"start_time,omitempty"`
	EndTime    string                    `json:"end_time,omitempty"`
	Pagination *model.PaginationParams   `json:"pagination_params,omitempty"`
}

type ListCrossMarginLiquidationsResponse struct {
	model.PaginationMixin
	Liquidations  []*model.XMLiquidationSummary       `json:"liquidations"`
	Request       *ListCrossMarginLiquidationsRequest `json:"-"`
	service       FinancingService
	serviceConfig *model.ServiceConfig
}

func (r *ListCrossMarginLiquidationsResponse) Next(ctx context.Context) (*ListCrossMarginLiquidationsResponse, error) {
	if !r.HasNext() {
		return nil, nil
	}

	nextRequest := *r.Request
	nextRequest.Pagination = model.PrepareNextPagination(r.Request.Pagination, r.GetNextCursor())

	return r.service.ListCrossMarginLiquidations(ctx, &nextRequest)
}

func (r *ListCrossMarginLiquidationsResponse) Iterator() *model.PageIterator[*ListCrossMarginLiquidationsResponse, *model.XMLiquidationSummary] {
	return model.NewPageIteratorWithConfig(
		r,
		func(resp *ListCrossMarginLiquidationsResponse) []*model.XMLiquidationSummary {
			return resp.Liquidations
		},
		r.serviceConfig,
	)
}

func (s *financingServiceImpl) ListCrossMarginLiquidations(
	ctx context.Context,
	request *ListCrossMarginLiquidationsRequest,
) (*ListCrossMarginLiquidationsResponse, error) {

	path := fmt.Sprintf("/entities/%s/cross_margin/liquidations", request.EntityId)

	request.Pagination = utils.ApplyDefaultLimit(request.Pagination, s.serviceConfig)

	queryParams := core.EmptyQueryParams
	if request.Status != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "status", string(request.Status))
	}
	if request.StartTime != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "start_time", request.StartTime)
	}
	if request.EndTime != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "end_time", request.EndTime)
	}
	queryParams = utils.AppendPaginationParams(queryParams, request.Pagination)

	response := &ListCrossMarginLiquidationsResponse{
		Request:       request,
		service:       s,
		serviceConfig: s.serviceConfig,
	}

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
