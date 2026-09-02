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

package model

// FcmMarginHealthState represents the margin health state of an FCM account.
type FcmMarginHealthState string

const (
	FcmMarginHealthStateUnspecified    FcmMarginHealthState = "FCM_MARGIN_HEALTH_STATE_UNSPECIFIED"
	FcmMarginHealthStateHealthy        FcmMarginHealthState = "FCM_MARGIN_HEALTH_STATE_HEALTHY"
	FcmMarginHealthStateRestricted     FcmMarginHealthState = "FCM_MARGIN_HEALTH_STATE_RESTRICTED"
	FcmMarginHealthStatePreLiquidation FcmMarginHealthState = "FCM_MARGIN_HEALTH_STATE_PRE_LIQUIDATION"
	FcmMarginHealthStateLiquidation    FcmMarginHealthState = "FCM_MARGIN_HEALTH_STATE_LIQUIDATION"
)

// FcmMarginCallType represents the type of margin call
type FcmMarginCallType string

const (
	FcmMarginCallTypeUnspecified FcmMarginCallType = "FCM_MARGIN_CALL_TYPE_UNSPECIFIED"
	FcmMarginCallTypeUrgent      FcmMarginCallType = "FCM_MARGIN_CALL_TYPE_URGENT"
	FcmMarginCallTypeRegular     FcmMarginCallType = "FCM_MARGIN_CALL_TYPE_REGULAR"
)

// FcmMarginCallState represents the state of a margin call
type FcmMarginCallState string

const (
	FcmMarginCallStateUnspecified FcmMarginCallState = "FCM_MARGIN_CALL_STATE_UNSPECIFIED"
	FcmMarginCallStateClosed      FcmMarginCallState = "FCM_MARGIN_CALL_STATE_CLOSED"
	FcmMarginCallStateRolledOver  FcmMarginCallState = "FCM_MARGIN_CALL_STATE_ROLLED_OVER"
	FcmMarginCallStateDefault     FcmMarginCallState = "FCM_MARGIN_CALL_STATE_DEFAULT"
	FcmMarginCallStateOfficial    FcmMarginCallState = "FCM_MARGIN_CALL_STATE_OFFICIAL"
)

// FcmMarginCall represents an FCM margin call
type FcmMarginCall struct {
	Type            FcmMarginCallType  `json:"type"`
	State           FcmMarginCallState `json:"state"`
	InitialAmount   string             `json:"initial_amount"`
	RemainingAmount string             `json:"remaining_amount"`
	BusinessDate    string             `json:"business_date"`
	CureDeadline    string             `json:"cure_deadline"`
}

// FcmRiskLimits represents FCM risk limits for an entity
type FcmRiskLimits struct {
	CfmRiskLimit                  string `json:"cfm_risk_limit"`
	CfmRiskLimitUtilization       string `json:"cfm_risk_limit_utilization"`
	CfmTotalMargin                string `json:"cfm_total_margin"`
	CfmDeltaOte                   string `json:"cfm_delta_ote"`
	CfmUnsettledRealizedPnl       string `json:"cfm_unsettled_realized_pnl"`
	CfmUnsettledAccruedFundingPnl string `json:"cfm_unsettled_accrued_funding_pnl"`
}

// FcmSettings represents FCM settings for an entity
type FcmSettings struct {
	TargetDerivativesExcess string `json:"target_derivatives_excess"`
}

// FcmBalance represents FCM balance information for a portfolio
type FcmBalance struct {
	PortfolioId                   string `json:"portfolio_id"`
	CfmUsdBalance                 string `json:"cfm_usd_balance"`
	UnrealizedPnl                 string `json:"unrealized_pnl"`
	DailyRealizedPnl              string `json:"daily_realized_pnl"`
	ExcessLiquidity               string `json:"excess_liquidity"`
	FuturesBuyingPower            string `json:"futures_buying_power"`
	InitialMargin                 string `json:"initial_margin"`
	MaintenanceMargin             string `json:"maintenance_margin"`
	ClearingAccountId             string `json:"clearing_account_id"`
	CfmUnsettledAccruedFundingPnl string `json:"cfm_unsettled_accrued_funding_pnl,omitempty"`
}

// FcmEquity represents FCM equity data for an entity.
type FcmEquity struct {
	EodAccountEquity     string `json:"eod_account_equity,omitempty"`
	EodUnrealizedPnl     string `json:"eod_unrealized_pnl,omitempty"`
	CurrentExcessDeficit string `json:"current_excess_deficit,omitempty"`
	AvailableToSweep     string `json:"available_to_sweep,omitempty"`
}

// FcmPosition represents a futures position
type FcmPosition struct {
	ProductId         string `json:"product_id"`
	Side              string `json:"side"`
	NumberOfContracts string `json:"number_of_contracts"`
	DailyRealizedPnl  string `json:"daily_realized_pnl"`
	UnrealizedPnl     string `json:"unrealized_pnl"`
	CurrentPrice      string `json:"current_price"`
	AvgEntryPrice     string `json:"avg_entry_price"`
	ExpirationTime    string `json:"expiration_time"`
}

