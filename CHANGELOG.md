# Changelog

## [0.11.0] - 2026-SEP-02

### Added

#### New API Endpoints

- **`GetCrossMarginLiquidation`**: Get detailed cross-margin liquidation data for an entity (`GET /entities/{entity_id}/cross_margin/liquidation`)
- **`ListCrossMarginLiquidations`**: List historical cross-margin liquidations for an entity (`GET /entities/{entity_id}/cross_margin/liquidations`)
- **`ListTradeFinanceObligations`**: List trade finance obligations for an entity (`GET /entities/{entity_id}/tf_obligations`)
- **`GetFcmEquity`**: Retrieve FCM equity data for an entity (`GET /entities/{entity_id}/futures/equity`)
- **`GetEntityRewardsRate`**: Get current rewards rate and available tiers for an entity (beta) (`GET /entities/{entity_id}/rewards/rate`)
- **`GetPortfolioRewardsRate`**: Get current rewards rate and available tiers for a portfolio (beta) (`GET /portfolios/{portfolio_id}/rewards/rate`)

#### New & Updated Models

- **`Product`**: Added `option_product_details`
- **`OptionProductDetails`**: Option-specific product fields (strike, expiry, settlement, lot size, price increment steps)
- **`PriceIncrementStep`**: Tiered price increment override for option products
- **`FcmBalance`**: Added `cfm_unsettled_accrued_funding_pnl`
- **`FcmEquity`**: Prior EOD equity, unrealized P&L, excess/deficit, and available-to-sweep amounts
- **`TFObligation`**: Trade finance obligation (loan) for an entity
- **`XMLiquidationDetail`**, **`XMLiquidationSummary`**, **`XMLiquidatedAsset`**: Cross-margin liquidation detail, summary, and per-asset breakdown
- **`RewardsRateTier`**: A single tier in the rewards rate card

#### Error codes and subcodes

- **`model/errors`**: Typed Prime REST `ErrorCode` and `Subcode` constants generated from OpenAPI `x-error-codes` / `x-subcodes`, with spec descriptions on each constant
- **`client.HttpGet` / `HttpPost` / `HttpPut` / `HttpDelete` / `HttpPatch`**: SDK HTTP helpers that parse `{ code, message, subcode, trace_id }` into `*errors.APIError`
- Helpers: `errors.From`, `IsCode`, `IsSubcode`, `(*APIError).Format`, `Retryable`
- **`examples/advanced/errorHandling`**: Sample that inspects `*errors.APIError` with `From`, `IsSubcode`, `IsCode`, `Format`, and `Retryable`

#### New Enums

- **`SettlementPeriod`**
- **`SettlementModel`**
- **`RewardsRateTierType`**

## [0.10.0] - 2026-AUG-28

### Added

#### New API Endpoints

- **`GetConversionFees`**: Get organization stablecoin conversion fee tiers and month-to-date net conversion volume (`GET /conversion/fees`)
- **`GetDerivativesCurrencySummary`**: Retrieve per-currency international derivatives balances for a portfolio (`GET /portfolios/{portfolio_id}/derivatives/currency_summary`)
- **`GetDerivativePositions`**: Retrieve active derivative positions for a portfolio (`GET /portfolios/{portfolio_id}/derivatives/positions`)

#### New & Updated Models

- **`Product`**: `ProductType` now includes `OPTION`
- **`ConversionFee`**: Per-pair conversion fee row with month-to-date volume and progressive tiers
- **`ConversionFeeTier`**: A single tier in the progressive stablecoin conversion schedule
- **`DerivativePosition`**: A derivative position across product types, including options details
- **`DerivativesCurrencyBalance`**: Per-currency international derivatives balances
- **`OptionsDetails`**: Options greeks and strike for a derivative position

#### New Enums

- **`ProductTypeOption`**
- **`DerivativeProductType`**
- **`OptionType`**
- **`FcmPositionSide`**

## [0.9.2] - 2026-AUG-5

### Added

- **`Order`**: Added `is_buy_exact` field

## [0.9.1] - 2026-JUL-24

### Added

#### Updated Request/Response Fields

- **`TravelRuleParty`**: Added `vasp_address` field

## [0.9.0] - 2026-JUN-23

