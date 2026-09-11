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

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/coinbase/prime-sdk-go/client"
	"github.com/coinbase/prime-sdk-go/credentials"
	primeerrors "github.com/coinbase/prime-sdk-go/model/errors"
	"github.com/coinbase/prime-sdk-go/orders"
)

func main() {
	creds, err := credentials.ReadEnvCredentials("PRIME_CREDENTIALS")
	if err != nil {
		log.Fatalf("unable to read credentials from environment: %v", err)
	}

	httpClient, err := client.DefaultHttpClient()
	if err != nil {
		log.Fatalf("unable to load default http client: %v", err)
	}

	restClient := client.NewRestClient(creds, httpClient)
	ordersSvc := orders.NewOrdersService(restClient)

	// Default to a syntactically valid UUID that will not match an order so the
	// typed error path is exercised. Pass a real order ID as argv[1] to fetch it.
	orderId := "00000000-0000-0000-0000-000000000000"
	if len(os.Args) > 1 {
		orderId = os.Args[1]
	}

	resp, err := ordersSvc.GetOrder(context.Background(), &orders.GetOrderRequest{
		PortfolioId: creds.PortfolioId,
		OrderId:     orderId,
	})
	if err != nil {
		handlePrimeError(err)
		os.Exit(1)
	}

	if resp.Order == nil {
		log.Fatal("order missing from response")
	}
	fmt.Printf("order id=%s status=%s\n", resp.Order.Id, resp.Order.Status)
}

func handlePrimeError(err error) {
	apiErr, ok := primeerrors.From(err)
	if !ok {
		log.Printf("non-API error: %v", err)
		return
	}

	fmt.Println(apiErr.Format())

	switch {
	case primeerrors.IsSubcode(err, primeerrors.SubcodeOrderNotFound):
		fmt.Println("typed match: order was not found for this portfolio")
	case primeerrors.IsSubcode(err, primeerrors.SubcodeOrderIdInvalid):
		fmt.Println("typed match: order_id is not a valid UUID")
	case primeerrors.IsCode(err, primeerrors.ErrorCodeRateLimitExceeded):
		fmt.Println("typed match: rate limited; backoff before retrying")
	default:
		fmt.Printf("code=%s subcode=%s trace_id=%s\n", apiErr.Code, apiErr.Subcode, apiErr.TraceID)
	}

	if apiErr.Retryable() {
		fmt.Println("this error is typically safe to retry after backoff")
	}
}
