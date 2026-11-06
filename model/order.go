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

import "time"

// Order type constants
const (
	OrderTypeMarket = "MARKET"
	OrderTypeLimit  = "LIMIT"
	OrderTypeTwap   = "TWAP"
	OrderTypeBlock  = "BLOCK"
	OrderTypeSpread = "SPREAD"
)

// Time in force constants
const (
	TimeInForceGoodUntilTime      = "GOOD_UNTIL_DATE_TIME"
	TimeInForceGoodUntilCancelled = "GOOD_UNTIL_CANCELLED"
	TimeInForceImmediateOrCancel  = "IMMEDIATE_OR_CANCEL"
)

// OrderSide represents the side of an order (buy or sell)
type OrderSide string

const (
	OrderSideBuy     OrderSide = "BUY"
	OrderSideSell    OrderSide = "SELL"
	OrderSideUnknown OrderSide = "UNKNOWN_ORDER_SIDE"

	// OrderSideMultiple is response-only: a SPREAD order whose legs carry their own sides.
	OrderSideMultiple OrderSide = "MULTIPLE"
)

// Order represents a Prime order
type Order struct {
	PortfolioId string `json:"portfolio_id"`
	Side        string `json:"side"`

	// A client-generated order ID used for reference purposes (note: order will be rejected if this ID
	// is not unique among all currently active orders)
	ClientOrderId string `json:"client_order_id"`
	ProductId     string `json:"product_id"`
	Type          string `json:"type"`

	// Order size in base asset units (either `base_quantity` or `quote_value` is required).
	// For SPREAD orders, required and in the primary leg's units: a whole number of contracts
	// when it is a futures product.
	BaseQuantity string `json:"base_quantity"`

	// Order size in quote asset units, i.e. the amount the user wants to spend (when buying) or receive (when selling);
	// the quantity in base units will be determined based on the market liquidity and indicated `quote_value` (either
	// `base_quantity` or `quote_value` is required). Not supported for SPREAD orders.
	QuoteValue string `json:"quote_value,omitempty"`

	LimitPrice string `json:"limit_price,omitempty"`

	// The start time of the order in UTC (TWAP only)
	StartTime string `json:"start_time,omitempty"`

	// The expiry time of the order in UTC (TWAP and limit GTD only)
	ExpiryTime  string `json:"expiry_time,omitempty"`
	TimeInForce string `json:"time_in_force,omitempty"`

	// An optional self trade prevention id (in the form of a UUID). The value is only honored for certain
	// clients who are permitted to specify a custom self trade prevention id
	StpId string `json:"stp_id,omitempty"`

	// Optionally specify a display size. This is the maximum order size that will show up on venue order books.
	// Specifying a value here effectively makes a LIMIT order into an "iceberg" style order.
	// This property only applies to LIMIT orders and will be ignored for other order types.
	DisplayQuoteSize string `json:"display_quote_size,omitempty"`
	DisplayBaseSize  string `json:"display_base_size,omitempty"`

	// If you pass is_raise_exact = TRUE, you must use quote_value = n where n is the amount you want,
	// so $2000 will then cost you 1 ETH + fee, requiring > 1 ETH
	IsRaiseExact bool `json:"is_raise_exact,omitempty"`

	// Buy Exact order flag. When true, fees for a BUY order sized in quote_value are charged on top of
	// the requested quote_value instead of being carved out of it. Only valid for BUY orders sized in
	// quote_value on SPOT products.
	IsBuyExact bool `json:"is_buy_exact,omitempty"`

	// Used for describe order, create order preview, and list portfolio orders
	Id                    string `json:"id,omitempty"`
	UserId                string `json:"user_id,omitempty"`
	Created               string `json:"created_at,omitempty"`
	FilledQuantity        string `json:"filled_quantity,omitempty"`
	FilledValue           string `json:"filled_value,omitempty"`
	AverageFilledPrice    string `json:"average_filled_price,omitempty"`
	Commission            string `json:"commission,omitempty"`
	ExchangeFee           string `json:"exchange_fee,omitempty"`
	Total                 string `json:"order_total,omitempty"`
	BestBid               string `json:"best_bid,omitempty"`
	BestAsk               string `json:"best_ask,omitempty"`
	Slippage              string `json:"slippage,omitempty"`
	Status                string `json:"status,omitempty"`
	HistoricalPov         string `json:"historical_pov,omitempty"`
	StopPrice             string `json:"stop_price,omitempty"`
	NetAverageFilledPrice string `json:"net_average_filled_price,omitempty"`
	UserContext           string `json:"user_context,omitempty"`
	ClientProductId       string `json:"client_product_id,omitempty"`
	PostOnly              bool   `json:"post_only,omitempty"`
	// Deprecated: Use EditHistory instead
	OrderEditHistory      []*OrderEditHistory    `json:"order_edit_history,omitempty"`
	DisplaySize           string                 `json:"display_size,omitempty"`
	EditHistory           []*EditHistory         `json:"edit_history,omitempty"`
	PegOffsetType         string                 `json:"peg_offset_type,omitempty"`
	Offset                string                 `json:"offset,omitempty"`
	WigLevel              string                 `json:"wig_level,omitempty"`
	ProductType           ProductType            `json:"product_type,omitempty"`
	CommissionDetailTotal *CommissionDetailTotal `json:"commission_detail_total,omitempty"`

	// How limit_price is expressed (SPREAD and option orders)
	LimitPriceType LimitPriceType `json:"limit_price_type,omitempty"`

	// Request only: configuration of the two legs of a SPREAD order
	SpreadConfig *SpreadConfig `json:"spread_config,omitempty"`

	// Response only: details of a SPREAD order, including per-leg fills
	SpreadDetails *SpreadDetails `json:"spread_details,omitempty"`

	// Order preview only
	OptionGreeks  *OptionGreeks  `json:"option_greeks,omitempty"`
	SpreadPreview *SpreadPreview `json:"spread_preview,omitempty"`

	// Futures and SPREAD order previews only: derivatives margin limit usage before and after the order
	CurrentMarginUtilization        string `json:"current_margin_utilization,omitempty"`
	CurrentMarginUtilizationRatio   string `json:"current_margin_utilization_ratio,omitempty"`
	ProjectedMarginUtilization      string `json:"projected_margin_utilization,omitempty"`
	ProjectedMarginUtilizationRatio string `json:"projected_margin_utilization_ratio,omitempty"`

	// The requested settlement currency, echoed on order previews
	SettlCurrency string `json:"settl_currency,omitempty"`
}