### Changed

- **Breaking:** Renamed `SetFundingSettings` to `UpdateFundingSettings` (GA graduation from beta)
- **Breaking:** Fixed `UpdateFundingSettings` path from `/entities/{id}/funding/settings` to `/entities/{id}/funding_settings`
- `GetCrossMarginPrimeOverview` operationId updated to GA (`PrimeRESTAPI_GetCrossMarginPrimeOverview`)

### Added

#### New & Updated Models

- **`WalletStakingMetadata`**: Optional metadata for wallet stake/unstake requests (`external_id`)
- **`ValidatorProvider`**: ETH validator service provider enum for portfolio unstaking
- **`CustomStablecoinAsset`**, **`CustomStablecoinRewardDetails`**: Custom stablecoin reward payout details
- **`RewardSubtype`**: Added `TRANSACTION_REWARD`, `STAKING_FEE_REBATE_REWARD`, `BUIDL_DIVIDEND`, `CUSTOM_STABLECOIN_REWARD`
- **`RewardMetadata`**: Added `custom_stablecoin_reward_details` field
- **`TravelRuleData`**, **`CounterpartyDestination`**: Travel rule and counterparty destination types for withdrawals

#### Updated Request/Response Fields

- **`EditOrderRequest`**: Added `offset`, `wig_level` (PEG order fields)
- **`CreateStakeRequest`**, **`CreateUnstakeRequest`**: Added optional `metadata` (`WalletStakingMetadata`)
- **`PortfolioUnstakeRequest`**: Added `validator_provider`; `amount` is now optional
- **`CreateWalletWithdrawalRequest`**: Added `counterparty`, `travel_rule_data`
- **`CreateWalletWithdrawalResponse`**: Added `counterparty_destination`

## [0.8.1] - 2026-05-29

### Fixed

- `User-Agent` header now reports the correct SDK version via `sdkVersion` (was incorrectly `0.7.0` after the v0.8.0 module release).

## [0.8.0] - 2026-05-29

### Changed

- Relocated module from `github.com/coinbase-samples/prime-sdk-go` to `github.com/coinbase/prime-sdk-go`.
- Depends on `github.com/coinbase/core-go` v0.3.0 (replacing `github.com/coinbase-samples/core-go`).
- No intentional API changes; equivalent to v0.7.0 on the previous module path.

## [0.7.0] - 2026-MAY-11

### Added

- New Beta Financing endpoints
  - `GetCrossMarginRiskParameters` — retrieves XM 2.0 tier risk parameters and offset credit matrices for an entity
  - `GetCrossMarginPrimeOverview` — returns full live Prime cross-margin information (served from `/v2`)
  - `SetFundingSettings` — sets FCM funding configuration for an entity (creates a PCS proposal)
  - `GetMarketData` — retrieves paginated volatility and ADV market data for an entity
- `client.VersionedBaseUrl` and `client.WithBaseUrl` helpers for per-call API version overrides (used internally by `GetCrossMarginPrimeOverview` for the `/v2` path, without affecting other calls)
- `NewFinancingServiceWithConfig` constructor for pagination control on `GetMarketData`
- Financing models: `XMLiquidationStatus`, `ActiveLiquidationSummary`; `ActiveLiquidation` field on `CrossMarginOverview`
- Financing models: `MarginAddOn`, `XMPosition`, `XMRiskNettingInfo`; `XMMarginLimit`, `SpotEquity`, `FuturesEquity`, `RiskNettingInfo` fields on `XMSummary`
- Beta financing models: `PrimeXMControlStatus`, `PrimeXMMarginLevel`, `PrimeXMHealthStatus`, `PrimeXMMarginRequirementType`, `PrimeXMMarginThresholdType`, `CrossMarginRiskParameters`, `TierPairRateEntry`, `CrossMarginPrimeMarginSummary`, `CrossMarginPrimeSpotEquityBreakdown`, `CrossMarginPrimeDerivativesEquityBreakdown`, `CrossMarginPrimeRiskNettingInfo`, `PrimeXMMarginRequirementBreakdown`, `PrimeXMOffsetCreditBreakdown`, `CrossMarginPrimeXMPosition`, `PrimeXMMarginCallThresholds`, `PrimeXMMarginThreshold`, `MarketData`
- Staking model: `ValidatorUnstakePreview`; `WalletId`, `WalletAddress`, `CurrentTimestamp`, `Validators` fields on `PreviewUnstakeResponse`
- User model: `BUSINESS_MANAGER` user role
- RFQ: `QuoteDurationMs` optional field on `CreateQuoteRequest` and `CreateQuoteResponse`


