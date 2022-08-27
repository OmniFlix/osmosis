package types

import (
	"fmt"
	appparams "github.com/osmosis-labs/osmosis/v10/app/params"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys
var (
	KeySaleCreationFee               = []byte("SaleCreationFee")
	KeySaleCreationFeeRecipient      = []byte("SaleCreationFeeRecipient")
	KeyMinimumDurationUntilStartTime = []byte("MinimumDurationUntilStartTime")
	KeyMinimumSaleDuration           = []byte("MinimumSaleDuration")
)

// ParamTable for streamswap module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(saleCreationFee sdk.Coins, saleCreationFeeRecipient string, minimumDurationUntilStartTime, minimumSaleDuration time.Duration) Params {
	return Params{
		SaleCreationFee:           saleCreationFee,
		SaleCreationFeeRecipient:  saleCreationFeeRecipient,
		MinDurationUntilStartTime: minimumDurationUntilStartTime,
		MinSaleDuration:           minimumSaleDuration,
	}
}

// default streamswap module parameters
func DefaultParams() Params {
	return Params{
		SaleCreationFee:           sdk.NewCoins(sdk.NewInt64Coin(appparams.BaseCoinUnit, 200_000_000)), // 200 OSMO
		MinDurationUntilStartTime: time.Hour * 24,                                                      // 1 Day
		MinSaleDuration:           time.Hour * 24,                                                      // 1 Day
	}
}

// validate params
func (p Params) Validate() error {
	if err := validateSaleCreationFee(p.SaleCreationFee); err != nil {
		return err
	}
	if err := validateDuration(p.MinDurationUntilStartTime); err != nil {
		return err
	}
	if err := validateDuration(p.MinSaleDuration); err != nil {
		return err
	}
	return nil

}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeySaleCreationFee, &p.SaleCreationFee, validateSaleCreationFee),
		paramtypes.NewParamSetPair(KeySaleCreationFeeRecipient, &p.SaleCreationFeeRecipient, validateSaleCreationFeeRecipient),
		paramtypes.NewParamSetPair(KeyMinimumDurationUntilStartTime, &p.MinDurationUntilStartTime, validateDuration),
		paramtypes.NewParamSetPair(KeyMinimumSaleDuration, &p.MinSaleDuration, validateDuration),
	}
}
func validateSaleCreationFeeRecipient(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T, expected string bech32", i)
	}
	if _, err := sdk.AccAddressFromBech32(v); err != nil {
		return fmt.Errorf("invalid parameter type: expected string bech32")
	}
	return nil
}

func validateSaleCreationFee(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.Validate() != nil {
		return fmt.Errorf("invalid sale creation fee: %+v", i)
	}

	return nil
}

func validateDuration(i interface{}) error {
	_, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}
