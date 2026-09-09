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

package client

import (
	"context"

	"github.com/coinbase/core-go"
	apierrors "github.com/coinbase/prime-sdk-go/model/errors"
)

// HttpPost sends a JSON POST request. Unexpected status codes are returned as *errors.ApiError.
func HttpPost(
	ctx context.Context,
	cl RestClient,
	path,
	query string,
	expectedHttpStatusCodes []int,
	request,
	response interface{},
	headersFunc core.HttpHeaderFunc,
) error {
	return core.HttpPost(ctx, cl, path, query, expectedHttpStatusCodes, request, response, headersFunc, parsePrimeApiError)
}

// HttpGet sends a JSON GET request. Unexpected status codes are returned as *errors.ApiError.
func HttpGet(
	ctx context.Context,
	cl RestClient,
	path,
	query string,
	expectedHttpStatusCodes []int,
	request,
	response interface{},
	headersFunc core.HttpHeaderFunc,
) error {
	return core.HttpGet(ctx, cl, path, query, expectedHttpStatusCodes, request, response, headersFunc, parsePrimeApiError)
}

// HttpPut sends a JSON PUT request. Unexpected status codes are returned as *errors.ApiError.
func HttpPut(
	ctx context.Context,
	cl RestClient,
	path,
	query string,
	expectedHttpStatusCodes []int,
	request,
	response interface{},
	headersFunc core.HttpHeaderFunc,
) error {
	return core.HttpPut(ctx, cl, path, query, expectedHttpStatusCodes, request, response, headersFunc, parsePrimeApiError)
}

// HttpDelete sends a JSON DELETE request. Unexpected status codes are returned as *errors.ApiError.
func HttpDelete(
	ctx context.Context,
	cl RestClient,
	path,
	query string,
	expectedHttpStatusCodes []int,
	request,
	response interface{},
	headersFunc core.HttpHeaderFunc,
) error {
	return core.HttpDelete(ctx, cl, path, query, expectedHttpStatusCodes, request, response, headersFunc, parsePrimeApiError)
}

// HttpPatch sends a JSON PATCH request. Unexpected status codes are returned as *errors.ApiError.
func HttpPatch(
	ctx context.Context,
	cl RestClient,
	path,
	query string,
	expectedHttpStatusCodes []int,
	request,
	response interface{},
	headersFunc core.HttpHeaderFunc,
) error {
	return core.HttpPatch(ctx, cl, path, query, expectedHttpStatusCodes, request, response, headersFunc, parsePrimeApiError)
}

func parsePrimeApiError(body []byte, statusCode int, _ []int, callUrl string) error {
	resp, err := apierrors.ParseBody(body)
	if err != nil {
		resp = apierrors.Response{Message: string(body)}
	}
	return &apierrors.ApiError{
		Response:   resp,
		StatusCode: statusCode,
		URL:        callUrl,
	}
}
