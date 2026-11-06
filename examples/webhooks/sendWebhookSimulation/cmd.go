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
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/coinbase/prime-sdk-go/client"
	"github.com/coinbase/prime-sdk-go/credentials"
	"github.com/coinbase/prime-sdk-go/webhooks"
)

func main() {

	credentials, err := credentials.ReadEnvCredentials("PRIME_CREDENTIALS")
	if err != nil {
		log.Fatalf("unable to read credentials from environment: %v", err)
	}

	httpClient, err := client.DefaultHttpClient()
	if err != nil {
		log.Fatalf("unable to load default http client: %v", err)
	}

	restClient := client.NewRestClient(credentials, httpClient)

	svc := webhooks.NewWebhooksService(restClient)

	if len(os.Args) < 3 {
		log.Fatalf("usage: go run ./examples/webhooks/sendWebhookSimulation <subscription_id> <event_type>")
	}
	subscriptionId := os.Args[1]
	eventType := os.Args[2]

	request := &webhooks.SendWebhookSimulationRequest{
		PortfolioId:    credentials.PortfolioId,
		SubscriptionId: subscriptionId,
		EventType:      eventType,
	}

	response, err := svc.SendWebhookSimulation(context.Background(), request)
	if err != nil {
		log.Fatalf("unable to send webhook simulation: %v", err)
	}

	output, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatalf("error marshaling response to JSON: %v", err)
	}
	fmt.Println(string(output))
}