## [0.6.3] - 2026-APR-30

### Added
- Add RFQ information to products


## [0.6.2] - 2026-APR-21

### Added
- Add entity name to portfolio struct


## [0.6.1] - 2026-APR-20

### Added
- New attributes on List Assets

## [0.6.0] - 2026-MAR-30

### Added

- New `advancedtransfers` package with four endpoints
  - ListAdvancedTransfers
  - CreateAdvancedTransfer
  - CancelAdvancedTransfer
  - ListAdvancedTransferTransactions
- New Transaction endpoint: GetTransactionTravelRuleData
- New examples: listAdvancedTransfers, createAdvancedTransfer, cancelAdvancedTransfer, getTransactionTravelRuleData
- New models: `AdvancedTransfer`, `AdvancedTransferState`, `AdvancedTransferType`, `BlindMatchMetadata`, `FundMovement`, `TransferLocation`
- New product types: `ProductType`, `ContractExpiryType`, `ExpiringContractStatus`, `FutureProductDetails`, `PerpetualProductDetails`, `FcmTradingSessionDetails`
- New model types: `CommissionDetailTotal`, `UserRole`, `SecondaryPermission`, `FcmMarginHealthState`, `StakingRewardType`, `ValidatorAllocation`

### Updated

- ListProducts supports three new optional query params: `product_type`, `contract_expiry_type`, `expiring_contract_status`
- `Product` model has new fields: `product_type`, `fcm_trading_session_details`, `future_product_details`
- `Order` and `OrderFill` models have new fields: `product_type`, `commission_detail_total`
- `GetFcmRiskLimitsResponse` has new fields: `cfm_unsettled_accrued_funding_pnl`, `margin_utilization_percent`, `margin_health_state`
- `User` model has new fields: `roles`, `secondary_permissions`
- `CreateUnstakeInputs` has new field: `validator_allocations` (Alpha — ETH V2 validator-level unstaking)
- `StakingRewardType` has new enum value: `BUIDL_DIVIDEND`

## [0.5.2] - 2025-JUN-17

### Added

- Wallet Service has two new endpoints
  - listWalletAddresses
  - createWalletAddress

## [0.5.0] - 2025-JUN-11

### Fix

- Align financing models with rest of the SDK

## [0.4.3] - 2025-JUN-02

### Added

- Add Network info to Get and List Wallets
- Added TransactionMetadata to activities


## [0.4.1] - 2025-MAY-13

### Added

- Add disable dynamic nonce to EVM params for Onchain Txs
- Add settle currency to quote requests

## [0.4.0] - 2025-MAY-13

### Fix

- Request structs are now excluded from JSON serialization when marshaling a response

### Added

- Add pagination support for ListOpenOrders and slice support for product IDs

## [0.3.8] - 2025-MAY-07

### Added

- Adding support for new Financing endpoints
  - ListExistingLocations
  - ListInterestAccruals
  - ListPortfolioInterestAccruals
  - ListMarginCallSummaries
  - ListMarginConversions
  - GetEntityLocateAvailabilities
  - GetMarginInformation
  - GetPortfolioBuyingPower
  - GetPortfolioCreditInformation
  - GetPortfolioWithdrawalPower
  - GetTieredPricingFees
  - CreateNewLocates
- Adding support for new Positions endpoints
  - ListAggregateEntityPositions
  - ListEntityPositions
- Adding support for new Balance endpoint
  - ListEntityBalances
- Added support for new staking endpoints
  - CreateStake
  - CreateUnstake

## [0.3.7] - 2025-MAY-01

### Fix

- Fix missing allocations on list request

## [0.3.6] - 2025-APR-30

### Added

- Added NetworkFamily to CreateWallet endpoints
- Now supports evm and solana network families

## [0.3.5] - 2025-APR-08

### Added

- Added OnchainEvmParams option to transation
