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

import (
	"fmt"
	"time"

	"github.com/coinbase/core-go"
	"github.com/shopspring/decimal"
)

// TravelRuleWalletType represents the type of wallet for travel rule compliance
type TravelRuleWalletType string

const (
	TravelRuleWalletTypeUnspecified   TravelRuleWalletType = "TRAVEL_RULE_WALLET_TYPE_UNSPECIFIED"
	TravelRuleWalletTypeVASP          TravelRuleWalletType = "TRAVEL_RULE_WALLET_TYPE_VASP"
	TravelRuleWalletTypeSelfCustodied TravelRuleWalletType = "TRAVEL_RULE_WALLET_TYPE_SELF_CUSTODIED"
)

// NaturalPersonName represents natural person name components
type NaturalPersonName struct {
	FirstName  string `json:"first_name,omitempty"`
	MiddleName string `json:"middle_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
}

// DetailedAddress represents detailed address information
type DetailedAddress struct {
	Address1    string `json:"address_1,omitempty"`
	Address2    string `json:"address_2,omitempty"`
	Address3    string `json:"address_3,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
}

// TravelRuleDate represents a date for travel rule (year, month, day)
type TravelRuleDate struct {
	Year  int32 `json:"year,omitempty"`
	Month int32 `json:"month,omitempty"`
	Day   int32 `json:"day,omitempty"`
}

// TravelRuleData contains travel rule information for withdrawals.
type TravelRuleData struct {
	Beneficiary                   *TravelRuleParty `json:"beneficiary,omitempty"`
	Originator                    *TravelRuleParty `json:"originator,omitempty"`
	IsSelf                        bool             `json:"is_self,omitempty"`
	IsIntermediary                bool             `json:"is_intermediary,omitempty"`
	OptOutOfOwnershipVerification bool             `json:"opt_out_of_ownership_verification,omitempty"`
	AttestVerifiedWalletOwnership bool             `json:"attest_verified_wallet_ownership,omitempty"`
}

// CounterpartyDestination represents a destination for a counterparty payment.
type CounterpartyDestination struct {
	CounterpartyId string `json:"counterparty_id,omitempty"`
}

// TravelRuleParty represents a party in a travel rule transaction
type TravelRuleParty struct {
	Name              string               `json:"name,omitempty"`
	NaturalPersonName *NaturalPersonName   `json:"natural_person_name,omitempty"`
	Address           *DetailedAddress     `json:"address,omitempty"`
	WalletType        TravelRuleWalletType `json:"wallet_type,omitempty"`
	VaspId            string               `json:"vasp_id,omitempty"`
	VaspName          string               `json:"vasp_name,omitempty"`
	PersonalId        string               `json:"personal_id,omitempty"`
	DateOfBirth       *TravelRuleDate      `json:"date_of_birth,omitempty"`
}

// EstimatedNetworkFees represents estimated network fees for a transaction
type EstimatedNetworkFees struct {
	LowerBound string `json:"lower_bound,omitempty"`
	UpperBound string `json:"upper_bound,omitempty"`
}

// MatchMetadata represents metadata for matched transactions
type MatchMetadata struct {
	ReferenceId    string `json:"reference_id,omitempty"`
	SettlementDate string `json:"settlement_date,omitempty"`
}

// RewardSubtype represents the reward subtype
type RewardSubtype string

const (
	RewardSubtypeUnknown              RewardSubtype = "REWARD_SUBTYPE_UNKNOWN"
	RewardSubtypeMEV                  RewardSubtype = "MEV_REWARD"
	RewardSubtypeInflation            RewardSubtype = "INFLATION_REWARD"
	RewardSubtypeBlock                RewardSubtype = "BLOCK_REWARD"
	RewardSubtypeTransaction          RewardSubtype = "TRANSACTION_REWARD"
	RewardSubtypeStakingFeeRebate     RewardSubtype = "STAKING_FEE_REBATE_REWARD"
	RewardSubtypeBuidlDividend        RewardSubtype = "BUIDL_DIVIDEND"
	RewardSubtypeCustomStablecoin     RewardSubtype = "CUSTOM_STABLECOIN_REWARD"
)

// CustomStablecoinAsset contains currency metadata for a custom stablecoin reward program.
type CustomStablecoinAsset struct {
	Symbol string `json:"symbol,omitempty"`
}

// CustomStablecoinRewardDetails contains details for a custom stablecoin reward payout.
type CustomStablecoinRewardDetails struct {
	StartDate string                 `json:"start_date,omitempty"`
	EndDate   string                 `json:"end_date,omitempty"`
	Asset     *CustomStablecoinAsset `json:"asset,omitempty"`
}

// RewardMetadata represents metadata for reward transactions
type RewardMetadata struct {
	Subtype                       RewardSubtype                  `json:"subtype,omitempty"`
	CustomStablecoinRewardDetails *CustomStablecoinRewardDetails `json:"custom_stablecoin_reward_details,omitempty"`
}