// FcmPositionSide represents the side of an FCM or derivative position.
type FcmPositionSide string

const (
	FcmPositionSideLong  FcmPositionSide = "LONG"
	FcmPositionSideShort FcmPositionSide = "SHORT"
)

// DerivativeProductType represents the general type of a derivative product.
type DerivativeProductType string

const (
	DerivativeProductTypeUnspecified      DerivativeProductType = "DERIVATIVE_PRODUCT_TYPE_UNSPECIFIED"
	DerivativeProductTypeSpot             DerivativeProductType = "DERIVATIVE_PRODUCT_TYPE_SPOT"
	DerivativeProductTypeFuture           DerivativeProductType = "DERIVATIVE_PRODUCT_TYPE_FUTURE"
	DerivativeProductTypeEquity           DerivativeProductType = "DERIVATIVE_PRODUCT_TYPE_EQUITY"
	DerivativeProductTypePredictionMarket DerivativeProductType = "DERIVATIVE_PRODUCT_TYPE_PREDICTION_MARKET"
	DerivativeProductTypeOption           DerivativeProductType = "DERIVATIVE_PRODUCT_TYPE_OPTION"
	DerivativeProductTypeBasis            DerivativeProductType = "DERIVATIVE_PRODUCT_TYPE_BASIS"
	DerivativeProductTypeEquityOption     DerivativeProductType = "DERIVATIVE_PRODUCT_TYPE_EQUITY_OPTION"
	DerivativeProductTypeFutureCombo      DerivativeProductType = "DERIVATIVE_PRODUCT_TYPE_FUTURE_COMBO"
	DerivativeProductTypeOptionCombo      DerivativeProductType = "DERIVATIVE_PRODUCT_TYPE_OPTION_COMBO"
)

// OptionType represents the type of an option position.
type OptionType string

const (
	OptionTypeCall OptionType = "OPTION_TYPE_CALL"
	OptionTypePut  OptionType = "OPTION_TYPE_PUT"
)

// OptionsDetails contains options-specific details for a derivative position, including greeks.
type OptionsDetails struct {
	Delta      string     `json:"delta,omitempty"`
	Gamma      string     `json:"gamma,omitempty"`
	Theta      string     `json:"theta,omitempty"`
	Vega       string     `json:"vega,omitempty"`
	Strike     string     `json:"strike,omitempty"`
	OptionType OptionType `json:"option_type,omitempty"`
}

// DerivativePosition is a single derivative position across all derivative product types.
type DerivativePosition struct {
	ProductId         string                `json:"product_id,omitempty"`
	Side              FcmPositionSide       `json:"side,omitempty"`
	NumberOfContracts string                `json:"number_of_contracts,omitempty"`
	DailyRealizedPnl  string                `json:"daily_realized_pnl,omitempty"`
	UnrealizedPnl     string                `json:"unrealized_pnl,omitempty"`
	CurrentPrice      string                `json:"current_price,omitempty"`
	AvgEntryPrice     string                `json:"avg_entry_price,omitempty"`
	ExpirationTime    string                `json:"expiration_time,omitempty"`
	ProductType       DerivativeProductType `json:"product_type,omitempty"`
	Currency          string                `json:"currency,omitempty"`
	OptionsDetails    *OptionsDetails       `json:"options_details,omitempty"`
	VenueId           string                `json:"venue_id,omitempty"`
}

// DerivativesCurrencyBalance contains balances for a single settlement currency
// within an international derivatives portfolio.
type DerivativesCurrencyBalance struct {
	Currency          string               `json:"currency,omitempty"`
	Balance           string               `json:"balance,omitempty"`
	UnrealizedPnl     string               `json:"unrealized_pnl,omitempty"`
	RealizedPnl       string               `json:"realized_pnl,omitempty"`
	InitialMargin     string               `json:"initial_margin,omitempty"`
	MaintenanceMargin string               `json:"maintenance_margin,omitempty"`
	MarginBalance     string               `json:"margin_balance,omitempty"`
	OptionValue       string               `json:"option_value,omitempty"`
	MarginExcess      string               `json:"margin_excess,omitempty"`
	MarginUtilization string               `json:"margin_utilization,omitempty"`
	MarginHealthState FcmMarginHealthState `json:"margin_health_state,omitempty"`
}

// FcmSweep represents a futures sweep
type FcmSweep struct {
	Id              string           `json:"id"`
	RequestedAmount *RequestedAmount `json:"requested_amount"`
	ShouldSweepAll  bool             `json:"should_sweep_all"`
	Status          string           `json:"status"`
	ScheduledTime   string           `json:"scheduled_time"`
}

// RequestedAmount represents a requested amount with currency
type RequestedAmount struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}
