# Webhook examples

Run these from the repository root. Every example reads `PRIME_CREDENTIALS` (the JSON object documented in the [SDK README](../../README.md)).

```bash
export PRIME_CREDENTIALS='{"accessKey":"","passphrase":"","signingKey":"","portfolioId":"","svcAccountId":"","entityId":""}'
```

`CreateWebhookSubscription` and `RotateWebhookHmacSecret` return the HMAC secret only once. The examples redact it and print whether one was returned. Store the real value in a secret manager.

## List event types

```bash
go run ./examples/webhooks/getWebhookEventTypes
```

## Create a subscription

The first argument is a comma-separated list of names from the previous command.

```bash
go run ./examples/webhooks/createWebhookSubscription prime.transaction.created https://example.com/webhooks
go run ./examples/webhooks/createWebhookSubscription prime.transaction.created,prime.transaction.updated https://example.com/webhooks
```

## List subscriptions

```bash
go run ./examples/webhooks/listWebhookSubscriptions
```

## Get a subscription

```bash
go run ./examples/webhooks/getWebhookSubscription <subscription_id>
```

## Update the target URL

Reads the current subscription first and sends its version, so the update is rejected if the subscription changed in the meantime.

```bash
go run ./examples/webhooks/updateWebhookSubscription <subscription_id> https://example.com/webhooks
```

## Send a simulated event

```bash
go run ./examples/webhooks/sendWebhookSimulation <subscription_id> prime.transaction.created
```

## Delivery logs

```bash
go run ./examples/webhooks/getWebhookDeliveryLogs <subscription_id>
```

## Delivery attempts for one event

```bash
go run ./examples/webhooks/getWebhookDeliveryAttempts <subscription_id> <event_id>
```

## Subscription stats

Uses a 7 day period.

```bash
go run ./examples/webhooks/getWebhookSubscriptionStats <subscription_id>
```

## Portfolio stats

Uses a 7 day period.

```bash
go run ./examples/webhooks/getWebhookOwnerStats
```

## Rotate the HMAC secret

Pass `true` to invalidate the previous secret immediately. Omit it to keep the previous secret valid for the server grace window.

```bash
go run ./examples/webhooks/rotateWebhookHmacSecret <subscription_id>
go run ./examples/webhooks/rotateWebhookHmacSecret <subscription_id> true
```

## Replay an event

The subscription id is optional. Omit it to replay the event to every matching subscription.

```bash
go run ./examples/webhooks/replayWebhookEvent <event_id>
go run ./examples/webhooks/replayWebhookEvent <event_id> <subscription_id>
```

## Delete a subscription

```bash
go run ./examples/webhooks/deleteWebhookSubscription <subscription_id>
```