// LimitPriceType describes how a limit price is expressed.
type LimitPriceType string

const (
	LimitPriceTypeBaseAsset         LimitPriceType = "LIMIT_PRICE_TYPE_BASE_ASSET"
	LimitPriceTypeQuoteAsset        LimitPriceType = "LIMIT_PRICE_TYPE_QUOTE_ASSET"
	LimitPriceTypeImpliedVolatility LimitPriceType = "LIMIT_PRICE_TYPE_IMPLIED_VOLATILITY"
)

// SpreadLimitType is the unit of a SPREAD order limit.
type SpreadLimitType string

const (
	SpreadLimitTypeBps SpreadLimitType = "SPREAD_LIMIT_TYPE_BPS"
	SpreadLimitTypeUsd SpreadLimitType = "SPREAD_LIMIT_TYPE_USD"
)

// MaxDeltaType is the unit of a SPREAD order max delta size.
type MaxDeltaType string

const (
	MaxDeltaTypeSpot     MaxDeltaType = "MAX_DELTA_TYPE_SPOT"
	MaxDeltaTypePercent  MaxDeltaType = "MAX_DELTA_TYPE_PERCENT"
	MaxDeltaTypeContract MaxDeltaType = "MAX_DELTA_TYPE_CONTRACT"
)

// Leg is a single leg of a SPREAD order request.
type Leg struct {
	ProductId           string    `json:"product_id,omitempty"`
	Side                OrderSide `json:"side,omitempty"`
	SupportPassiveOrder bool      `json:"support_passive_order,omitempty"`
}

// SpreadConfig configures a SPREAD order (request).
type SpreadConfig struct {
	Limit        string          `json:"limit,omitempty"`
	LimitType    SpreadLimitType `json:"limit_type,omitempty"`
	MaxDeltaSize string          `json:"max_delta_size,omitempty"`
	MaxDeltaType MaxDeltaType    `json:"max_delta_type,omitempty"`
	PrimaryLeg   *Leg            `json:"primary_leg,omitempty"`
	SecondaryLeg *Leg            `json:"secondary_leg,omitempty"`
}

// SpreadLegDetails contains per-leg fill details of a SPREAD order (response).
type SpreadLegDetails struct {
	ProductId             string                 `json:"product_id,omitempty"`
	Side                  OrderSide              `json:"side,omitempty"`
	SupportPassiveOrder   bool                   `json:"support_passive_order,omitempty"`
	FilledQuantity        string                 `json:"filled_quantity,omitempty"`
	AverageFilledPrice    string                 `json:"average_filled_price,omitempty"`
	Commission            string                 `json:"commission,omitempty"`
	CommissionDetailTotal *CommissionDetailTotal `json:"commission_detail_total,omitempty"`
	ExchangeFee           string                 `json:"exchange_fee,omitempty"`
	FilledValue           string                 `json:"filled_value,omitempty"`
}

