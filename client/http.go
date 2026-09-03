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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/coinbase/core-go"
	apierrors "github.com/coinbase/prime-sdk-go/model/errors"
)

type apiRequest struct {
	Path                    string
	Query                   string
	HttpMethod              string
	Body                    []byte
	ExpectedHttpStatusCodes []int
	Client                  RestClient
}

// HttpPost sends a JSON POST request. Unexpected status codes are returned as *errors.APIError.
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
	return call(ctx, cl, path, query, http.MethodPost, expectedHttpStatusCodes, request, response, headersFunc)
}

// HttpGet sends a JSON GET request. Unexpected status codes are returned as *errors.APIError.
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
	return call(ctx, cl, path, query, http.MethodGet, expectedHttpStatusCodes, request, response, headersFunc)
}

// HttpPut sends a JSON PUT request. Unexpected status codes are returned as *errors.APIError.
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
	return call(ctx, cl, path, query, http.MethodPut, expectedHttpStatusCodes, request, response, headersFunc)
}

// HttpDelete sends a JSON DELETE request. Unexpected status codes are returned as *errors.APIError.
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
	return call(ctx, cl, path, query, http.MethodDelete, expectedHttpStatusCodes, request, response, headersFunc)
}

// HttpPatch sends a JSON PATCH request. Unexpected status codes are returned as *errors.APIError.
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
	return call(ctx, cl, path, query, http.MethodPatch, expectedHttpStatusCodes, request, response, headersFunc)
}

func call(
	ctx context.Context,
	cl RestClient,
	path,
	query,
	httpMethod string,
	expectedHttpStatusCodes []int,
	request,
	response interface{},
	headersFunc core.HttpHeaderFunc,
) error {
	body, err := json.Marshal(request)
	if err != nil {
		return err
	}

	resp, err := makeCall(ctx, &apiRequest{
		Path:                    path,
		Query:                   query,
		HttpMethod:              httpMethod,
		Body:                    body,
		ExpectedHttpStatusCodes: expectedHttpStatusCodes,
		Client:                  cl,
	}, headersFunc)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(resp, response); err != nil {
		return err
	}

	return nil
}

func makeCall(ctx context.Context, request *apiRequest, headersFunc core.HttpHeaderFunc) ([]byte, error) {
	callUrl := fmt.Sprintf("%s%s%s", request.Client.HttpBaseUrl(), request.Path, request.Query)

	parsedUrl, err := url.Parse(callUrl)
	if err != nil {
		return nil, &core.ApiError{
			Message:      fmt.Sprintf("invalid URL: %s - %v", callUrl, err),
			ParsedUrl:    callUrl,
			CodeReceived: 0,
		}
	}

	var requestBody []byte
	if request.HttpMethod == http.MethodPost || request.HttpMethod == http.MethodPut || request.HttpMethod == http.MethodPatch {
		requestBody = request.Body
	}

	req, err := http.NewRequestWithContext(ctx, request.HttpMethod, callUrl, bytes.NewReader(requestBody))
	if err != nil {
		return nil, &core.ApiError{
			Message:      err.Error(),
			CodeReceived: 0,
		}
	}

	if headersFunc != nil {
		headersFunc(req, parsedUrl.Path, requestBody, request.Client, time.Now())
	}

	res, err := request.Client.HttpClient().Do(req)
	if err != nil {
		return nil, &core.ApiError{
			Message:      err.Error(),
			CodeReceived: 0,
		}
	}

	defer func() { _ = res.Body.Close() }()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, &core.ApiError{
			Message:      err.Error(),
			CodeReceived: 0,
		}
	}

	for _, code := range request.ExpectedHttpStatusCodes {
		if res.StatusCode == code {
			return body, nil
		}
	}

	return nil, parseAPIError(body, res.StatusCode, callUrl)
}

func parseAPIError(body []byte, statusCode int, callUrl string) error {
	resp, err := apierrors.ParseBody(body)
	if err != nil {
		resp = apierrors.Response{Message: string(body)}
	}
	return &apierrors.APIError{
		Response:   resp,
		StatusCode: statusCode,
		URL:        callUrl,
	}
}