// Web3TransactionMetadata represents metadata for web3 transactions
type Web3TransactionMetadata struct {
	Label                 string        `json:"label,omitempty"`
	ConfirmedAssetChanges []AssetChange `json:"confirmed_asset_changes,omitempty"`
}

// TransactionMetadata represents additional metadata for a transaction
type TransactionMetadata struct {
	MatchMetadata           *MatchMetadata           `json:"match_metadata,omitempty"`
	Web3TransactionMetadata *Web3TransactionMetadata `json:"web3_transaction_metadata,omitempty"`
	RewardMetadata          *RewardMetadata          `json:"reward_metadata,omitempty"`
}

// AssetChangeType represents the type of asset change
type AssetChangeType string

const (
	AssetChangeTypeBalanceTransfer AssetChangeType = "BALANCE_TRANSFER"
	AssetChangeTypeBalanceApproval AssetChangeType = "BALANCE_APPROVAL"
	AssetChangeTypeItemTransfer    AssetChangeType = "ITEM_TRANSFER"
	AssetChangeTypeItemApproval    AssetChangeType = "ITEM_APPROVAL"
	AssetChangeTypeItemApprovalAll AssetChangeType = "ITEM_APPROVAL_ALL"
)

// NFTCollection represents an NFT collection
type NFTCollection struct {
	Name string `json:"name,omitempty"`
}

// NFTItem represents an NFT item
type NFTItem struct {
	Name string `json:"name,omitempty"`
}

// AssetChange represents a change in asset for a transaction
type AssetChange struct {
	Type       AssetChangeType `json:"type,omitempty"`
	Symbol     string          `json:"symbol,omitempty"`
	Amount     string          `json:"amount,omitempty"`
	Collection *NFTCollection  `json:"collection,omitempty"`
	Item       *NFTItem        `json:"item,omitempty"`
}

// RiskAssessment represents risk assessment results for a transaction
type RiskAssessment struct {
	ComplianceRiskDetected bool `json:"compliance_risk_detected"`
	SecurityRiskDetected   bool `json:"security_risk_detected"`
}

// OnchainDetail represents on-chain details for a transaction
type OnchainDetail struct {
	SignedTransaction     string          `json:"signed_transaction"`
	RiskAssessment        *RiskAssessment `json:"risk_assessment"`
	ChainId               string          `json:"chain_id"`
	Nonce                 string          `json:"nonce"`
	ReplacedTransactionId string          `json:"replaced_transaction_id"`
	DestinationAddress    string          `json:"destination_address"`
	SkipBroadcast         bool            `json:"skip_broadcast"`
	FailureReason         string          `json:"failure_reason"`
	SigningStatus         string          `json:"signing_status"`
}

// Transfer represents a transfer from or to in a transaction
type Transfer struct {
	Type              string `json:"type"`
	Value             string `json:"value"`
	Address           string `json:"address"`
	AccountIdentifier string `json:"account_identifier"`
}

// ValueNum converts the transfer value string to a decimal
func (tr Transfer) ValueNum() (amount decimal.Decimal, err error) {
	amount, err = core.StrToNum(tr.Value)
	if err != nil {
		err = fmt.Errorf("invalid transfer value: %s - type: %s - msg: %v", tr.Value, tr.Type, err)
	}
	return
}

// Transaction represents a Prime transaction
type Transaction struct {
	Id                    string                `json:"id"`
	WalletId              string                `json:"wallet_id"`
	PortfolioId           string                `json:"portfolio_id"`
	Type                  string                `json:"type"`
	Status                string                `json:"status"`
	Symbol                string                `json:"symbol"`
	Created               time.Time             `json:"created_at"`
	Completed             time.Time             `json:"completed_at"`
	Amount                string                `json:"amount"`
	TransferFrom          *Transfer             `json:"transfer_from,omitempty"`
	TransferTo            *Transfer             `json:"transfer_to,omitempty"`
	NetworkFees           string                `json:"network_fees"`
	Fees                  string                `json:"fees"`
	FeeSymbol             string                `json:"fee_symbol"`
	BlockchainIds         []string              `json:"blockchain_ids"`
	TransactionId         string                `json:"transaction_id"`
	DestinationSymbol     string                `json:"destination_symbol"`
	EstimatedNetworkFees  *EstimatedNetworkFees `json:"estimated_network_fees,omitempty"`
	Network               string                `json:"network"`
	EstimatedAssetChanges []AssetChange         `json:"estimated_asset_changes"`
	Metadata              *TransactionMetadata  `json:"metadata,omitempty"`
	IdempotencyKey        string                `json:"idempotency_key"`
	OnchainDetails        *OnchainDetail        `json:"onchain_details,omitempty"`
}