// SpreadDetails contains the details of a SPREAD order (response).
type SpreadDetails struct {
	Limit          string            `json:"limit,omitempty"`
	LimitType      SpreadLimitType   `json:"limit_type,omitempty"`
	MaxDeltaSize   string            `json:"max_delta_size,omitempty"`
	MaxDeltaType   MaxDeltaType      `json:"max_delta_type,omitempty"`
	GroupProductId string            `json:"group_product_id,omitempty"`
	RealizedSpread string            `json:"realized_spread,omitempty"`
	PrimaryLeg     *SpreadLegDetails `json:"primary_leg,omitempty"`
	SecondaryLeg   *SpreadLegDetails `json:"secondary_leg,omitempty"`
}

// LegPreview contains the preview of a single leg of a SPREAD order.
type LegPreview struct {
	ProductId           string    `json:"product_id,omitempty"`
	Side                OrderSide `json:"side,omitempty"`
	SupportPassiveOrder bool      `json:"support_passive_order,omitempty"`
	LegCommission       string    `json:"leg_commission,omitempty"`
	QuoteValue          string    `json:"quote_value,omitempty"`
	BestBid             string    `json:"best_bid,omitempty"`
	BestAsk             string    `json:"best_ask,omitempty"`
	AverageFilledPrice  string    `json:"average_filled_price,omitempty"`
	Slippage            string    `json:"slippage,omitempty"`
}

// SpreadPreview contains the preview of a SPREAD order.
type SpreadPreview struct {
	Limit          string          `json:"limit,omitempty"`
	LimitType      SpreadLimitType `json:"limit_type,omitempty"`
	MaxDeltaSize   string          `json:"max_delta_size,omitempty"`
	MaxDeltaType   MaxDeltaType    `json:"max_delta_type,omitempty"`
	GroupProductId string          `json:"group_product_id,omitempty"`
	PrimaryLeg     *LegPreview     `json:"primary_leg,omitempty"`
	SecondaryLeg   *LegPreview     `json:"secondary_leg,omitempty"`
}

// OptionGreeks contains the greeks of an option order preview.
type OptionGreeks struct {
	Delta string `json:"delta,omitempty"`
	Gamma string `json:"gamma,omitempty"`
	Theta string `json:"theta,omitempty"`
	Vega  string `json:"vega,omitempty"`
}

// CommissionDetailTotal contains a breakdown of all commission charges for an order or fill.
type CommissionDetailTotal struct {
	TotalCommission      string `json:"total_commission,omitempty"`
	ClientCommission     string `json:"client_commission,omitempty"`
	VenueCommission      string `json:"venue_commission,omitempty"`
	CesCommission        string `json:"ces_commission,omitempty"`
	FinancingCommission  string `json:"financing_commission,omitempty"`
	RegulatoryCommission string `json:"regulatory_commission,omitempty"`
	ClearingCommission   string `json:"clearing_commission,omitempty"`
}

// OrderFill represents a fill on an order
type OrderFill struct {
	Id                    string                 `json:"id"`
	OrderId               string                 `json:"order_id"`
	Side                  string                 `json:"side"`
	ProductId             string                 `json:"product_id"`
	ClientProductId       string                 `json:"client_product_id"`
	FilledQuantity        string                 `json:"filled_quantity"`
	FilledValue           string                 `json:"filled_value"`
	Price                 string                 `json:"price"`
	Time                  time.Time              `json:"time"`
	Commission            string                 `json:"commission"`
	Venue                 string                 `json:"venue"`
	VenueFees             string                 `json:"venue_fees"`
	CesCommission         string                 `json:"ces_commission"`
	ProductType           ProductType            `json:"product_type,omitempty"`
	CommissionDetailTotal *CommissionDetailTotal `json:"commission_detail_total,omitempty"`
}

// EditHistory represents an order edit entry (new format)
type EditHistory struct {
	Price            string `json:"price"`
	BaseQuantity     string `json:"base_quantity"`
	QuoteValue       string `json:"quote_value"`
	DisplayBaseSize  string `json:"display_base_size"`
	DisplayQuoteSize string `json:"display_quote_size"`
	StopPrice        string `json:"stop_price"`
	ExpiryTime       string `json:"expiry_time"`
	AcceptTime       string `json:"accept_time"`
	ClientOrderId    string `json:"client_order_id"`
}

// OrderEditHistory represents an order edit entry (deprecated format)
// Deprecated: Use EditHistory instead
type OrderEditHistory struct {
	Price          string `json:"price"`
	Size           string `json:"size"`
	DisplaySize    string `json:"display_size"`
	StopPrice      string `json:"stop_price"`
	StopLimitPrice string `json:"stop_limit_price"`
	EndTime        string `json:"end_time"`
	AcceptTime     string `json:"accept_time"`
	ClientOrderId  string `json:"client_order_id"`
}
